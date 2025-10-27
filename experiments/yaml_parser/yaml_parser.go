package yamlparser

import (
	"encoding/json"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

var (
	TestFileName = "yaml_parser/test.yaml"
	testJsonStr  = `{"humans":{"name":"Petya","age":30,"city":"Moscow"},"plemya":"indeec"}`
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

	// Кодируем map в YAML
	yamlBytes, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Printf("failed to convert to YAML: %v", err)

		return
	}

	// Выводим результат
	fmt.Println(string(yamlBytes))
}
