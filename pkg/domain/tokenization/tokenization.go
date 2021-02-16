package tokenization

import (
	"strings"

	"github.com/jdkato/prose/v2"
)

// TermRepository is an interface for the repositories that
// operate with terms on data level, a.k.a storing and loding
type TermRepository interface {
	Store(Term) error
	LoadByToken(string) (*Term, error)
}

// Term has Token and DocIDs
// Token is a word and DocIDs is an array of the document's ids in which occurrs the word
type Term struct {
	Token  string
	DocIDs []string
}

// binarySearch find the position where to place ID in IDs so that IDs stay
// sorted or returns true in secound parameter if element is already in slice
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

// Insert insert docunment id to a term's DocIDs
// And after insetion DocIDs stay sorted
// It returns any write error encountered.
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

// New creates a new term
func New(token string, di []string) (*Term, error) {
	return &Term{
		Token:  token,
		DocIDs: di,
	}, nil
}

// Tokenize extracts the normalized words out of the text
func Tokenize(text string) ([]string, error) {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	s := []string{}

	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		s = append(s, tok.Text)
	}

	return s, nil
}
