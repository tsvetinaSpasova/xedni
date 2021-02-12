package document

import (
	"net/http"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateRequest is the payload shape for demo creation
type CreateRequest struct {
	Text string
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
