package tdd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func ValidationUsingJSONSchema(inputPayload string) (string, error) {
	var jsonByte = []byte(inputPayload)
	var jsonString = string(jsonByte)
	var documentLoader = gojsonschema.NewStringLoader(jsonString)

	var filePath = "file://./schemas/login-param.json"
	schemaLoader := gojsonschema.NewReferenceLoader(filePath)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return "", err
	}

	if result.Valid() {
		return "The document is valid", nil
	} else {
		var errorMessages []string
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
			errorMessages = append(errorMessages, desc.Description())
		}
		return "", errors.New(strings.Join(errorMessages, ", "))
	}
}
