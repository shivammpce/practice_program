package main

import (
	"encoding/json"
	"log"

	"gopkg.in/yaml.v2"
)

func YamlToJson(yamlStr string) (string, error) {
	// Parse the YAML string
	var data interface{}
	err := yaml.Unmarshal([]byte(yamlStr), &data)
	if err != nil {
		return "", err
	}

	// Convert the parsed data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func main() {
	yamlString := `name: John
                  age: 30
                   city: New York`

	jsonString, err := YamlToJson(yamlString)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(jsonString)
}

// package main

// import (
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "log"

//     "gopkg.in/yaml.v2"
// )

// func convertYAMLToJSON(yamlFilePath string) ([]byte, error) {
//     // Read the YAML file
//     yamlFile, err := ioutil.ReadFile(yamlFilePath)
//     if err != nil {
//         return nil, err
//     }

//     // Convert YAML to JSON
//     var jsonData interface{}
//     err = yaml.Unmarshal(yamlFile, &jsonData)
//     if err != nil {
//         return nil, err
//     }

//     jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
//     if err != nil {
//         return nil, err
//     }

//     return jsonBytes, nil
// }

// func main() {
//     jsonBytes, err := convertYAMLToJSON("path/to/file.yaml")
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println(string(jsonBytes))
// }

// This function uses the ioutil.ReadFile function from the io/ioutil package to read the YAML file from disk.
//  It then uses the yaml.Unmarshal function from the gopkg.in/yaml.v2 package to convert the YAML data to a Go data structure.
//   Finally, it uses the json.MarshalIndent function from the encoding/json package to convert the Go data structure to JSON format.
//    The resulting JSON data is returned as a byte slice.

// Note that this function assumes that the YAML file contains valid YAML data that can be successfully unmarshalled to a Go data structure. If the YAML data is invalid, the yaml.Unmarshal function will return an error.
