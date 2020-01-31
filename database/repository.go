package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sinbadflyce/dictcrawler/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName = "Dictionary"
var collectionName = "Longman"

// Repository ...
type Repository struct {
	uri    string
	client *mongo.Client
}

// Open ...
func (r *Repository) Open(uri string) bool {
	r.uri = uri
	clientOptions := options.Client().ApplyURI(r.uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
		return false
	}

	r.client = client
	return true
}

// Close ...
func (r *Repository) Close() {
	if r.client != nil {
		err := r.client.Disconnect(context.TODO())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Connection to MongoDB closed.")
		r.client = nil
	}
}

// Save ...
func (r *Repository) Save(word models.Word) {
	existW := r.Find(word.Name)

	if len(existW.Name) > 0 {
		fmt.Println("Word is already existing. Don not save!")
		return
	}

	jsonData, err := json.Marshal(word)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	// Check the connection
	err = r.client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to MongoDB!")

	collection := r.client.Database(databaseName).Collection(collectionName)

	insertResult, err := collection.InsertOne(context.TODO(), word)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// Find ...
func (r *Repository) Find(aWord string) models.Word {
	var result models.Word
	collection := r.client.Database(databaseName).Collection(collectionName)

	filter := bson.D{primitive.E{Key: "name", Value: aWord}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

// DictRepo ...
var DictRepo = Repository{}
