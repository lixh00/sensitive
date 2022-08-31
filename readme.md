## 脱敏 Desensitization

### 作用 Effect
利用反射处理结构体，处理带`sen`标签的字段(规则: `规则名称,占位符`，如:`phone,*`)，将其脱敏。支持自定义处理函数

Use reflection to process the structure, process the fields with the `sen` tag (rule: `rule name, placeholder`, such as: `phone,*`), and desensitize it. Support custom handlers


### 定义结构体示例 Example
```go
package main

import (
    "github.com/lixh00/sensitive"
)


type User struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Phone string `json:"phone" sen:"phone,*"`
    Email string `json:"email" sen:"email,*"`
}

func main() {
    data :=  User{
        Name: "lixh",
        Age: 18,
        Phone: "13888888888",
        Email: "lixh@gmail.com",
    }

    // 添加自定义处理函数
    // Add custom handler
    sensitive.AddHandler("email", func(src, p string) string {
        // 将@符号后面的替换为*
        // Replace the after @ sign with *
        idx := strings.Index(src, "@")
        dst := src[:idx+1] + strings.Repeat(p, utf8.RuneCountInString(src)-idx-1)
        
        return dst
    })
    
    if err := sensitive.Desensitize(data); err != nil {
        fmt.Println(err)
    }
    bs, _ := json.Marshal(response)
    log.Printf("after processing data: %v", string(bs))
}
```