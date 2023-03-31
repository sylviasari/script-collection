package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	input "github.com/sylviasari/script-collection/tools/json_struct_go_converter/struct_to_json/input"
)

// if no value, use default value
// read from template
func main() {
	outputPath, _ := filepath.Abs("output/output.json")

	// Convert to json
	output, err := json.MarshalIndent(input.Input, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write to output
	_, err2 := f.WriteString(string(output))
	if err2 != nil {
		log.Fatal(err2)
	}
}
