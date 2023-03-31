// Define the struct that needs to convert to json.
package input

// The struct you want to convert. Struct `input` is the default.
// You might change the contents of the struct. Please do not rename struct `input`.
// The result will be generated at output/output.json.

type input struct {
}

var Input interface{} = &input{}

// Leave it empty if you want. It will be filled with default value based on the type data.
// var Input interface{} = &input{}
