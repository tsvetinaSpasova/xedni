package tokenization

import (
	"log"
	"xedni/pkg/domain/document"

	"github.com/jdkato/prose/v2"
)

type TermRepository interface {
	Store(Term) error
	LoadByToken(string) (*Term, error)
}

type Term struct {
	Token  string
	DocIDs []string
}

func New(token string, di []string) (*Term, error) {
	return &Term{
		Token:  token,
		DocIDs: di,
	}, nil
}

func Tokenizing(d document.Document) ([]string, error) {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(d.Text)
	if err != nil {
		log.Fatal(err)
	}

	s := []string{}

	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		s = append(s, tok.Text)
	}

	return s, nil
}
