package main

import (
	"fmt"
	"github.com/SmartsYoung/gbp/interface/add"
	"github.com/SmartsYoung/gbp/interface/add/maths"
)

func main() {
	a, b := 1, 1
	m := &maths.Maths{}

	sum := m.AddSum(a, b)
	fmt.Println(sum)

	//调用接口前实例化接口   不能直接实例化接口， 接口不是变量类型
	add.MathsAdd = &maths.Maths{}

	sum1 := add.MathsAdd.AddSum(a, b)
	fmt.Println(sum1)

}

type inter struct {
}

func (i *inter) AddSum(a, b int) int {
	c, d := 1, 1
	add.MathsAdd = &maths.Maths{}
	sum1 := add.MathsAdd.AddSum(c, d)
	fmt.Println(sum1)

	return a + b + sum1
}
