package _type

import (
	"fmt"
	"os"
)

//定义一个接口
type Abc interface {
	String() string
}

// 类型
type Efg struct {
	data string
}

// 类型Efg实现Abc接口   方法
func (e *Efg) String() string {
	return e.data
}

// 获取一个*Efg实例
func GetEfg() *Efg {
	return &Efg{data: ""}
}

// 比较
func CheckAE(a Abc) bool {
	return a == nil
}
func main() {
	efg := GetEfg()
	b := CheckAE(efg)
	fmt.Println(b)
	os.Exit(1)
}
