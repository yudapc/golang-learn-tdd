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
