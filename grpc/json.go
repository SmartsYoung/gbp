package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float32
}

func main() {

	gobook := Book{
		"Go programming",
		[]string{"XuShiwei", "HughLv", "Johnson"},
		"isturing.com.cn",
		true,
		9.99,
	}

	// encode
	b, err := json.Marshal(gobook) // 变量b 是一个[]byte类型
	if err == nil {
		fmt.Println(b)
	}

	// decode
	var book Book
	json.Unmarshal(b, &book)
	fmt.Println(book)
}
