package utils

import (
	"context"
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"strings"
)

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
	locale        = "chinese"
)

func TransInit(ctx context.Context) context.Context {
	//设置支持语言
	chinese := zh.New()
	english := en.New()
	//设置国际化翻译器
	uni := ut.New(chinese, chinese, english)
	//设置验证器
	val := validator.New()
	//根据参数取翻译器实例
	trans, _ := uni.GetTranslator(locale)
	//翻译器注册到validator
	switch locale {
	case "chinese":
		zhTranslations.RegisterDefaultTranslations(val, trans)
		//使用fld.Tag.Get("comment")注册一个获取tag的自定义方法
		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})
		val.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
			phone := fl.Field().String()
			//使用正则表达式验证手机号码
			pattern := `^1[3456789]\d{9}$`
			matched, _ := regexp.MatchString(pattern, phone)
			return matched
		})
		zhTranslations.RegisterDefaultTranslations(val, trans)
		val.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
			return ut.Add("phone", "{0}格式不正确，必须为手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("phone", fe.Field())
			return t
		})

	case "english":
		enTranslations.RegisterDefaultTranslations(val, trans)
		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("en_comment")
		})
		val.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
			phone := fl.Field().String()
			//使用正则表达式验证手机号码
			pattern := `^1[3456789]\d{9}$`
			matched, _ := regexp.MatchString(pattern, phone)
			return matched
		})
		zhTranslations.RegisterDefaultTranslations(val, trans)
		val.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
			return ut.Add("phone", "The {0} format is incorrect;it must be a cell phone number", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("phone", fe.Field())
			return t
		})
	}

	ctx = context.WithValue(ctx, ValidatorKey, val)
	ctx = context.WithValue(ctx, TranslatorKey, trans)
	return ctx
}

func DefaultGetValidParams(ctx context.Context, params interface{}) error {
	ctx = TransInit(ctx)
	err := validate(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func validate(ctx context.Context, params interface{}) error {
	//获取验证器
	val, ok := ctx.Value(ValidatorKey).(*validator.Validate)
	if !ok {
		return errors.New("Validator not found in context")
	}

	//获取翻译器
	tran, ok := ctx.Value(TranslatorKey).(ut.Translator)

	if !ok {
		return errors.New("Translator not found in context")
	}
	err := val.Struct(params)
	//如果数据效验不通过，则将所有err以切片形式输出
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			//使用validator.ValidationErrors类型里的Translate方法进行翻译
			sliceErrs = append(sliceErrs, e.Translate(tran))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
