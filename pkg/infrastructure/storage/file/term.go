package file

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"xedni/pkg/domain/tokenization"

	"github.com/rs/zerolog"
)

const termsPath = "/var/tmp/xedni/terms/"

type TermRepository struct {
	Logger *zerolog.Logger
}

func (r *TermRepository) Store(t tokenization.Term) error {
	bs, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(termsPath+t.Token, bs, 0644)
}

func (r *TermRepository) LoadByToken(token string) (*tokenization.Term, error) {

	if _, err := os.Stat(termsPath + token); err != nil {
		return tokenization.New(token, nil)
	}

	bs, err := ioutil.ReadFile(termsPath + token)
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
