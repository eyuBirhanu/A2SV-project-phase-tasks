package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser models.User
	err := UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	role := "user"
	count, _ := UserCollection.CountDocuments(ctx, bson.M{})
	if count == 0 {
		role = "admin"
	}

	newUser := models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	result, err := UserCollection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}

	newUser.ID = result.InsertedID.(primitive.ObjectID)
	return &newUser, nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func PromoteUser(userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"role": "admin"}}
	result, err := UserCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}