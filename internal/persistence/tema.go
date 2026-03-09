package persistence

import (
	"fmt"
	"log"
)

type Tema_sql struct {
	Titulo  string `sql:"Titulo"`
	Foto    int    `sql:"Foto"`
	Ordem   int    `sql:"ordem"`
	Periodo string `sql:"periodo"`
}

func NewTema_sql(titulo string, nome_foto string, ordem int, periodo string) *Tema_sql {
	id, err := Db.GetFotoID(nome_foto)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Tema_sql{
		Titulo:  titulo,
		Ordem:   ordem,
		Periodo: periodo,
		Foto:    id,
	}
}

func (t *Tema_sql) Insert(updates ...string) error {
	id := Db.GetTemaIdByTitulo(t.Titulo)
	if len(updates) >= 0 {

	}
	if id == -1 {
		if _, err := Db.Exec(
			"INSERT INTO tema(titulo, Foto, ordem, periodo) VALUES (?,?,?,?);", t.Titulo, t.Foto, t.Ordem, t.Periodo); err != nil {
			return err
		}
	} else {
		if _, err := Db.Exec(
			"UPDATE tema SET titulo=?, Foto=?, ordem=?, periodo=? WHERE tema.id=?;", t.Titulo, t.Foto, t.Ordem, t.Periodo, id); err != nil {
			return err
		}
	}

	return nil
}

func (d *DataBase) GetallTemas() []Tema_sql {
	retorno := []Tema_sql{}
	rows, err := d.db.Query("SELECT titulo, foto, ordem, periodo FROM tema;")
	if err != nil {
		log.Println(err)
		return retorno
	}
	defer rows.Close()
	for rows.Next() {
		var t Tema_sql
		err = rows.Scan(&t.Titulo, &t.Foto, &t.Ordem, &t.Periodo)
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
