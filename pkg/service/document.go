package service

import (
	"strings"
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

	tokens, err := tokenization.Tokenize(doc.Text)
	if err != nil {
		return nil, err
	}

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

func Merge(ids1 []string, ids2 []string) []string {
	answer := []string{}

	if ids2 == nil {
		return ids1
	}

	for i, j := 0, 0; i < len(ids1) && j < len(ids2); {
		switch strings.Compare(ids1[i], ids2[j]) {
		case 0:
			answer = append(answer, ids1[i])
			i++
			j++
		case -1:
			i++
		default:
			j++
		}
	}
	return answer
}

func (ds DocumentService) Search(words []string) ([]document.Document, error) {
	terms := []tokenization.Term{}

	for _, w := range words {
		t, err := ds.TermRepository.LoadByToken(w)
		if err != nil {
			return nil, err
		}
		terms = append(terms, *t)
	}

	var ids []string

	for _, t := range terms {
		ids = Merge(t.DocIDs, ids)
	}

	docs := []document.Document{}

	for _, id := range ids {
		d, err := ds.Repository.LoadByID(id)
		if err != nil {
			return nil, err
		}
		docs = append(docs, *d)
	}

	return docs, nil
}
