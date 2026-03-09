package persistence

import (
	"database/sql"
	"fmt"
)

type MockDb struct{}

func (m *MockDb) GetFotoID(titulo string) (int, error) {
	if titulo == "foto_nao_existe" {
		return -1, fmt.Errorf("Foto não encontrada")
	}
	return 1, nil
}
func (m *MockDb) GetFotoById(id int) (Foto_sql, error) {
	return Foto_sql{}, nil
}
func (m *MockDb) GetFotoByTitulo(titulo string) (Foto_sql, error) {
	return Foto_sql{}, nil
}

func (m *MockDb) GetObrasByTema(temaID int) ([]Obra, error) {
	return []Obra{}, nil
}

func (m *MockDb) GetallTemas() []Tema_sql {
	return []Tema_sql{}
}
func (m *MockDb) GetTemaIdByTitulo(titulo string) int {
	if titulo == "tema_nao_existe" {
		return -1
	}
	return 1
}
func (m *MockDb) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (m *MockDb) GetObraIdByTitulo(titulo string) int {
	return 1
}

func (m *MockDb) DeleteTema(titulo string) error {
	return nil
}

func (m *MockDb) DeleteObra(titulo string) error {
	return nil
}

func (m *MockDb) DeleteFoto(id int) error {
	return nil
}
