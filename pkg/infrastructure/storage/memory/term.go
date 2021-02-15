package memory

import (
	"context"
	"sync"
	"xedni/pkg/domain/tokenization"

	"github.com/rs/zerolog"
)

type TermRepository struct {
	m sync.Map
}

// Store mock for in-memory storage
func (r *TermRepository) Store(t tokenization.Term) error {
	r.m.Store(t.Token, t)
	return nil
}

// LoadByToken mock for in-memory storage.
func (r *TermRepository) LoadByToken(token string) (*tokenization.Term, error) {
	v, ok := r.m.Load(token)
	if !ok {
		// Match returning an empty term
		return tokenization.New(token, nil)
	}

	d := v.(tokenization.Term)
	return &d, nil
}

// NewTermRepository instantiates an example memory repository
func NewTermRepository(_ context.Context, options map[string]interface{}, logger *zerolog.Logger) (*TermRepository, error) {
	return &TermRepository{}, nil
}
