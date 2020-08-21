package main

import (
	"errors"
	"github.com/SmartsYoung/gbp/interface/add"
	"github.com/SmartsYoung/gbp/interface/addnew"
	"log"
)

// 官方文档中常用这种实现接口调用的方法，参见 io.Reader接口
func main() {
	sum, err := useAdd(addnew.NewAdd())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sum)
}

func useAdd(add add.Add) (int, error) {
	sum := add.AddSum(3, 4)
	if sum < 0 {
		return 0, errors.New("error add")
	}
	return sum, nil
}
