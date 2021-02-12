package document_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"xedni/pkg/configuration"
	"xedni/pkg/infrastructure/log"
	"xedni/pkg/web"
	"xedni/pkg/web/document"
	weberror "xedni/pkg/web/error"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var configPath = flag.String("config-path", "", "Directory for the specific environment yaml configuration file.")

type DocumentSuite struct {
	suite.Suite
	r chi.Router
}

func (s *DocumentSuite) SetupTest() {
	appConfiguration := &configuration.AppConfiguration{}
	absoluteConfigPath, err := filepath.Abs(*configPath)
	if err != nil {
		fmt.Printf("could not load configuration path with error [%s]", err.Error())
		os.Exit(1)
	}

	err = configuration.LoadYAML(appConfiguration, &absoluteConfigPath, nil, nil)
	if err != nil {
		fmt.Printf("could not load configuration path with error [%s]", err.Error())
		os.Exit(1)
	}

	logger, err := log.NewZerolog(os.Stdout, appConfiguration.LogLevel)
	if err != nil {
		fmt.Printf("could not instantiate the logger [%s]", err.Error())
		os.Exit(1)
	}
	s.r = web.NewRouter(context.Background(), appConfiguration, logger)
}

func (s *DocumentSuite) TestDocumentHandlerCreate() {
	assert := assert.New(s.T())

	// Any label gets the Document ID
	createPayload, err := json.Marshal(document.CreateRequest{
		Text: "I better be there",
	})
	assert.Nil(err)

	createRequest := httptest.NewRequest("POST", "/api/document", bytes.NewReader(createPayload))
	rr := httptest.NewRecorder()

	s.r.ServeHTTP(rr, createRequest)
	assert.Equal(http.StatusCreated, rr.Code)

	var happyResponse document.CreateResponse
	err = json.Unmarshal(rr.Body.Bytes(), &happyResponse)
	assert.Nil(err)
	assert.Equal("demo", happyResponse.ID, "ID of created record did not match expectation.")

	// Creation with empty label does not pass validation
	createPayloadEmpty, err := json.Marshal(document.CreateRequest{
		Text: "",
	})
	assert.Nil(err)

	rr = httptest.NewRecorder()
	emptyCreateRequest := httptest.NewRequest("POST", "/api/document", bytes.NewReader(createPayloadEmpty))
	s.r.ServeHTTP(rr, emptyCreateRequest)
	assert.Equal(http.StatusBadRequest, rr.Code)

	var errorResponse weberror.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResponse)
	assert.Nil(err)
	assert.Equal(document.ErrCreateDocumentParam, errorResponse.Message, "Error response for create validation did not match.")

}

func TestAll(t *testing.T) {
	suite.Run(t, &DocumentSuite{})
}
