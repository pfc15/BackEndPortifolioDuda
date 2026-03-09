package persistence

import (
	"os"
	"path/filepath"
)

type Foto_sql struct {
	Titulo    string `sql:"Titulo"`
	File_name string `sql:"path_foto"`
	Descricao string `sql:"Descricao"`
}

func (f *Foto_sql) Insert() (*os.File, error) {
	// Create an images directory if it doesn’t exist
	if _, err := os.Stat("static/images"); os.IsNotExist(err) {
		os.Mkdir("static/images", 0755)
	}

	// Build the file path and create it
	dst, err := os.Create(filepath.Join("static/images", f.File_name))
	if err != nil {
		return nil, err
	}

	isFotoNew, _ := Db.GetFotoID(f.Titulo)
	if isFotoNew != -1 {
		if _, err = Db.Exec(
			"UPDATE foto SET Titulo =?, path_foto=?, Descricao=? WHERE foto.id=?;", f.Titulo, f.File_name, f.Descricao, isFotoNew); err != nil {
			return nil, err
		}
	} else {
		if _, err = Db.Exec(
			"INSERT INTO Foto(Titulo, path_foto, Descricao) VALUES (?, ?, ?);", f.Titulo, f.File_name, f.Descricao); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func (d *DataBase) GetFotoByTitulo(titulo string) (foto Foto_sql, err error) {
	err = d.db.QueryRow("SELECT Titulo, Descricao, path_foto FROM Foto WHERE Foto.Titulo=?", titulo).Scan(
		&foto.Titulo, &foto.Descricao, &foto.File_name)
	if err != nil {
		return foto, err
	}
	return foto, nil
}

func (d *DataBase) GetFotoById(id int) (foto Foto_sql, err error) {

	err = d.db.QueryRow("SELECT Titulo, Descricao, path_foto FROM Foto WHERE Foto.id=?;", id).Scan(
		&foto.Titulo, &foto.Descricao, &foto.File_name)
	if err != nil {
		return foto, err
	}
	return foto, nil
}

func (d *DataBase) GetFotoID(titulo string) (id int, err error) {
	err = d.db.QueryRow("SELECT id FROM Foto WHERE Foto.Titulo=?;", titulo).Scan(
		&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (d *DataBase) DeleteFoto(id int) error {
	var path_foto string
	err := d.db.QueryRow("SELECT path_foto FROM Foto WHERE Foto.id=?;", id).Scan(
		&path_foto)
	if err != nil {
		return err
	}

	if _, err = Db.Exec("DELETE FROM Foto WHERE id=?;", id); err != nil {
		return err
	}

	err = os.Remove("./static/images/" + path_foto)
	if err != nil {
		return err
	}
	return nil
}
