package file

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"xedni/pkg/domain/tokenization"

	"github.com/rs/zerolog"
)

type TermRepository struct {
	Logger *zerolog.Logger
}

func (r *TermRepository) Store(t tokenization.Term) error {
	bs, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return ioutil.WriteFile("/var/tmp/xedni/terms/"+t.Token, bs, 0644)
}

func (r *TermRepository) LoadByToken(token string) (*tokenization.Term, error) {
	bs, err := ioutil.ReadFile("/var/tmp/xedni/terms/" + token)
	if err != nil {
		return nil, err
	}
	var d tokenization.Term
	err = json.Unmarshal(bs, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

func NewTermRepository(_ context.Context, options map[string]interface{}, logger *zerolog.Logger) (*TermRepository, error) {
	return &TermRepository{
		Logger: logger,
	}, nil
}
