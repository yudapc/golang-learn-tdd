package tdd

import (
	"encoding/json"
	"fmt"
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

func GetUserFromJson(jsonData string) (User, error) {
	var jsonToByte = []byte(jsonData)
	var data User

	var err = json.Unmarshal(jsonToByte, &data)
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}
	return data, nil
}

func JsonToMapStringInterface(jsonData string) (map[string]interface{}, error) {
	var dataUser map[string]interface{}
	var jsonToByte = []byte(jsonData)
	var err = json.Unmarshal(jsonToByte, &dataUser)
	if err != nil {
		fmt.Println(err.Error())
		return dataUser, err
	}
	return dataUser, nil
}

func DecodeJsonToInterface(jsonData string) map[string]interface{} {
	var jsonToByte = []byte(jsonData)
	var data interface{}
	json.Unmarshal(jsonToByte, &data)
	var decodedData = data.(map[string]interface{})
	return decodedData
}
