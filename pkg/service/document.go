package service

import (
	"xedni/pkg/domain/document"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// DocumentService contains the business logic for documents
type DocumentService struct {
	Repository document.DocumentRepository
	Logger     *zerolog.Logger
}

// GetByID - Proxy to repository
func (ds DocumentService) GetByID(ID uuid.UUID) (*document.Document, error) {
	// Stuff happens
	return ds.Repository.LoadByID(ID)
}

// Store converts raw text to a Documents record and saves to the respective repository.
func (ds DocumentService) Store(text string) (*uuid.UUID, error) {
	// Stuff happens

	doc, err := document.New(text)
	if err != nil {
		return nil, err
	}

	if err = ds.Repository.Store(*doc); err != nil {
		return nil, err
	}

	return &doc.ID, nil
}
