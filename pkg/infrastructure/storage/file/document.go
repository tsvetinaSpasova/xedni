package file

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"xedni/pkg/domain/document"

	"github.com/rs/zerolog"
)

const docsPath = "/var/tmp/xedni/documents/"

type DocumentRepository struct {
	Logger *zerolog.Logger
}

// Store mock for in-memory storage
func (r *DocumentRepository) Store(d document.Document) error {
	bs, err := json.Marshal(d)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(docsPath+d.ID, bs, 0644)
}

// LoadByID mock for in-memory storage.
func (r *DocumentRepository) LoadByID(ID string) (*document.Document, error) {
	bs, err := ioutil.ReadFile(docsPath + ID)
	if err != nil {
		return nil, err
	}
	var d document.Document
	err = json.Unmarshal(bs, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// NewDocumentRepository instantiates an example memory repository
func NewDocumentRepository(_ context.Context, options map[string]interface{}, logger *zerolog.Logger) (*DocumentRepository, error) {
	return &DocumentRepository{
		Logger: logger,
	}, nil
}
