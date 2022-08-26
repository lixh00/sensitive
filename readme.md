## 脱敏

### 作用
利用反射处理结构体，处理带`sen`标签的字段(规则: `规则名称,占位符`，如:`phone,*`)，将其脱敏。支持自定义处理函数

### 定义结构体示例
```go
package main

import (
    "github.com/lixh00/sensitive"
)


type User struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Phone string `json:"phone" sen:"phone,*"`
}

data :=  User{
    Name: "lixh",
    Age: 18,
	Phone: "13888888888",
}

if err := sensitive.Desensitize(data); err != nil {
    fmt.Println(err)
}
bs, _ := json.Marshal(response)
log.Printf("脱敏后的数据: %v", string(bs))
```