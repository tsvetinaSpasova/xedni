package tokenization_test

import (
	"flag"
	"testing"
	"xedni/pkg/domain/tokenization"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var configPath = flag.String("config-path", "", "Directory for the specific environment yaml configuration file.")

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupTest() {

}

func (s *TestSuite) TestInsert() {

	assert := assert.New(s.T())

	t := tokenization.Term{
		Token: "no-problem",
		DocIDs: []string{
			"1", "2", "4",
		},
	}

	err := t.Insert("3")
	assert.NoError(err)
	assert.Equal([]string{"1", "2", "3", "4"}, t.DocIDs)

	err = t.Insert("0")
	assert.NoError(err)
	assert.Equal([]string{"0", "1", "2", "3", "4"}, t.DocIDs)

	err = t.Insert("5")
	assert.NoError(err)
	assert.Equal([]string{"0", "1", "2", "3", "4", "5"}, t.DocIDs)

	t.DocIDs = nil
	err = t.Insert("3")
	assert.NoError(err)
	assert.Equal([]string{"3"}, t.DocIDs)

}

func (s *TestSuite) TestNewTerm() {
	assert := assert.New(s.T())

	t, err := tokenization.New("token", []string{"1", "2"})
	assert.NoError(err)
	assert.Equal("token", t.Token)
	assert.Equal([]string{"1", "2"}, t.DocIDs)
}
func (s *TestSuite) TestTokenize() {
	assert := assert.New(s.T())

	t, err := tokenization.Tokenize("Test tokenize")
	assert.NoError(err)
	assert.Equal([]string{"Test", "tokenize"}, t)
}
func TestAll(t *testing.T) {
	suite.Run(t, &TestSuite{})
}
