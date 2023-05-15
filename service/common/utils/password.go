package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

func GeneratePassword(length int) string {
	// 定义密码包含的字符集
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~!@#$%^&*()_+`-={}|[]\\:\";'<>,.?/"

	// 定义密码长度
	// 这里可以根据实际需求进行调整
	passwordLength := length

	// 初始化密码切片
	password := make([]byte, passwordLength)

	// 生成随机密码
	for i := 0; i < passwordLength; i++ {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		password[i] = charset[charIndex.Int64()]
	}

	return string(password)
}

// md5加密
func Md5(pasaword string) string {
	hash := md5.New()
	hash.Write([]byte(pasaword))
	passwordHash := hash.Sum(nil)
	// 将密码转换为16进制储存
	passwordHash16 := hex.EncodeToString(passwordHash)
	return passwordHash16
}

// 加盐值加密
func Md5Password(password, salt string) string {
	return Md5(password + salt)
}

// 解密
func ValidMd5Password(password, salt, dataPwd string) bool {
	return Md5Password(password, salt) == dataPwd
}
