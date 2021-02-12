package document

type DocumentRepository interface {
	Store(Document) error
	LoadByID(string) (*Document, error)
}

type Document struct {
	ID   string
	Text string
}
