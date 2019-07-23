package tdd

import "testing"

func TestGetUserFromJson(test *testing.T) {
	var jsonString = `{"name": "Kamu lah", "age": 15}`
	var expected = "Kamu lah"
	user, _ := GetUserFromJson(jsonString)
	result := user.Name
	if result != expected {
		test.Errorf("The result from function GetUserFromJson is wrong!, user Name is %s not equal %s", expected, result)
	}
}

func TestGetUserFromJsonWhenHasError(test *testing.T) {
	var jsonString = `[{"namex": "Kamu lah", "ngapain": 15}]`
	_, err := GetUserFromJson(jsonString)
	if err == nil {
		test.Errorf("The result from function GetUserFromJson is correct!")
	}
}

func TestJsonToMapStringInterface(test *testing.T) {
	var jsonString = `{"name": "Kamu lah", "age": 15}`
	var expected float64 = 15
	data, _ := JsonToMapStringInterface(jsonString)
	if data["age"] != expected {
		test.Errorf("The result from function JsonToMapStringInterface is wrong!, data[\"age\"] = %0.f", data["age"])
	}
}

func TestJsonToMapStringInterfaceHasError(test *testing.T) {
	var jsonString = `[{"namex": "Kamu lah", "ngapain": 15}]`
	_, err := JsonToMapStringInterface(jsonString)
	if err == nil {
		test.Errorf("The result from function JsonToMapStringInterface is correct!, please make to wrong")
	}
}

func TestDecodeJsonToInterface(test *testing.T) {
	var jsonString = `{"name": "Jhon Doe", "age": 15}`
	data := DecodeJsonToInterface(jsonString)
	expected := "Jhon Doe"
	result := data["name"]
	if result != expected {
		test.Errorf("The result from function JsonToMapStringInterface is wrong!, data[\"name\"] = %s", data["name"])
	}
}

func TestDecodeArrayJSONToObject(test *testing.T) {
	var jsonString = `[
		{"name": "Dhani", "age": 35},
		{"name": "Jhon Doe", "age": 15}
	]`

	users, _ := DecodeArrayJSONToObject(jsonString)
	expectedName := "Dhani"
	expectedAge := 15
	resultName := users[0].Name
	resultAge := users[1].Age

	if expectedName != resultName || expectedAge != resultAge {
		test.Errorf("The result from function JsonToMapStringInterface is wrong!, resultName (%s) != %s, resultAge (%d) != %d", resultName, expectedName, resultAge, expectedAge)
	}
}

func TestDecodeArrayJSONToObjectHasError(test *testing.T) {
	var jsonString = `[
		{"name": "Dhani", "age": 35},
		{"name": "Jhon Doe", "age": 15
	]`

	_, err := DecodeArrayJSONToObject(jsonString)
	if err == nil {
		test.Errorf("The result from function DecodeArrayJSONToObject is correct, please change to error")
	}
}

func TestEncodeObjectToJSON(test *testing.T) {
	var jsonString = User{Name: "Jhon Doe", Age: 15}
	result := EncodeObjectToJSON(jsonString)
	if result == "" {
		test.Errorf("The result from function EncodeObjectToJSON is wrong!")
	}
}
