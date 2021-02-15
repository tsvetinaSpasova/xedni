package document

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateRequest is the payload shape for demo creation
type CreateRequest struct {
	Text string `json:"text"`
}

// Validate is a proxy method to confirm payload satisfies expectations
func (cr *CreateRequest) Validate() error {
	return ozzo.ValidateStruct(
		cr,
		ozzo.Field(&cr.Text, ozzo.Required),
	)
}

// Binder interface for chi
func (cr *CreateRequest) Bind(r *http.Request) error {
	return cr.Validate()
}

// SearchRequest is the payload shape for demo creation
type SearchRequest struct {
	Words []string `json:"words"`
}

// Validate is a proxy method to confirm payload satisfies expectations
func (sr *SearchRequest) Validate() error {
	return ozzo.ValidateStruct(
		sr,
		ozzo.Field(&sr.Words, ozzo.Required, ozzo.Length(1, 0)),
	)
}

// Binder interface for chi
func (sr *SearchRequest) Bind(r *http.Request) error {
	return sr.Validate()
}
