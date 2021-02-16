package document

import (
	"net/http"

	"xedni/pkg/domain/document"
	"xedni/pkg/service"

	"github.com/go-chi/render"
)

// FetchResponse is the shape of data for a loaded document record
type FetchResponse struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// Render satisfies the chi interface
func (fr *FetchResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

// CreateResponse contains the ID post document creation
type CreateResponse struct {
	ID string `json:"id"`
}

// Render setups up the correct http status code.
func (cr *CreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

// SearchResponse contains the Docs post done search
type SearchResponse struct {
	Docs []document.Document `json:"docs"`
}

// Render satisfies the chi interface
func (sr *SearchResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d document.Document, _ *service.IndexService) *FetchResponse {
	return &FetchResponse{
		ID:   d.ID,
		Text: d.Text,
	}
}

// NewCreateResponse instantiates a new response when document is created
func NewCreateResponse(ID string, _ *service.IndexService) *CreateResponse {
	return &CreateResponse{ID: ID}
}

// NewSearchResponse instantiates a new response when search is done
func NewSearchResponse(docs []document.Document, _ *service.IndexService) *SearchResponse {
	return &SearchResponse{
		Docs: docs,
	}
}
