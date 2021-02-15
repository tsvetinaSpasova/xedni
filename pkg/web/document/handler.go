package document

import (
	"net/http"

	"xedni/pkg/service"
	weberror "xedni/pkg/web/error"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-ozzo/ozzo-validation/is"
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rs/zerolog"
)

const (
	// You decide if you want to wrap errors or
	// will use values.
	ErrGetDocumentParam      = "get_document_param"
	ErrGetDocumentLoad       = "get_document_load"
	ErrCreateDocumentParam   = "create_document_param"
	ErrCreateDocumentStore   = "create_document_store"
	ErrSearchDocumentsParam  = "search_document_param"
	ErrSearchDocumentsSearch = "search_document_search"
)

// Handler is just a route collection
type Handler struct{}

// GetDocument Load a specific Document by ID - only "Document" will be found
func (h Handler) GetDocument(logger *zerolog.Logger, ds *service.DocumentService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")

		if err := ozzo.Validate(ID, is.UUIDv4); err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrGetDocumentParam, http.StatusBadRequest, err, logger))
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

		id, err := ds.Store(request.Text)
		if err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrCreateDocumentStore, http.StatusBadRequest, err, logger))
			return
		}

		render.Render(w, r, NewCreateResponse(*id, ds))
	}
}

// // Searchallows HTTP creation.
func (h Handler) Search(logger *zerolog.Logger, ds *service.DocumentService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := SearchRequest{}
		if err := render.Bind(r, &request); err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrSearchDocumentsParam, http.StatusBadRequest, err, logger))
			return
		}

		docs, err := ds.Search(request.Words)
		if err != nil {
			render.Render(w, r, weberror.NewErrorResponse(ErrSearchDocumentsSearch, http.StatusBadRequest, err, logger))
			return
		}

		render.Render(w, r, NewSearchResponse(docs, ds))
	}
}

// Routes for document create/read
func (h Handler) Routes(logger *zerolog.Logger, ds *service.DocumentService) chi.Router {
	r := chi.NewRouter()

	r.Post("/document", h.CreateDocument(logger, ds))
	r.Get("/document/{ID}", h.GetDocument(logger, ds))

	r.Post("/search", h.Search(logger, ds))

	return r
}
