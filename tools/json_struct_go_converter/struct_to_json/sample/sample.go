// This is only sample.
package template

// The struct you want to convert. Struct `input` is the default.
// You might change the contents of the struct. Please do not rename struct `input`.
// The result will be generated at output/output.json.

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

// Leave it empty if you want. It will be filled with default value based on the type data.
// var Input interface{} = &input{}
