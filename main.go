package tdd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func Addition(first int, second int) int {
	return first + second
}

func Reduction(first int, second int) int {
	return first - second
}

func Division(first int, second int) int {
	return first / second
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUserFromJson(jsonParams string) (User, error) {
	var jsonToByte = []byte(jsonParams)
	var data User

	var err = json.Unmarshal(jsonToByte, &data)
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}
	return data, nil
}

func JsonToMapStringInterface(jsonParams string) (map[string]interface{}, error) {
	var dataUser map[string]interface{}
	var jsonToByte = []byte(jsonParams)
	var err = json.Unmarshal(jsonToByte, &dataUser)
	if err != nil {
		fmt.Println(err.Error())
		return dataUser, err
	}
	return dataUser, nil
}

func DecodeJsonToInterface(jsonParams string) map[string]interface{} {
	var jsonToByte = []byte(jsonParams)
	var data interface{}
	json.Unmarshal(jsonToByte, &data)
	var decodedData = data.(map[string]interface{})
	return decodedData
}

func DecodeArrayJSONToObject(jsonParams string) ([]User, error) {
	var jsonToByte = []byte(jsonParams)
	var users []User
	var err = json.Unmarshal(jsonToByte, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func EncodeObjectToJSON(jsonParams User) string {
	var jsonFromString, _ = json.Marshal(jsonParams)
	var data = string(jsonFromString)
	return data
}

func ValidationUsingJSONSchema(inputPayload string) (string, error) {
	var jsonByte = []byte(inputPayload)
	var jsonString = string(jsonByte)
	var documentLoader = gojsonschema.NewStringLoader(jsonString)

	var filePath = "file:///Users/yudaprabucogati/Workspace/go-workspace/src/belajar_validation_json_schema/login-param.json"
	schemaLoader := gojsonschema.NewReferenceLoader(filePath)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return "", err
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
		return "The document is valid", nil
	} else {
		var errorMessages []string
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
			errorMessages = append(errorMessages, desc.Description())
		}
		return "", errors.New(strings.Join(errorMessages, ", "))
	}
}
