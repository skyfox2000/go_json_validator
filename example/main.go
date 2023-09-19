package main

import (
	"fmt"
	"reflect"

	validator "github.com/skyfox2000/go_json_validator"
)

func main() {
	options := `{
	}`
	myValidator, err := validator.NewValidator(options)
	if err != nil {
		return
	}

	// 定义 JSON Schema 和数据
	schema := `{ "type": "string" }`
	data := "{'type':'rainect.com'}"

	// 验证数据
	result, err := myValidator.Validate(schema, data)
	if err != nil {
		return
	}

	resultType := result.ExportType()

	if resultType == reflect.TypeOf(true) {
		// 处理布尔类型结果
		valid := result.ToBoolean()
		fmt.Println("Validation result (boolean):", valid)
	} else if resultType == reflect.TypeOf([]interface{}{}) {
		// 处理数组类型结果
		errors := result.Export()
		fmt.Println("Validation errors (array):", errors)
	} else {
		// 处理其他类型的结果，或者报告错误
		fmt.Println("Unknown validation result type:", resultType)
	}

}
