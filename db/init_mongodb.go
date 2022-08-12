package db

import (
	"context"
	"fmt"
	"github.com/Ccc-me/for-golang-test/db/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var mongoInstance *mongo.Client

func InitMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		AuthSource: "admin",
		Username:   "root",
		Password:   "qwer1234",
	}

	mongoUrl := "mongodb://mongoreplicae6dc0c93351f0.mongodb.volces.com:3717"

	clientOpts := options.Client().ApplyURI(mongoUrl).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		return err
	}

	coll := client.Database("douyincloud").Collection("count")
	doc := &model.MongoCount{
		Type:  "mongodb",
		Count: 2022,
	}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("documents inserted with result:%v\n", result)

	mongoInstance = client
	return err
}

func GetMongo() *mongo.Client {
	return mongoInstance
}
