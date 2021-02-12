package service

import (
	"xedni/pkg/domain/document"

	"github.com/rs/zerolog"
)

// DocumentService contains the business logic for documents
type DocumentService struct {
	Repository document.DocumentRepository
	Logger     *zerolog.Logger
}

// GetByID - Proxy to repository
func (ds DocumentService) GetByID(ID string) (*document.Document, error) {
	// Stuff happens
	return ds.Repository.LoadByID(ID)
}

// Store converts raw text to a Documents record and saves to the respective repository.
func (ds DocumentService) Store(text string) error {
	// Stuff happens
	return ds.Repository.Store(document.Document{ID: "demo", Text: text})
}
