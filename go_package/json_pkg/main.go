package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	//var json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	//last := gjson.Get(json, "name.last")
	//age := gjson.Get(json, "age")
	//fmt.Printf("last=%v,last-type=%T\n", last.String(), last.String())
	//fmt.Printf("age=%v,age-type=%T\n", age.Int(), age.Int())

	jsonData := `{
		"products": [
			{"name": "苹果", "money": 11},
			{"name": "香蕉", "money": 1},
			{"name": "橘子", "money": 1},
		]
	}`

	result := gjson.Get(jsonData, "products.#(money < 2)")
	cheapProducts := result.Array()
	fmt.Println("便宜的水果:", cheapProducts)
}
