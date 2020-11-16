package models

import "time"

var SnippetList = Snippets{
	&Snippet{
		ID:        1,
		Title:     "Hello World",
		Content:   "print('Hello World!')",
		Language:  "Python",
		IsLoved:   false,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		CreatedBy: UserList[0].ID,
	},
	&Snippet{
		ID:    2,
		Title: "Sum Two Numbers",
		Content: `
			function add(a, b) {
				return a + b;
			}

			console.log(add(7, 4));
		`,
		Language:  "JavaScript",
		IsLoved:   true,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		CreatedBy: UserList[0].ID,
	},
}

var UserList = []*User{
	&User{
		ID:        1,
		FirstName: "Michael",
		LastName:  "Maysonet",
		Email:     "michaelmaysonet@test.com",
	},
}
