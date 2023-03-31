# Golang struct to json converter
This is a script that will convert golang struct into json.

# Input
golang struct (non bulk struct)

*Sample*
```
type input struct {
	Name          string `json:"name"`
	Address       string `json:"address,omitempty"`
	FavoriteFoods []string
	Age           int   `json:"age"`
	Order         order `json:"order"`
}

type order struct {
	OrderID    int64     `json:"order_id"`
	Product    []product `json:"products"`
	GrandTotal float64   `json:"grand_total"`
}

type product struct {
	ProductName string `json:"product_name"`
}

var Input interface{} = &input{
	Name:          "John",
	FavoriteFoods: []string{"Noodles", "Fried chicken"},
	Age:           20,
	Order: order{
		OrderID: 123,
		Product: []product{
			{
				ProductName: "USB Cable",
			},
			{
				ProductName: "Charger",
			},
		},
		GrandTotal: float64(125000),
	},
}
```


# Output
json

*Sample*

```
{
    "name": "John",
    "FavoriteFoods": [
        "Noodles",
        "Fried chicken"
    ],
    "age": 20,
    "order": {
        "order_id": 123,
        "products": [
            {
                "product_name": "USB Cable"
            },
            {
                "product_name": "Charger"
            }
        ],
        "grand_total": 125000
    }
}
```


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
