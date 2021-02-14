package document

import (
	"github.com/google/uuid"
)

type DocumentRepository interface {
	Store(Document) error
	LoadByID(string) (*Document, error)
}

type Document struct {
	ID   string
	Text string
}

func New(text string) (*Document, error) {
	return &Document{
		ID:   uuid.New().String(),
		Text: text,
	}, nil
}
