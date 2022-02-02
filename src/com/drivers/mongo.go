package drivers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDrivers struct {
	Client   *mongo.Client
	Database string
}

var MgoClient *mongo.Client
var MgoDbName string

func Init() {
	MgoClient = Connect()
	MgoDbName = "data"
}

func Connect() *mongo.Client {
	// 设置客户端参数mongodb://user01:*****@localhost:27017/?authSource=data&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false
	//credential := options.Credential{
	//	Username: "user01",
	//	Password: "123456",
	//}
	clientOptions := options.Client().ApplyURI("mongodb://user01:123456@localhost:27017/?authSource=data&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	//defer client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// 检查链接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func Close() {
	err := MgoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
