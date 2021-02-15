package document_test

import (
	"flag"
	"testing"
	"xedni/pkg/domain/document"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var configPath = flag.String("config-path", "", "Directory for the specific environment yaml configuration file.")

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupTest() {

}

func (s *TestSuite) TestNewDocument() {
	assert := assert.New(s.T())

	d, err := document.New("Test New")
	assert.NoError(err)
	assert.Equal("Test New", d.Text)
}

func TestAll(t *testing.T) {
	suite.Run(t, &TestSuite{})
}
