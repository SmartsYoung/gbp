package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	//连接
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//通过Database.C()方法切换集合（Collection）
	//func (db Database) C(name string) *Collection
	c := session.DB("test").C("people")

	//插入
	//func (c *Collection) Insert(docs ...interface{}) error
	err = c.Insert(&Person{"superWang", "13478808311"},
		&Person{"David", "15040268074"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	//查询
	//func (c Collection) Find(query interface{}) Query
	err = c.Find(bson.M{"name": "superWang"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Phone:", result.Phone)
}
