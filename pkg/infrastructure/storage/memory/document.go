package memory

import (
	"context"
	"errors"
	"sync"

	"xedni/pkg/domain/document"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type DocumentRepository struct {
	m sync.Map
}

// Store mock for in-memory storage
func (r *DocumentRepository) Store(d document.Document) error {
	r.m.Store(d.ID, d)
	return nil
}

// LoadByID mock for in-memory storage.
func (r *DocumentRepository) LoadByID(ID uuid.UUID) (*document.Document, error) {
	v, ok := r.m.Load(ID)
	if !ok {
		// Match upper/db4 db.ErrNoMoreRows
		return nil, errors.New(`upper: no more rows in this result set`)
	}

	d := v.(document.Document)
	return &d, nil
}

// NewDocumentRepository instantiates an example memory repository
func NewDocumentRepository(_ context.Context, options map[string]interface{}, logger *zerolog.Logger) (*DocumentRepository, error) {
	return &DocumentRepository{}, nil
}
