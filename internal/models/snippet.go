package models

import (
	"fmt"
)

type Snippet struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Language  string `json:"language"`
	IsLoved   bool   `json:"isLoved"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	CreatedBy int    `json:"createdBy"`
}

type Snippets []*Snippet

func generateID() int {
	return len(SnippetList) + 1
}

func findSnippet(id int) (*Snippet, int, error) {
	for i, s := range SnippetList {
		if s.ID == id {
			return s, i, nil
		}
	}

	return nil, 0, fmt.Errorf("Error: Snippet with ID[%v] was not found", id)
}

/**
 *	Public Methods
 */

func GetSnippets() Snippets {
	return SnippetList
}

func GetSnippetByID(id int) (*Snippet, error) {
	snippet, _, err := findSnippet(id)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

func CreateSnippet(s *Snippet) *Snippet {
	s.ID = generateID()
	s.CreatedBy = UserList[0].ID

	SnippetList = append(SnippetList, s)
	return s
}

func UpdateSnippet(id int, s *Snippet) (*Snippet, error) {
	_, i, err := findSnippet(id)
	if err != nil {
		return nil, err
	}

	SnippetList[i] = s
	return s, nil
}

func DeleteSnippet(id int) (Snippets, error) {
	_, i, err := findSnippet(id)
	if err != nil {
		return nil, err
	}

	// Remove found snippet from List
	SnippetList = append(SnippetList[:i], SnippetList[i+1:]...)
	return SnippetList, nil
}
