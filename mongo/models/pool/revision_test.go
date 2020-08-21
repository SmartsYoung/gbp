package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

const (
	DBName = "kie"

	CollectionLabel         = "label"
	CollectionKV            = "kv"
	CollectionKVRevision    = "kv_revision"
	CollectionPollingDetail = "polling_detail"
	CollectionCounter       = "counter"
	CollectionView          = "view"
	DefaultTimeout          = 5 * time.Second
	DefaultValueType        = "text"
)

func TestIncreaseAndGetRevision(t *testing.T) {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://kie:123@localhost:27017/kie")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	s1 := Counter{Name: "test", Count: 0}
	collection := client.Database("kie").Collection("rev")
	insertResult, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	s := &Service{}
	n, _ := s.GetRevision(client, context.TODO(), "default")
	t.Log(n)

	next, _ := ApplyRevision(client, context.TODO(), "default")
	t.Log(next)
	assert.Equal(t, n+1, next)
}

//Counter is db schema
type Counter struct {
	Name  string `bson:"name,omitempty"`
	Count int64  `bson:"count,omitempty"`
}

const revision = "revision_counter"

//Service is the implementation
type Service struct {
}

//GetRevision return current revision number
func (s *Service) GetRevision(client *mongo.Client, ctx context.Context, domain string) (int64, error) {
	collection := client.Database("kie").Collection("rev")

	filter := bson.M{"name": revision, "domain": domain}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("err", err)
		if err.Error() == context.DeadlineExceeded.Error() {
			msg := "operation timeout"
			fmt.Println(msg)
			return 0, errors.New(msg)
		}
		return 0, err
	}
	defer cur.Close(ctx)
	c := &Counter{}
	for cur.Next(ctx) {
		if err := cur.Decode(c); err != nil {
			fmt.Println("decode error: " + err.Error())
			return 0, err
		}
	}
	return c.Count, nil
}

func ApplyRevision(client *mongo.Client, ctx context.Context, domain string) (int64, error) {
	collection := client.Database("kie").Collection("rev")
	filter := bson.M{"name": revision, "domain": domain}
	sr := collection.FindOneAndUpdate(ctx, filter,
		bson.D{
			{"$inc", bson.D{
				{"count", 1},
			}}}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var result interface{}
	fmt.Println("hello:", sr.Decode(&result))
	if sr.Err() != nil {
		fmt.Println("sr:", sr.Err())
		return 0, sr.Err()
	}
	c := &Counter{}
	err := sr.Decode(c)
	if err != nil {
		fmt.Println("decode error: " + err.Error())
		return 0, err
	}
	return c.Count, nil
}
