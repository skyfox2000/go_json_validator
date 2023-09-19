package go_json_validator

import (
	"errors"
	"log"
	"strings"

	"github.com/buger/jsonparser"

	"github.com/dop251/goja"

	"github.com/dop251/goja_nodejs/require"
)

// Validator represents the Fastest Validator wrapper
type Validator struct {
	vm *goja.Runtime
}

func isValidJSON(s string) bool {
	s = strings.TrimSpace(s)
	if s == "{}" || s == "[]" {
		return true
	}

	if s == "" {
		return false
	}

	_, _, _, err := jsonparser.Get([]byte(s))
	return err == nil
}

// New creates a new Validator instance
func NewValidator(initOptions string) (*Validator, error) {
	if initOptions != "" && !isValidJSON(initOptions) {
		msg := "Invalid Validator options, " + initOptions
		log.Println(msg)
		return nil, errors.New(msg)
	}
	registry := new(require.Registry) // this can be shared by multiple runtimes
	vm := goja.New()
	registry.Enable(vm)

	// Load Fastest Validator JavaScript files
	_, err := vm.RunString(`
			const GoJsonValidator = require('./fastest-validator/initialize');
			
			let options = ` + initOptions + `;
			if (options == "") options = {};
		   const instance = new GoJsonValidator(options);
		`,
	)

	if err != nil {
		msg := "Init Error: " + err.Error() + ", " + initOptions
		log.Println(msg)
		return nil, errors.New(msg)
	}

	return &Validator{
		vm: vm,
	}, nil
}

// Validate validates the data by schema
func (v *Validator) Validate(schema string, jsonData string) (goja.Value, error) {
	if !isValidJSON(schema) {
		msg := "Invalid JSON Schema: " + schema
		log.Println(msg)
		return nil, errors.New(msg)
	}
	if !isValidJSON(jsonData) {
		msg := "Invalid JSON Data: " + jsonData
		log.Println(msg)
		return nil, errors.New(msg)
	}
	_, err := v.vm.RunString(`
		let schema = ` + schema + `;
		const validResult = instance.validate(schema, ` + jsonData + `);
		validResult;`)

	if err != nil {
		msg := "Validate Error: " + err.Error()
		log.Println(msg)
		return nil, errors.New(msg)
	}

	validResult := v.vm.Get("validResult")

	return validResult, nil
}
