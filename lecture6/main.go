package main

import (
	"encoding/json"
	"fmt"
)

func codeGenerator(data map[string]interface{}, indent int) string {
	code := ""
	for key, value := range data {
		code += fmt.Sprintf("%s%s := ", getIndent(indent), key)

		switch valueType := value.(type) {
		case int, float64, bool:
			code += fmt.Sprintf("%v\n", value)
		case string:
			code += fmt.Sprintf("\"%s\"\n", value)
		case map[string]interface{}:
			code += "{\n" + codeGenerator(value.(map[string]interface{}), indent+1) + getIndent(indent) + "}\n"
		case []interface{}:
			code += "{\n" + sliceCode(value.([]interface{}), indent+1) + getIndent(indent) + "}\n"
		default:
			fmt.Printf("Unsupported data type for key %s: %T\n", key, valueType)
		}
	}
	return code
}

func sliceCode(data []interface{}, indent int) string {
	code := ""
	for i, item := range data {
		code += fmt.Sprintf("%s[%d] := ", getIndent(indent), i)

		switch itemType := item.(type) {
		case int, float64, bool:
			code += fmt.Sprintf("%v\n", item)
		case string:
			code += fmt.Sprintf("\"%s\"\n", item)
		case map[string]interface{}:
			code += "{\n" + codeGenerator(item.(map[string]interface{}), indent+1) + getIndent(indent) + "}\n"
		case []interface{}:
			code += "{\n" + sliceCode(item.([]interface{}), indent+1) + getIndent(indent) + "}\n"
		default:
			fmt.Printf("Unsupported data type in array at index %d: %T\n", i, itemType)
		}
	}
	return code
}

func generateFunctions(functions []interface{}) string {
	code := ""
	for _, function := range functions {
		functionData, ok := function.(map[string]interface{})
		if !ok {
			fmt.Println("Invalid function format")
			continue
		}

		functionName, nameOk := functionData["function_name"].(string)
		parametersData, paramsOk := functionData["parameters"].([]interface{})
		returnType, returnTypeOk := functionData["return_type"].(string)
		codeData, codeOk := functionData["code"].(string)

		if !nameOk || !paramsOk || !codeOk || !returnTypeOk {
			fmt.Println("Invalid function format")
			continue
		}

		code += fmt.Sprintf("func %s(", functionName)
		for i, param := range parametersData {
			paramData, ok := param.(map[string]interface{})
			if !ok {
				fmt.Println("Invalid parameter format")
				continue
			}
			paramName, nameOk := paramData["param_name"].(string)
			paramType, typeOk := paramData["param_type"].(string)
			if !nameOk || !typeOk {
				fmt.Println("Invalid parameter format")
				continue
			}
			if i > 0 {
				code += ", "
			}
			code += fmt.Sprintf("%s %s", paramName, paramType)
		}
		code += fmt.Sprintf(") %s {\n", returnType)
		code += fmt.Sprintf("%s", codeData)
		code += "\n}\n\n"
	}
	return code
}

func getIndent(indent int) string {
	return "    "
}

func main() {
	jsonStr := `{
		"variables": {
			"var1": 10,
			"var2": "Hello, World!",
			"var3": true,
			"var4": [1, 2, 3, 4, 5],
			"var5": {
				"nested_var1": "ABC",
				"nested_var2": 42.5
			}
		},
		"functions": [
			{
				"function_name": "myFunction",
				"parameters": [
				  {
					"param_name": "a",
					"param_type": "int"
				  },
				  {
					"param_name": "b",
					"param_type": "int"
				  }
				],
				"return_type": "int",
				"code": "result := a + b\\nreturn result"
			}						
		]
	}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	goCode := codeGenerator(data["variables"].(map[string]interface{}), 0)
	functionCode := generateFunctions(data["functions"].([]interface{}))

	fmt.Printf("package main\n\n")
	fmt.Printf("import \"fmt\"\n\n")
	fmt.Printf("func main() {\n")
	fmt.Printf("%s", goCode)
	fmt.Printf("%s", functionCode)
	fmt.Printf("}\n")
}
