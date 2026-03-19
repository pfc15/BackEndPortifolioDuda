package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

type DataBaseInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)

	GetFotoById(id int) (Foto_sql, error)
	GetFotoByTitulo(titulo string) (Foto_sql, error)
	GetFotoID(titulo string) (int, error)
	DeleteFoto(id int) error

	GetObrasByTema(temaID int) ([]Obra, error)
	GetObraIdByTitulo(titulo string) int
	DeleteObra(titulo string) error
	GetObraByTitulo(titulo string) (Obra, error)

	GetallTemas() []Tema_sql
	GetTemaIdByTitulo(titulo string) int
	DeleteTema(titulo string) error
	GetTemaByTitulo(titulo string) (Tema_sql, error)
}
type DataBase struct {
	db *sql.DB
}

var Db DataBaseInterface

func StartDataBase() {
	var err error

	exPath := os.Getenv("ROOT")

	conn, err := sql.Open("sqlite", "file:"+exPath+"/db/db.sqlite3?_foreign_keys=on&_busy_timeout=5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Exec("PRAGMA foreign_keys = ON")
	Db = &DataBase{
		db: conn,
	}

	query, err := os.ReadFile(exPath + "/db/create_db.sql")
	if err != nil {
		log.Println("erro lendo arquivo sql")
		panic(err)
	}
	if _, err = Db.Exec(string(query)); err != nil {
		log.Println("erro executando arquivo create sql")
		panic(err)
	}
}

func (d *DataBase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}
