package handle

import (
	"strings"
	"unicode/utf8"
)

//	IdCard
//	@description: 脱敏规则: 身份证号码
//	@param src string: 待处理字符串
//	@param placeholder string: 占位符
//	@return dst string: 脱敏后的数据
func IdCard(src, placeholder string) (dst string) {
	// 保留前六位后两位
	dst = src[:6] + strings.Repeat(placeholder, utf8.RuneCountInString(src)-7) + src[len(src)-2:]
	return
}
