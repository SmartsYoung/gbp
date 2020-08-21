package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string `bson:"name,omitempty"`
	Age  int64  `bson:"age,omitempty"`
}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://kie:123@localhost:27017/kie")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	s1 := Student{"xiaohong", 12}
	s2 := Student{"xiaolan", 10}
	s3 := Student{"xiaohuang", 11}

	collection := client.Database("kie").Collection("student")

	insertResult, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	students := []interface{}{s2, s3}
	insertManyResult, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	filter := bson.D{{"name", "xiaolan"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// var stu Student
	/*  var stu = make(map[string]interface{})
	    res := collection.FindOneAndUpdate(context.TODO(), bson.D{{"name", "xiaohong"}}, bson.M{"$set": bson.M{"name": "modify"}}).Decode(stu)

	    fmt.Println(stu)
	    fmt.Println("_id", stu["_id"])
	    log.Println(res)*/

	sr := collection.FindOneAndUpdate(context.TODO(), bson.D{{"name", "xiaohong"}},
		bson.D{
			{"$inc", bson.D{
				{"age", 1},
			}}}, options.FindOneAndUpdate().SetReturnDocument(options.After))

	if sr.Err() != nil {
		fmt.Println("sr:", sr.Err())
		return
	}

	st := &Student{}
	err = sr.Decode(st)
	if err != nil {
		fmt.Println("decode error: " + err.Error())
		return
	}
	fmt.Printf("FindOneAndUpdate查询到的数据:%v, %v\n", st.Name, st.Age)

}
