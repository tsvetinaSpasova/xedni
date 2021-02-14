package document

import (
	"github.com/google/uuid"
)

type DocumentRepository interface {
	Store(Document) error
	LoadByID(uuid.UUID) (*Document, error)
}

type Document struct {
	ID   uuid.UUID
	Text string
}

func New(text string) (*Document, error) {
	return &Document{
		ID:   uuid.New(),
		Text: text,
	}, nil
}

// type Tokenization struct {
// }

// func (t Tokenization) Tokenizing() ([]string, error) {
// 	// Create a new document with the default configuration:
// 	doc, err := prose.NewDocument("Go is an open-source programming language lenguages created at Google GOOGLE.")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Iterate over the doc's tokens:
// 	for _, tok := range doc.Tokens() {
// 		fmt.Println(tok.Text)
// 		// Go NNP B-GPE
// 		// is VBZ O
// 		// an DT O
// 		// ...
// 	}

// 	// // Iterate over the doc's sentences:
// 	// for _, sent := range doc.Sentences() {
// 	// 	fmt.Println(sent.Text)
// 	// 	// Go is an open-source programming language created at Google.
// 	// }
// 	return nil, nil

// }
