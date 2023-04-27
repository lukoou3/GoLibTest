package test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到mongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("连接失败")
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("ping 失败")
	}

}

func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

func TestMoCreateCollection(t *testing.T) {
	client := GetMgoCli()
	db := client.Database("test")
	db.CreateCollection(context.TODO(), "test_db")
}

func TestMoInsert(t *testing.T) {
	client := GetMgoCli()
	db := client.Database("test")
	collection := db.Collection("test_db")
	data := bson.M{"_id": "2", "name": "青丝2", "age": 18}
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("插入数据失败")
		return
	}
	fmt.Println(*insertResult)
}

func TestMoUpdate(t *testing.T) {
	client := GetMgoCli()
	db := client.Database("test")
	collection := db.Collection("test_db")
	filter := bson.M{"_id": "2"}
	update := bson.M{"$set": bson.M{"name": "青丝22", "age": 28}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
	}
}

func TestMoInsertUpsert(t *testing.T) {
	client := GetMgoCli()
	db := client.Database("test")
	collection := db.Collection("test_db")
	filter := bson.M{"_id": "3"}
	update := bson.M{"$set": bson.M{"name": "青丝", "age": 18}}
	opts := options.Update().SetUpsert(true) //关键是这个
	r, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		fmt.Println("插入数据失败")
		return
	}
	fmt.Println(r)
}

// findAll 根据条件查找全部符合的数据
func TestMofindAll(t *testing.T) {
	client := GetMgoCli()
	db := client.Database("test")
	collection := db.Collection("test_db")
	filter := bson.M{}
	res, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var persons []bson.M
	err = res.All(context.TODO(), &persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(persons)
}

type Data struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func TestMofindAll2(t *testing.T) {
	client := GetMgoCli()
	db := client.Database("test")
	collection := db.Collection("test_db")
	filter := bson.M{}
	res, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var persons []Data
	err = res.All(context.TODO(), &persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(persons)
}

func TestMofindOne(t *testing.T) {
	client := GetMgoCli()
	db := client.Database("test")
	collection := db.Collection("test_db")
	filter := bson.M{"_id": "3"}
	var person Data
	err := collection.FindOne(context.TODO(), filter).Decode(&person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
}
