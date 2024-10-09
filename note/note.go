package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title     string
	Content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Print("Your not titled %v has the following content: \n\n%v", note.Title, note.Content)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input")
	}

	return Note{
		Title:     title,
		Content:   content,
		createdAt: time.Now(),
	}, nil
}
