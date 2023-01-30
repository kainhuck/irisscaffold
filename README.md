# irisscaffold

iris开发模板

## 命名规范(参考go-zero规范)

命名准则
- 当变量名称在定义和最后一次使用之间的距离很短时，简短的名称看起来会更好。 
- 变量命名应尽量描述其内容，而不是类型
- 常量命名应尽量描述其值，而不是如何使用这个值
- 在遇到for，if等循环或分支时，推荐单个字母命名来标识参数和返回值
- method、interface、type、package推荐使用单词命名
- package名称也是命名的一部分，请尽量将其利用起来
- 使用一致的命名风格

文件命名规范
- 全部小写
- 除unit test外避免下划线(_)
- 文件名称不宜过长

变量命名规范参考
- 首字母小写
- 驼峰命名
- 见名知义，避免拼音替代英文
- 不建议包含下划线(_)
- 不建议包含数字

适用范围
- 局部变量
- 函数出参、入参

函数、常量命名规范
- 驼峰式命名
- 可exported的必须首字母大写
- 不可exported的必须首字母小写
- 避免全部大写与下划线(_)组


## 编码规范(参考go-zero规范)

import
- 单行import不建议用圆括号包裹
- 按照官方包，NEW LINE，当前工程包，NEW LINE，第三方依赖包顺序引入
```go
import (
    "context"
    "string"

    "greet/user/internal/config"

    "google.golang.org/grpc"
)
```

函数返回
- 对象避免非指针返回
- 遵循有正常值返回则一定无error，有error则一定无正常值返回的原则

错误处理
- 有error必须处理，如果不能处理就必须抛出。
- 避免下划线(_)接收error

函数体编码
- 建议一个block结束空一行，如if、for等
```go
func main (){
    if x==1{
        // do something
    }

    fmt.println("xxx")
}
```
- return前空一行
```go
func getUser(id string)(string,error){
....

    return "xx",nil
}
```