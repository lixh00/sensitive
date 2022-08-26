package sensitive

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

// Response 返回值
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// PageData 分页数据通用结构体
type PageData struct {
	Current   int   `json:"current"`    // 当前页码
	Size      int   `json:"size"`       // 每页数量
	Total     int64 `json:"total"`      // 总数
	TotalPage int   `json:"total_page"` // 总页数
	Records   any   `json:"records"`    // 返回数据
}

// Data 模拟返回数据结构体
type Data struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone" sen:"phone,*"`       // 使用脱敏标签，规则为手机号，替换占位字符为*
	IdNumber string `json:"idNumber" sen:"idNumber,#"` // 使用脱敏标签，规则为身份证号码，替换占位字符为*
}

// 模拟测试一下
func TestDeal(t *testing.T) {
	data := []Data{{
		Id:       "123",
		Name:     "张三",
		Phone:    "13800138000",
		IdNumber: "420102199010101010",
	}, {
		Id:       "234",
		Name:     "李四",
		Phone:    "13800138001",
		IdNumber: "420102199010101011",
	}}

	pageData := PageData{
		Current:   1,
		Size:      10,
		Total:     2,
		TotalPage: 1,
		Records:   &data,
	}

	response := Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: pageData,
	}

	bs, _ := json.Marshal(response)
	log.Printf("脱敏前的数据: %v", string(bs))

	//if err := response.Desensitization(true); err != nil {
	//	log.Println(err)
	//}

	//bs, _ = json.Marshal(response)
	//log.Printf("假设是管理员，需要跳过处理，脱敏后的数据: %v", string(bs))

	if err := Desensitization(response, false); err != nil {
		log.Println(err)
	}

	bs, _ = json.Marshal(response)
	log.Printf("脱敏后的数据: %v", string(bs))
}
