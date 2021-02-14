package tokenization

import (
	"log"
	"strings"
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

func binarySearch(ID string, IDs []string) (int, bool) {

	low := -1
	high := len(IDs)

	for low+1 < high {
		median := (low + high) / 2

		switch strings.Compare(IDs[median], ID) {
		case 0:
			return 0, true
		case -1:
			low = median
		default:
			high = median
		}
	}

	return low, false
}

func (t *Term) Insert(id string) error {

	pos, ok := binarySearch(id, t.DocIDs)
	if ok {
		return nil
	}

	newIDs := make([]string, len(t.DocIDs)+1, len(t.DocIDs)+1)
	if pos >= 0 {
		copy(newIDs[:pos+1], t.DocIDs[:pos+1])
	}
	newIDs[pos+1] = id
	if pos+1 < len(t.DocIDs) {
		copy(newIDs[pos+2:], t.DocIDs[pos+1:])
	}
	t.DocIDs = newIDs

	return nil

}

func New(token string, di []string) (*Term, error) {
	return &Term{
		Token:  token,
		DocIDs: di,
	}, nil
}

func Tokenize(d document.Document) ([]string, error) {
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
