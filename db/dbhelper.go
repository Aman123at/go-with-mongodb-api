package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Aman123at/usermanage/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURL = "mongodb://127.0.0.1:27017"
const dbName = "usermanage"
const collectionName = "user"

var collection *mongo.Collection

func init() {
	fmt.Println("Initializing DB")

	clientOpt := options.Client().ApplyURI(mongoURL)

	client, err := mongo.Connect(context.TODO(), clientOpt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected successfully.")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Printf("%v database is ready", dbName)

	fmt.Printf("\n%v collection is ready", collectionName)
}

func InsertOneUserInDB(user model.User) {
	_, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
}

func UpdateOneUserInDB(userId string, updateData map[string]string) {
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	update := bson.M{"$set": updateData}

	_, errors := collection.UpdateOne(context.Background(), filter, update)

	if errors != nil {
		log.Fatal(err)
	}
}

func GetAllUsersFromDB() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var users []primitive.M

	for cur.Next(context.Background()) {
		var user bson.M
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer cur.Close(context.Background())

	return users
}
func GetOneUserFromDB(userId string) []primitive.M {
	id, hexErr := primitive.ObjectIDFromHex(userId)

	if hexErr != nil {
		log.Fatal(hexErr)
	}

	filter := bson.M{"_id": id}

	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	var users []primitive.M

	for cur.Next(context.Background()) {
		var user bson.M
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer cur.Close(context.Background())

	return users
}

func DeleteOneUserFromDB(userId string) {
	id, hexErr := primitive.ObjectIDFromHex(userId)

	if hexErr != nil {
		log.Fatal(hexErr)
	}

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
}
