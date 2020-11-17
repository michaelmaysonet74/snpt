package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SnippetList = Snippets{
	&Snippet{
		ID:        primitive.NewObjectID(),
		Title:     "Hello World",
		Content:   "print('Hello World!')",
		Language:  "Python",
		IsLoved:   false,
		CreatedOn: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedOn: primitive.NewDateTimeFromTime(time.Now()),
		CreatedBy: UserList[0].ID,
	},
	&Snippet{
		ID:    primitive.NewObjectID(),
		Title: "Sum Two Numbers",
		Content: `
			function add(a, b) {
				return a + b;
			}

			console.log(add(7, 4));
		`,
		Language:  "JavaScript",
		IsLoved:   true,
		CreatedOn: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedOn: primitive.NewDateTimeFromTime(time.Now()),
		CreatedBy: UserList[0].ID,
	},
}

var UserList = []*User{
	&User{
		ID:        primitive.NewObjectID(),
		FirstName: "Michael",
		LastName:  "Maysonet",
		Email:     "michaelmaysonet@test.com",
	},
}
