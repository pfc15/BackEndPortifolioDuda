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

	tema := NewTema_sql("titulo", "foto", 1, "2025-02-02")

	s.NotNil(tema)
	s.Equal(tema.Titulo, "titulo")
	s.Equal(tema.Periodo, "2025-02-02")
	s.Equal(tema.Foto, 1)
}

func (s *TemaTestSuite) TestNewTema_sqlFailure() {
	tema := NewTema_sql("titulo", "foto_nao_existe", 1, "2025-02-02")
	s.Nil(tema)
}

func (s *TemaTestSuite) TearDownTest() {
	Db = s.original
}

func TestTemaSuite(t *testing.T) {
	suite.Run(t, new(TemaTestSuite))
}
