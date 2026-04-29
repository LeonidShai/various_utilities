package yamlparser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var (
	TestFileName = "yaml_parser/test.yaml"
	testJsonStr  = `{"humans":{"name":"Petya","age":30,"parents":[{"mom":"Alex"},{"dad":"Gordon","city":"Moscow"}],"plemya":"indeec"}}`
)

func ParseYamlFile(filename string) {
	type Fru struct {
		Fruits []string `yaml:"fruits"`
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)

		return
	}

	fru := &Fru{}
	if err := yaml.Unmarshal(data, fru); err != nil {
		fmt.Printf("failed to unmarshal yaml: %v", err)

		return
	}

	for _, elem := range fru.Fruits {
		fmt.Println(elem)
	}
}

func ConvertJsonToYaml() {
	// Декодируем JSON в map
	var data map[string]interface{}
	err := json.Unmarshal([]byte(testJsonStr), &data)
	if err != nil {
		fmt.Printf("failed to decode JSON: %v", err)

		return
	}

	dt, ok := data["humans"]
	if ok {
		fmt.Println("Found humans")
	} else {
		fmt.Println("Failed to find humans")
	}

	_, ok = data["plemya"]
	if ok {
		fmt.Println("Found plemya")
	} else {
		fmt.Println("Failed to find plemya")
	}

	// Кодируем map в YAML
	// yamlBytes, err := yaml.Marshal(&dt)
	// if err != nil {
	// 	fmt.Printf("failed to convert to YAML: %v", err)

	// 	return
	// }

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)

	if err := encoder.Encode(dt); err != nil {
		fmt.Printf("failed to decode: %v", err)

		return
	}

	if err := encoder.Close(); err != nil {
		fmt.Printf("failed to close: %v", err)

		return
	}

	// Выводим результат
	fmt.Println(buf.String())
}
