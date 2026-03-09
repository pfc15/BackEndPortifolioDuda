package persistence

import "fmt"

type Obra struct {
	Foto      int    `sql:"foto"`
	Titulo    string `sql:"titulo"`
	Periodo   string `sql:"periodo"`
	Descricao string `sql:"descricao"`
	Ordem     int    `sql:"ordem"`
	Tema      int    `sql:"tema"`
	Link      string `sql:"link"`
}

func NewObra(titulo, foto, Periodo, descricao, tema, link string, ordem int) *Obra {
	tema_id := Db.GetTemaIdByTitulo(tema)
	if tema_id == -1 {
		return nil
	}
	foto_id, err := Db.GetFotoID(foto)
	if err != nil {
		return nil
	}

	return &Obra{
		Titulo:    titulo,
		Foto:      foto_id,
		Periodo:   Periodo,
		Descricao: descricao,
		Ordem:     ordem,
		Tema:      tema_id,
		Link:      link,
	}
}

func (o *Obra) Insert() error {
	id := Db.GetObraIdByTitulo(o.Titulo)
	if id == -1 {
		if o.Link == "" {
			if _, err := Db.Exec("INSERT INTO Obra(titulo, foto, ordem, periodo, descricao, tema) VALUES (?,?,?,?,?,?);",
				o.Titulo, o.Foto, o.Ordem, o.Periodo, o.Descricao, o.Tema); err != nil {
				return err
			}
		} else {
			if _, err := Db.Exec("INSERT INTO Obra(titulo, foto, ordem, periodo, descricao, tema, link) VALUES (?,?,?,?,?,?,?);",
				o.Titulo, o.Foto, o.Ordem, o.Periodo, o.Descricao, o.Tema, o.Link); err != nil {
				return err
			}
		}
	} else {
		if o.Link == "" {
			if _, err := Db.Exec("UPDATE obra SET titulo=?, foto=?, ordem=?, periodo=?, descricao=?, tema=?, link=? WHERE id=?;",
				o.Titulo, o.Foto, o.Ordem, o.Periodo, o.Descricao, o.Tema, nil, id); err != nil {
				return err
			}
		} else {
			if _, err := Db.Exec("UPDATE obra SET titulo=?, foto=?, ordem=?, periodo=?, descricao=?, tema=?, link=?"+
				" WHERE id=?;",
				o.Titulo, o.Foto, o.Ordem, o.Periodo, o.Descricao, o.Tema, o.Link, id); err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *DataBase) GetObrasByTema(temaID int) ([]Obra, error) {

	rows, err := d.db.Query(
		`SELECT titulo, foto, periodo, descricao, ordem, link 
		 FROM Obra 
		 WHERE Obra.tema = ?;`,
		temaID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var obras []Obra

	for rows.Next() {
		var o Obra

		err = rows.Scan(
			&o.Titulo,
			&o.Foto,
			&o.Periodo,
			&o.Descricao,
			&o.Ordem,
			&o.Link,
		)
		if err != nil {
			return nil, err
		}

		obras = append(obras, o)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return obras, nil
}

func (d *DataBase) GetObraIdByTitulo(titulo string) int {
	var id int
	if err := d.db.QueryRow("SELECT id FROM obra WHERE obra.titulo=?;", titulo).
		Scan(&id); err != nil {
		return -1
	}
	return id
}

func (d *DataBase) DeleteObra(titulo string) error {
	id := Db.GetObraIdByTitulo(titulo)
	if id == -1 {
		return fmt.Errorf("Obra does not exist")
	} else {
		var foto_id int
		err := d.db.QueryRow("SELECT foto FROM obra WHERE obra.id=?;", id).Scan(&foto_id)
		if err != nil {
			return err
		}

		err = Db.DeleteFoto(foto_id)
		if err != nil {
			return err
		}

		if _, err = Db.Exec("DELETE FROM obra WHERE id=?;", id); err != nil {
			return err
		}
	}
	return nil
}
