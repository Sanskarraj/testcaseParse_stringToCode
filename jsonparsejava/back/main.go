package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func handleJSONObject(object interface{}, key, indentation string) string {
	var result string

	switch t := object.(type) {
	case string:
		// If the value is a string, wrap it in double quotes <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[1](https://stackoverflow.com/questions/38867692/how-to-parse-json-array-in-go)</sup> <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[4](https://www.codemio.com/2021/02/advanced-golang-tutorials-dynamic-json-parsing.html)</sup> 
		result = fmt.Sprintf("%s%s[] = {\"%s\"};", indentation, key, t)
	case float64:
		// If the value is a number, convert it to an integer <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[1](https://stackoverflow.com/questions/38867692/how-to-parse-json-array-in-go)</sup> <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[4](https://www.codemio.com/2021/02/advanced-golang-tutorials-dynamic-json-parsing.html)</sup> 
		result = fmt.Sprintf("%s%s[] = {%d};", indentation, key, int(t))
	case []interface{}:
		var arrayElements []string
		for index, v := range t {
			// Recursively handle nested arrays <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[4](https://www.codemio.com/2021/02/advanced-golang-tutorials-dynamic-json-parsing.html)</sup> <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[8](https://www.sohamkamani.com/golang/json/)</sup> 
			arrayElements = append(arrayElements, handleJSONObject(v, key+strconv.Itoa(index+1), indentation+"\t"))
		}
		result = strings.Join(arrayElements, "\n")
	}

	return result
}

func main() {
	jsonArray := `{
		"data": [
			[2, 7, 11, 15],
			[3, 2, 4],
			[1, 5, 5],
			[1, 2, 3, 4, 5],
			[5, 3, 5, 7]
		]
	}`

	var results map[string]interface{}

	// Unmarshal JSON to the interface <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[4](https://www.codemio.com/2021/02/advanced-golang-tutorials-dynamic-json-parsing.html)</sup> <sup className="rounded-full text-xs cursor-pointer [&>*]:!text-white h-4 w-4 px-1 bg-zinc-400 hover:bg-zinc-500 dark:bg-zinc-700 hover:dark:bg-zinc-600">[8](https://www.sohamkamani.com/golang/json/)</sup> 
	json.Unmarshal([]byte(jsonArray), &results)

	var output string
	for k, v := range results {
		output += handleJSONObject(v, k, "") + "\n"
	}

	fmt.Println(output)
}
