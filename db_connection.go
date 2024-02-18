package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI = "mongodb+srv://jitesh:hrishi3688@cluster0.l6a8xav.mongodb.net/?retryWrites=true&w=majority"

type Database struct {
	Context            context.Context
	Database           *mongo.Database
	EmployeeCollection *mongo.Collection
	LeaveCollection    *mongo.Collection
}

func connectDB() *Database {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPI)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	database := client.Database("test")
	employeeCollection := database.Collection("employees")
	// employeeCollection.Find(ctx, bson.M{"team": "Rockwell"})
	leaveCollection := database.Collection("leaves")

	// Send a ping to confirm a successful connection
	if err := database.RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return &Database{
		Context:            ctx,
		Database:           database,
		EmployeeCollection: employeeCollection,
		LeaveCollection:    leaveCollection,
	}
}

// func (db Database) insert(collectionName string, document User) *mongo.InsertOneResult {
// 	collection := db.Database.Collection(collectionName)
// 	result, err := collection.InsertOne(ctx, document)
// 	if err != nil {
// 		log.Fatal("Could not insert document", err)
// 	}
// 	return result
// }

func (db Database) insert(collectionName string, document []interface{}) (*mongo.InsertManyResult, error) {
	collection := db.Database.Collection(collectionName)
	result, err := collection.InsertMany(db.Context, document)
	if err != nil {
		log.Fatal("Could not insert document. Error:-\n\t", err)
		return nil, err
	}
	return result, nil
}

func (db Database) find(collectionName string, filter bson.D) ([]interface{}, error) {
	var result []interface{}
	collection := db.Database.Collection(collectionName)
	res, err := collection.Find(db.Context, filter)
	res.Decode(&result)
	if err != nil {
		log.Fatal("Could not find documents. Error:-\n\t", err)
		return nil, err
	}
	return result, nil
}
