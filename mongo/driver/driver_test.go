package driver

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"testing"
)

func Test_Driver(t *testing.T) {
	Init("mongodb://localhost:27017")
	mgo := NewMgo("test", "test")
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
