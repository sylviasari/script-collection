# Golang struct to json converter
This is a script that will convert golang struct into json.

# Input
golang struct (non bulk struct)

# Output
json

# How to Run?
1. Fill the struct that needs to be converted into json on `input/input.go`. 
You can set the struct on `input` struct as default struct. You can fill any struct there and also the value. 
*Please do not rename struct `input`*
See the sample on `sample/sample.go`.
- If no value is filled, it will generate default value based on each data type
2. After the struct and the value are finished to fill, on terminal run command to go to directory:
`cd tools/json_struct_go_converter/struct_to_json`
3. Run `make run`
4. It will automatically convert the struct into json and generate the output on `output/output.json`

# Notes

# Library
Converter use `encoding/json`
- It will generate based on json tag: name, omitempty
- It's already using `MarshalIndent`, therefore the generated json is already applied indent.