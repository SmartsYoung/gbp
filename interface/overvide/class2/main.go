package main

import (
	"fmt"
)

type Animal struct {
	eat   string
	spark string
}

type Dog struct {
	Animal
	legs int
}

type Cat struct {
	Animal
	legs int
}

func (a Animal) Spark() {
	fmt.Println("Any animal can spark")
}

func main() {
	animal := Animal{"food", "wowo"}
	animal.Spark()
	dog := Dog{Animal{"bone", "wangwangwang"}, 4}
	dog.Spark()
	cat := Cat{Animal{"fish", "miaomiaomiao"}, 4}
	cat.Spark()
}

/*method的重写
上面的例子，如果Dog想要实现自己的Spark()方法怎么办？简单，和匿名字段冲突一样的道理，可以在Dog上定义一个Spark方法，重写匿名字段的Spark方法。*/

func (dog Dog) Spark() {
	fmt.Println("dog can spark:wangwangwang")
}

func (cat Cat) Spark() {
	fmt.Println("cat can spark:miaomiaomiao")
}
