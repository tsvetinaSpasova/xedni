package document

import (
	"errors"
	"net/http"

	"xedni/pkg/domain/document"
	"xedni/pkg/service"
	weberror "xedni/pkg/web/error"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

const (
	// You decide if you want to wrap errors or
	// will use values.
	ErrGetDocumentParam    = "get_document_param"
	ErrGetDocumentLoad     = "get_document_load"
	ErrCreateDocumentParam = "create_document_param"
	ErrCreateDocumentStore = "create_document_store"
)

// Handler is just a route collection
type Handler struct{}

// GetDocument Load a specific Document by ID - only "Document" will be found
func (h Handler) GetDocument(logger *zerolog.Logger, ds *service.DocumentService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		if ID == "" {
			render.Render(w, r, weberror.NewErrorResponse(ErrGetDocumentParam, http.StatusBadRequest, errors.New("passed an empty ID"), logger))
			return
		}

		document, err := ds.GetByID(ID)
		if err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrGetDocumentLoad, http.StatusBadRequest, err, logger))
			return
		}

		render.Render(w, r, NewFetchResponse(*document, ds))
	}
}

// CreateDocument allows HTTP creation.
func (h Handler) CreateDocument(logger *zerolog.Logger, ds *service.DocumentService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateRequest{}
		if err := render.Bind(r, &request); err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrCreateDocumentParam, http.StatusBadRequest, err, logger))
			return
		}

		err := ds.Store(request.Text)
		if err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrCreateDocumentStore, http.StatusBadRequest, err, logger))
			return
		}

		render.Render(w, r, NewCreateResponse(document.Document{ID: "demo", Text: request.Text}, ds))
	}
}

// Routes for document create/read
func (h Handler) Routes(logger *zerolog.Logger, ds *service.DocumentService) chi.Router {
	r := chi.NewRouter()

	r.Post("/document", h.CreateDocument(logger, ds))
	r.Get("/document/{ID}", h.GetDocument(logger, ds))

	return r
}
