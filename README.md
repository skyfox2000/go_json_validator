# go_json_validator

`go_json_validator` 是一个使用 Go 语言包装的基于 [fastest-validator](https://github.com/icebob/fastest-validator) 的 JSON 格式验证器。它允许你使用与 `fastest-validator` 完全相同的 JSON Schema 格式来验证 JSON 数据。

## 安装

要使用 `go_json_validator`，你需要首先安装 Go 语言。然后，你可以使用以下命令来获取和安装包：

```bash
go get github.com/skyfox2000/go_json_validator
```

## 使用方法

### 初始化

首先，你需要初始化验证器。你可以在初始化过程中传入一些选项。以下是初始化验证器的示例代码：

```go
import (
	"fmt"
	validator "github.com/skyfox2000/go_json_validator/validator"
)

func main() {
	options := `{
		// 你的初始化选项，根据需要自定义
	}`
	myValidator, err := validator.NewValidator(options)
	if err != nil {
		fmt.Println("初始化验证器时出错：", err)
		return
	}

	// 后续的验证操作将使用 myValidator 进行。
}
```

### 验证数据

初始化后，你可以使用验证器来验证 JSON 数据。

你需要提供一个 JSON Schema 和待验证的 JSON 数据。

JSON Schema 需符合 [fastest-validator](https://github.com/icebob/fastest-validator) 规范

以下是验证数据的示例代码：

```go
// 定义 JSON Schema 和数据
schema := `{ "type": "string" }`
data := `{"type":"rainect.com"}`

// 验证数据
result, err := myValidator.Validate(schema, data)
if err != nil {
	fmt.Println("验证数据时出错：", err)
	return
}

// 处理验证结果
resultType := result.ExportType()
if resultType == reflect.TypeOf(true) {
	// 格式正确返回true
	valid := result.ToBoolean()
	fmt.Println("Validation result (boolean):", valid)
} else if resultType == reflect.TypeOf([]interface{}{}) {
	// 格式错误返回一个JSON数组，描述错误相关的字段和信息
	errors := result.Export()
	fmt.Println("Validation errors (array):", errors)
} else {
	// 处理其他类型的结果，或者报告错误
	fmt.Println("Unknown validation result type:", resultType)
}
```

## 许可证

该项目基于 MIT 许可证进行分发。更多详情请参阅 [LICENSE](LICENSE) 文件。
