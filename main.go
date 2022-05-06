package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fileCount := map[string]int{
		"cpp": 10,
		"js":  8,
		"go":  10,
	}
	bytes, _ := json.MarshalIndent(fileCount, "", "  ")
	fmt.Println(string(bytes))
}
