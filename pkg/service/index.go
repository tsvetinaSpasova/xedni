package service

import (
	"strings"
	"sync"
	"xedni/pkg/domain/document"
	"xedni/pkg/domain/tokenization"

	"github.com/rs/zerolog"
)

// IndexService contains the business logic for documents and terms
type IndexService struct {
	TermRepository     tokenization.TermRepository
	DocumentRepository document.DocumentRepository
	Logger             *zerolog.Logger
	m                  sync.Mutex
}

// GetByID - Proxy to repository
func (ins *IndexService) GetByID(ID string) (*document.Document, error) {
	return ins.DocumentRepository.LoadByID(ID)
}

// Index converts raw text to a Documents record and saves to the respective repository.
func (ins *IndexService) Index(text string) (*string, error) {

	doc, err := document.New(text)
	if err != nil {
		return nil, err
	}

	if err = ins.DocumentRepository.Store(*doc); err != nil {
		return nil, err
	}

	tokens, err := tokenization.Tokenize(doc.Text)
	if err != nil {
		return nil, err
	}

	ins.m.Lock()
	for _, token := range tokens {
		term, err := ins.TermRepository.LoadByToken(token)
		if err != nil {
			return nil, err
		}

		if err = term.Insert(doc.ID); err != nil {
			return nil, err
		}

		err = ins.TermRepository.Store(*term)
		if err != nil {
			return nil, err
		}
	}
	ins.m.Unlock()

	return &doc.ID, nil
}

//Merge merges to arrays only if the elements  on the both arrays are the same
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

//Search searches by words the documents that contain the words
func (ins *IndexService) Search(words []string) ([]document.Document, error) {
	terms := []tokenization.Term{}

	for _, w := range words {
		t, err := ins.TermRepository.LoadByToken(w)
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
		d, err := ins.DocumentRepository.LoadByID(id)
		if err != nil {
			return nil, err
		}
		docs = append(docs, *d)
	}

	return docs, nil
}
