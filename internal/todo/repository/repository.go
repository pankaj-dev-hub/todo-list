package repository

import (
	"context"
	"fmt"
	"log"
	"pankaj-katyare/todo-list/internal/todo/config"
	"pankaj-katyare/todo-list/internal/todo/model"

	"go.mongodb.org/mongo-driver/bson"
)

func Create(todo *model.Todo) (*model.Todo, error) {

	db, collection := config.NewConnection()

	fmt.Printf("Create Todo Tasks: %+v\n", todo)
	// Insert the user into the 	collection
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Print the inserted user's ID
	fmt.Println("Inserted user ID:", insertResult.InsertedID)

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Println("Disconnected from MongoDB")
	return todo, nil
}

func Update(id string, todo model.Todo) (*model.Todo, error) {

	db, colloction := config.NewConnection()

	fmt.Printf("Update Todo Tasks: %+v\n", todo)
	fmt.Println("Updated id: ", id)

	filter := map[string]string{"tasks.id": id}
	update := bson.M{"$set": bson.M{"tasks": todo.Tasks}}

	result, err := colloction.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Print the inserted user's ID
	fmt.Println("Inserted user ID:", result.UpsertedID)

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &todo, nil
}

func Get(id string) (*model.Todo, error) {

	db, collection := config.NewConnection()

	nestedJSON := map[string]string{"tasks.id": id}

	var todo model.Todo
	// Query the users collection
	err := collection.FindOne(context.Background(), nestedJSON).Decode(&todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Printf("Todo Tasks: %+v\n", todo.Tasks)

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &todo, nil

}

func GetAll() ([]*model.Todo, error) {

	db, collection := config.NewConnection()

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var todos []*model.Todo
	// Iterate over the cursor and print the users
	for cursor.Next(context.Background()) {
		var todo model.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		todos = append(todos, &todo)
		fmt.Printf("Todo Tasks: %+v\n", todo)
	}

	// Disconnect from MongoDB
	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return todos, err

}

func GetCompletedTask() ([]*model.Todo, error) {

	db, collection := config.NewConnection()

	var todo model.Todo

	fiter := map[string]string{"tasks.status": "completed"}

	cursor, err := collection.Find(context.Background(), fiter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("Todo Cursor: %+v\n", cursor)
	var todos []*model.Todo
	for cursor.Next(context.Background()) {

		err := cursor.Decode(&todo)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		todos = append(todos, &todo)
		fmt.Printf("Todo Tasks: %+v\n", todo)
	}

	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return todos, nil
}

func GetPendingTask() ([]*model.Todo, error) {

	db, collection := config.NewConnection()

	fiter := map[string]string{"tasks.status": "pending"}

	cursor, err := collection.Find(context.Background(), fiter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("Todo Cursor: %+v\n", cursor)
	var todos []*model.Todo
	for cursor.Next(context.Background()) {
		var todo model.Todo
		err := cursor.Decode(&todo)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		todos = append(todos, &todo)
	}
	for _, t := range todos {
		fmt.Printf("Todo Tasks: %+v\n", t)
	}

	err = db.Client().Disconnect(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return todos, nil
}
