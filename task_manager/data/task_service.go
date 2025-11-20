package data

import (
	"context"
	"log"
	"time"

	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func InitDB() {
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	TaskCollection = client.Database("task_manager").Collection("tasks")
	UserCollection = client.Database("task_manager").Collection("users")

	log.Println("Connected to MongoDB!")
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err = cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func CreateTask(task models.Task) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}
	task.ID = result.InsertedID.(primitive.ObjectID)
	return &task, nil
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	_, err = TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}
	updatedTask.ID = objID
	return &updatedTask, nil
}

func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}