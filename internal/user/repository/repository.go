package repository

import (
	"context"
	"fmt"
	"log"
	"pankaj-katyare/todo-list/internal/user/config"
	"pankaj-katyare/todo-list/internal/user/model"

	"go.mongodb.org/mongo-driver/bson"
)

func Create(user *model.User) *model.User {

	db, collection := config.NewConnection()

	fmt.Printf("Create User: %+v\n", user)
	// Insert the user into the 	collection
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	// Print the inserted user's ID
	fmt.Println("Inserted user ID:", insertResult.InsertedID)

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB")
	return user
}

func Update(id string, user model.User) model.User {

	db, colloction := config.NewConnection()

	fmt.Printf("Update User : %+v\n", user)
	fmt.Println("Updated id: ", id)

	filter := map[string]string{"id": id}
	update := bson.M{"$set": user}

	result, err := colloction.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// Print the inserted user's ID
	fmt.Println("Inserted user ID:", result.UpsertedID)

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func Get(id string) model.User {

	db, collection := config.NewConnection()

	nestedJSON := map[string]string{"id": id}

	var user model.User
	// Query the users collection
	err := collection.FindOne(context.Background(), nestedJSON).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User : %+v\n", user)

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return user

}

func GetAll() []*model.User {

	db, collection := config.NewConnection()

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var todos []*model.User
	// Iterate over the cursor and print the users
	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, &user)
		fmt.Printf("User : %+v\n", user)
	}

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return todos

}
