package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

type Fruit struct {
	Name     string `json":Name"`
	PriceTag string `json:"PriceTag"`
}

type FruitBasket struct {
	Name    string
	Fruit   []Fruit
	Id      int64 `json:"ref"` // 声明对应的json key
	Created time.Time
}

func Test_1(t *testing.T) {
	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit" : {"Name": "Apple", "PriceTag": "$1"},
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(basket)
	fmt.Println(string(jsonData))
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	fmt.Println(basket.Created)

}

func Test_2(t *testing.T) {
	jsondata := []byte(`
    {
        "Name": "Standard",
        "Fruit" :  [
	             {
	               "Name": "Apple",
                    "PriceTag": "$1"
	               },
	             {
	                "Name": "Pear",
                     "PriceTag": "$1.5"}
                  ],
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)
	var basket FruitBasket
	err := json.Unmarshal(jsondata, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(basket)
	fmt.Println(string(jsondata))
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	fmt.Println(basket.Created)
}

func Test_3(t *testing.T) {
	type Fruit struct {
		Name     string `json:"Name"`
		PriceTag string `json:"PriceTag"`
	}

	type FruitBasket struct {
		Name    string
		Fruit   map[string]Fruit
		Id      int64 `json:"ref"` // 声明对应的json key
		Created time.Time
	}
	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit" : {
	    "1": {
		"Name": "Apple",
		"PriceTag": "$1"
	    },
	    "2": {
		"Name": "Pear",
		"PriceTag": "$1.5"
	    }
        },
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range basket.Fruit {
		fmt.Println(item.Name, item.PriceTag)
	}

}

func Test_Marshal(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string `json:"name"`
		Colors []string
		note   string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(reflect.TypeOf(b)) // []uint8
	os.Stdout.Write(b)
}

func Test_UnMarshal(t *testing.T) {
	var jsonBlob = []byte(` [ 
        { "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
    ] `)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(reflect.TypeOf(animals)) // []main.Animal
	fmt.Printf("%+v", animals)
}

func Test_String(t *testing.T) {
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}

	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	type Color struct {
		Space string
		RGB   RGB
		YCbCr YCbCr
		Point json.RawMessage // delay parsing until we know the color space
	}

	var j = []byte(` [ 
        { "Space" : "YCbCr" , "YCbCr" : { "Y" : 255 , "Cb" : 0 , "Cr" : -10 } } , 
        { "Space" : "RGB" ,   "RGB" : { "R" : 98 , "G" : 218 , "B" : 255 } } 
    ] `)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}
	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}

func Test_str(t *testing.T) {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}
	var j = []byte(` [ 
        { "Space" : "YCbCr" , "Point" : { "Y" : 255 , "Cb" : 0 , "Cr" : -10 } } , 
        { "Space" : "RGB" ,   "Point" : { "R" : 98 , "G" : 218 , "B" : 255 } } 
    ] `)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}
	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}
