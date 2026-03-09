package persistence

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type obraTestSuite struct {
	suite.Suite
	db       DataBaseInterface
	original DataBaseInterface
}

func (s *obraTestSuite) SetupTest() {
	s.db = &MockDb{}
	s.original = Db
	Db = s.db
}

func (s *obraTestSuite) TestNewObraSuccess() {

	tema := NewObra("titulo", "foto", "2025-02-02",
		"descricao",
		"tema", "link", 1)

	s.NotNil(tema)
	s.Equal(tema.Titulo, "titulo")
	s.Equal(tema.Ordem, 1)
	s.Equal(tema.Periodo, "2025-02-02")
	s.Equal(tema.Tema, 1)
	s.Equal(tema.Descricao, "descricao")
	s.Equal(tema.Link, "link")
	s.Equal(tema.Foto, 1)
}

func (s *obraTestSuite) TestNewObraFailure() {

	tema := NewObra("titulo", "foto_nao_existe", "2025-02-02",
		"descricao",
		"tema", "link", 1)

	s.Nil(tema)

	tema = NewObra("titulo", "foto", "2025-02-02",
		"descricao",
		"tema_nao_existe", "link", 1)
	s.Nil(tema)
}

func (s *obraTestSuite) TearDownTest() {
	Db = s.original
}

func TestObraSuite(t *testing.T) {
	suite.Run(t, new(obraTestSuite))
}
