package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Tema_sql struct {
	Titulo string `sql:"Titulo"`
	Foto   int    `sql:"Foto"`
}

func NewTema_sql(titulo string, nome_foto string) *Tema_sql {
	id, err := Db.GetFotoID(nome_foto)
	if errors.Is(err, sql.ErrNoRows) {
		tema, err2 := Db.GetTemaByTitulo(titulo)
		if err2 != nil {
			log.Println(err2)
			return nil
		}
		id = tema.Foto
	} else if err != nil {
		log.Println(err)
		return nil
	}
	return &Tema_sql{
		Titulo: titulo,
		Foto:   id,
	}
}

func (t *Tema_sql) Insert() error {
	id := Db.GetTemaIdByTitulo(t.Titulo)
	if id == -1 {
		if _, err := Db.Exec(
			"INSERT INTO tema(titulo, Foto) VALUES (?,?);", t.Titulo, t.Foto); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("titulo %s already exists", t.Titulo)
	}

	return nil
}

func (t *Tema_sql) Update(titulo_novo string) error {
	id := Db.GetTemaIdByTitulo(t.Titulo)
	if id != -1 {
		if _, err := Db.Exec(
			"UPDATE tema SET titulo=?, Foto=? WHERE tema.id=?;", titulo_novo, t.Foto, id); err != nil {
			return err
		}
		t.Titulo = titulo_novo
	} else {
		return fmt.Errorf("titulo doesn't exists")
	}

	return nil
}

func (d *DataBase) GetallTemas() []Tema_sql {
	retorno := []Tema_sql{}
	rows, err := d.db.Query("SELECT titulo, foto FROM tema;")
	if err != nil {
		log.Println(err)
		return retorno
	}
	defer rows.Close()
	for rows.Next() {
		var t Tema_sql
		err = rows.Scan(&t.Titulo, &t.Foto)
		if err != nil {
			log.Println(err)
			return []Tema_sql{}
		}
		retorno = append(retorno, t)
	}
	return retorno
}

func (d *DataBase) GetTemaIdByTitulo(titulo string) int {
	var id int
	if err := d.db.QueryRow("SELECT id FROM tema WHERE tema.titulo=?;", titulo).
		Scan(&id); err != nil {
		return -1
	}
	return id
}

func (d *DataBase) GetTemaByTitulo(titulo string) (Tema_sql, error) {
	retorno := Tema_sql{}
	row := d.db.QueryRow("SELECT titulo, foto FROM tema WHERE tema.titulo=?;", titulo)
	err := row.Scan(&retorno.Titulo, &retorno.Foto)
	if err != nil {
		log.Println(err)
		return Tema_sql{}, err
	}
	return retorno, nil
}

func (d *DataBase) DeleteTema(titulo string) error {
	id := Db.GetTemaIdByTitulo(titulo)
	if id == -1 {
		return fmt.Errorf("Tema does not exist")
	}
	var foto_id int
	err := d.db.QueryRow("SELECT foto FROM tema WHERE tema.id=?;", id).Scan(&foto_id)
	if err != nil {
		return err
	}

	err = Db.DeleteFoto(foto_id)
	if err != nil {
		return err
	}

	rows, err := d.db.Query("SELECT titulo FROM obra WHERE obra.tema=?;", id)
	if err != nil {
		return err
	}

	var tobras []string
	for rows.Next() {
		var t_obra string
		err = rows.Scan(&t_obra)
		if err != nil {
			return err
		}
		tobras = append(tobras, t_obra)

	}
	rows.Close()
	for _, t := range tobras {
		err = d.DeleteObra(t)
		if err != nil {
			return err
		}
	}

	if _, err = Db.Exec(
		"DELETE FROM tema WHERE tema.id=?;", id); err != nil {
		return err
	}
	return nil
}
