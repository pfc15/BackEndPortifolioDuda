package persistence

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TemaTestSuite struct {
	suite.Suite
	db       DataBaseInterface
	original DataBaseInterface
}

func (s *TemaTestSuite) SetupTest() {
	s.db = &MockDb{}
	s.original = Db
	Db = s.db
}

func (s *TemaTestSuite) TestNewTema_sqlSuccess() {

	tema := NewTema_sql("titulo", "foto")

	s.NotNil(tema)
	if tema != nil {
		s.Equal(tema.Titulo, "titulo")
		s.Equal(tema.Foto, 1)
	}

}

func (s *TemaTestSuite) TestNewTema_sqlFailure() {
	tema := NewTema_sql("titulo", "foto_nao_existe")
	s.Nil(tema)
}

func (s *TemaTestSuite) TearDownTest() {
	Db = s.original
}

func TestTemaSuite(t *testing.T) {
	suite.Run(t, new(TemaTestSuite))
}
