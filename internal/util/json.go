package util

import (
	"encoding/json"
	"fmt"
)

// JSON 파싱 함수
func ParseJSON(data []byte) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("JSON Parsing Error: %v\n", err)
		return nil
	}
	return result
}

// JSON 출력 함수
func PrintPrettyJSON(data []byte, title string) {

	var parsedData interface{}
	err := json.Unmarshal(data, &parsedData)
	if err != nil {
		fmt.Printf("[%s] JSON Parsing Error: %v\n", title, err)
		return
	}

	prettyJSON, err := json.MarshalIndent(parsedData, "", "  ")
	if err != nil {
		fmt.Printf("[%s] JSON Formatting Error: %v\n", title, err)
		return
	}

	fmt.Printf("[%s]\n%s\n", title, string(prettyJSON))
}
