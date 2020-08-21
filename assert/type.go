package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var f interface{}
	b := []byte(`[{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}]`)
	json.Unmarshal(b, &f)

	for k, v := range f.([]interface{}) {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int ", vv)
		case float64:
			fmt.Println(k, "is float64 ", vv)
		case []interface{}:
			fmt.Println(k, "is array:")
			for i, j := range vv {
				fmt.Println(i, j)
			}
		case interface{}:
			fmt.Println(k, "is :", vv)
		}
	}

}
