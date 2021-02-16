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

// Store stores document as a file in a directory with path: docsPath
// The file is named by document's ID and contains document's Text
func (r *DocumentRepository) Store(d document.Document) error {
	bs, err := json.Marshal(d)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(docsPath+d.ID, bs, 0644)
}

// LoadByID loads a document by finding the file that is named like ID
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

// NewDocumentRepository instantiates an example file repository
func NewDocumentRepository(_ context.Context, options map[string]interface{}, logger *zerolog.Logger) (*DocumentRepository, error) {
	return &DocumentRepository{
		Logger: logger,
	}, nil
}
