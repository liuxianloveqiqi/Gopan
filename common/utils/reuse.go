package utils

import (
	"strconv"
	"strings"
)

// api返回错误
func ApiError(err error) (c int, m string) {
	str := err.Error()
	// 提取状态码
	codeStart := strings.Index(str, "desc = ") + len("desc = ")
	codeEnd := strings.Index(str[codeStart:], ":")
	code := str[codeStart : codeStart+codeEnd]

	// 提取错误信息
	message := str[codeStart+codeEnd+1:]
	intcode, _ := strconv.Atoi(code)
	return intcode, message

}
