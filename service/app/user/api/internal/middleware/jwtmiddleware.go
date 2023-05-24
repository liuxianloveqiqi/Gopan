package middleware

import (
	"Gopan/service/common/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("开始jwt middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// JWTAuthMiddleware implementation
		fmt.Println(000000)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusBadRequest)
			err, _ := json.Marshal(errors.New("请求头中auth为空"))
			w.Write(err)
			fmt.Println("一错")
			return
		}
		fmt.Println(22222)
		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 3 && parts[0] == "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			err, _ := json.Marshal(errors.New("请求头中auth格式有误"))
			w.Write(err)
			fmt.Println("二错")

			return
		}
		fmt.Println(333333)
		parseToken, isExpire, err := utils.ParseToken(parts[1], parts[2])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			err, _ := json.Marshal(errors.New("token验证错误"))
			w.Write(err)
			fmt.Println("三错")
			return
		}
		if isExpire {
			parts[1], parts[2] = utils.GetToken(parseToken.ID, parseToken.State)
			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s %s", parts[1], parts[2]))
		}
		r = r.WithContext(context.WithValue(r.Context(), "user_id", parseToken.ID))
		fmt.Println(parseToken.ID, "gggggg")
		// Passthrough to next handler if need
		next(w, r)
	}
}
