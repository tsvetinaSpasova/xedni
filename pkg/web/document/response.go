package document

import (
	"net/http"

	"xedni/pkg/domain/document"
	"xedni/pkg/service"

	"github.com/go-chi/render"
)

// FetchResponse is the shape of data for a loaded demo record
type FetchResponse struct {
	document.Document
}

// Render satisfies the chi interface
func (fr *FetchResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

// CreateResponse contains the ID post demo creation
type CreateResponse struct {
	ID string
}

// Render setups up the correct http status code.
func (cr *CreateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

// NewFetchResponse instantiate a new response post load
func NewFetchResponse(d document.Document, _ *service.DocumentService) *FetchResponse {
	return &FetchResponse{d}
}

// NewCreateResponse instantiates a new response when demo is created
func NewCreateResponse(d document.Document, _ *service.DocumentService) *CreateResponse {
	return &CreateResponse{ID: "demo"}
}
