package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type Snippet struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Language  string `json:"language"`
	IsLoved   bool   `json:"isLoved"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	CreatedBy *User  `json:"-"`
}

type Snippets []*Snippet

func (s *Snippet) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(s)
}

func (s *Snippets) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(s)
}

func generateID() int {
	return len(SnippetList) + 1
}

func findSnippet(id int) (*Snippet, int, error) {
	for i, s := range SnippetList {
		if s.ID == id {
			return s, i, nil
		}
	}

	return nil, 0, fmt.Errorf("Snippet with id: " + string(id) + " not found.")
}

func GetSnippets() Snippets {
	return SnippetList
}

func GetSnippetById(id int) (*Snippet, error) {
	snippet, _, err := findSnippet(id)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

func CreateSnippet(s *Snippet) *Snippet {
	s.ID = generateID()
	s.CreatedBy = UserList[0]

	SnippetList = append(SnippetList, s)
	return s
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
