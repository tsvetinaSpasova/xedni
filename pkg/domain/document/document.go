package document

import (
	"github.com/google/uuid"
)

// DocumentRepository is an interface for the repositories that
// operate with documents on data level, a.k.a storing and loding
type DocumentRepository interface {
	Store(Document) error
	LoadByID(string) (*Document, error)
}

// Document is abstract struct of document that has ID and Text
type Document struct {
	ID   string
	Text string
}

// New creates a new term
func New(text string) (*Document, error) {
	return &Document{
		ID:   uuid.New().String(),
		Text: text,
	}, nil
}
