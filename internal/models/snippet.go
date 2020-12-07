package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Snippet struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title         string             `json:"title,omitempty" bson:"title,omitempty"`
	Content       string             `json:"content,omitempty" bson:"content,omitempty"`
	FileExtension string             `json:"fileExtension,omitempty" bson:"fileExtension,omitempty"`
	IsLoved       bool               `json:"isLoved" bson:"isLoved,omitempty"`
	CreatedOn     primitive.DateTime `json:"-" bson:"createdOn,omitempty"`
	UpdatedOn     primitive.DateTime `json:"-" bson:"updatedOn,omitempty"`
	CreatedBy     primitive.ObjectID `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
}

type Snippets []*Snippet

func generateID() int {
	return len(SnippetList) + 1
}

func findSnippet(c *mongo.Collection, id primitive.ObjectID) (*Snippet, error) {
	var snippet *Snippet

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.FindOne(ctx, bson.M{"_id": id}).Decode(&snippet)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

/**
 *	Public Methods
 */

func GetSnippets(c *mongo.Collection) (Snippets, error) {
	var snippets Snippets

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := c.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var snippet *Snippet
		cursor.Decode(&snippet)
		snippets = append(snippets, snippet)
	}

	return snippets, nil
}

func GetSnippetByID(c *mongo.Collection, id primitive.ObjectID) (*Snippet, error) {
	return findSnippet(c, id)
}

func CreateSnippet(c *mongo.Collection, s *Snippet) (*Snippet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.CreatedOn = primitive.NewDateTimeFromTime(time.Now())
	s.CreatedBy = UserList[0].ID

	result, err := c.InsertOne(ctx, s)
	if err != nil {
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)
	return findSnippet(c, id)
}

func UpdateSnippet(c *mongo.Collection, id primitive.ObjectID, s *Snippet) (*Snippet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.UpdateOne(ctx, Snippet{ID: id}, bson.M{"$set": s})
	if err != nil {
		return nil, err
	}

	return findSnippet(c, id)
}

func DeleteSnippet(c *mongo.Collection, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.DeleteOne(ctx, Snippet{ID: id})
	if err != nil {
		return err
	}

	return nil
}
