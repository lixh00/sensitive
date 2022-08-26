package sensitive

import "github.com/lixh00/sensitive/handle"

// RuleHandler 脱敏规则处理接口
type RuleHandler func(src, placeholder string) (dst string)

var senRuleMap map[string]RuleHandler // 脱敏规则Map

// 初始化函数，初始化默认的脱敏规则
func init() {
	senRuleMap = make(map[string]RuleHandler)
	AddHandler("phone", handle.Phone)
	AddHandler("idCard", handle.IdCard)
}

//	AddHandler
//	@description: 添加脱敏规则处理函数
//	@param name string: 脱敏规则名称
//	@param handler SensitiveRuleHandler: 脱敏规则处理函数
func AddHandler(name string, handler RuleHandler) {
	senRuleMap[name] = handler
}
