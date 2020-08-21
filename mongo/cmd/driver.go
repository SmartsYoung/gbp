package main

import (
	"context"
	"github.com/SmartsYoung/gbp/mongo/driver"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func main() {
	driver.Init("mongodb://kie:123@localhost:27017/kie")
	mgo := driver.NewMgo("kie", "test")
	var query []bson.M
	typeFilter := bson.M{"type": "go"}
	query = append(query, typeFilter)
	startTimeFilter := bson.M{"updateTimestamp": bson.M{"$gte": 0}}
	query = append(query, startTimeFilter)
	endTimeFilter := bson.M{"updateTimestamp": bson.M{"$lte": 1574870400000}}
	query = append(query, endTimeFilter)
	cursor, e := mgo.FindManyByFilters(query)
	if e != nil {
		log.Fatal(e)
	}
	for cursor.Next(context.TODO()) {
		println(cursor.Current.String())
	}
}
