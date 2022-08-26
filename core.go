package sensitive

import (
	"reflect"
	"strings"
)

//	Desensitization
//	@description: 脱敏
//	@param r any: 要脱敏的数据(只支持结构体)
//	@param skip bool: 是否跳过脱敏
//	@return err error: 错误信息
func Desensitization(r any, skip bool) (err error) {
	// 处理Data的值，如果也是结构体，就处理每个字段的标签，如果包含sen，就把sen改为sen_id
	if r == nil || skip {
		return
	}
	// 判断是不是结构体
	isPointer := reflect.TypeOf(r).Kind() == reflect.Ptr
	dataType := reflect.TypeOf(r).Kind()
	if isPointer {
		dataType = reflect.TypeOf(r).Elem().Kind()
	}
	//log.Printf("数据类型: %v，是否为指针: %v", dataType, isPointer)

	// 处理是数组的情况
	if dataType == reflect.Slice || dataType == reflect.Array {
		//log.Println("传入类型为数组")
		rs := reflect.ValueOf(r)
		if isPointer {
			rs = rs.Elem()
		}
		for i := 0; i < rs.Len(); i++ {
			id := rs.Index(i)
			err = Desensitization(id.Addr().Interface(), skip)
			if err != nil {
				return
			}
		}
	}

	// 如果是指针结构体，处理脱敏
	if dataType == reflect.Struct {
		val := reflect.ValueOf(r)
		if isPointer {
			val = val.Elem()
		}

		for i := 0; i < val.NumField(); i++ {
			f := val.Field(i)
			tag := val.Type().Field(i).Tag.Get("sen")

			//log.Printf("类型: %v -> %v 值: %v 脱敏标签是否存在: %v", f.Type(), f.Kind(), f.Interface(), tag != "")
			// 如果是结构体，递归调用
			if f.Kind() == reflect.Interface {
				//log.Println("开始处理子级")
				err = Desensitization(f.Interface(), skip)
			}

			if tag != "" {
				// 脱敏标签存在，处理一下，取出规则Id和占位符
				ruleId := strings.Split(tag, ",")[0]
				placeholder := strings.Split(tag, ",")[1]
				//log.Printf("脱敏规则Id: %v, 占位符: %v", ruleId, placeholder)

				// 处理脱敏
				if handle, ok := senRuleMap[ruleId]; ok {
					newData := handle(f.Interface().(string), placeholder)
					//log.Printf("脱敏后的值: %v", newData)
					f.SetString(newData)
				}
			}
		}
	}

	return
}
