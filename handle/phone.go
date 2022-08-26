package handle

import (
	"strings"
	"unicode/utf8"
)

//	Phone
//	@description: 脱敏规则: 手机号
//	@param src string: 待处理字符串
//	@param placeholder string: 占位符
//	@return dst string: 脱敏后的数据
func Phone(src, placeholder string) (dst string) {
	// 不足7位，直接返回
	if utf8.RuneCountInString(src) <= 7 {
		return src
	}
	// 取前三位和后四位
	dst = src[:3] + strings.Repeat(placeholder, utf8.RuneCountInString(src)-7) + src[len(src)-4:]
	return
}
