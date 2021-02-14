package service

import (
	"log"
	"xedni/pkg/domain/document"
	"xedni/pkg/domain/tokenization"

	"github.com/rs/zerolog"
)

// DocumentService contains the business logic for documents
type DocumentService struct {
	TermRepository tokenization.TermRepository
	Repository     document.DocumentRepository
	Logger         *zerolog.Logger
}

// GetByID - Proxy to repository
func (ds DocumentService) GetByID(ID string) (*document.Document, error) {
	// Stuff happens
	return ds.Repository.LoadByID(ID)
}

// Store converts raw text to a Documents record and saves to the respective repository.
func (ds DocumentService) Store(text string) (*string, error) {
	// Stuff happens

	doc, err := document.New(text)
	if err != nil {
		return nil, err
	}

	if err = ds.Repository.Store(*doc); err != nil {
		return nil, err
	}

	tokens, err := tokenization.Tokenize(*doc)
	if err != nil {
		return nil, err
	}
	log.Println(tokens)

	for _, token := range tokens {
		term, err := ds.TermRepository.LoadByToken(token)
		if err != nil {
			return nil, err
		}

		if err = term.Insert(doc.ID); err != nil {
			return nil, err
		}

		err = ds.TermRepository.Store(*term)
		if err != nil {
			return nil, err
		}
	}

	return &doc.ID, nil
}
