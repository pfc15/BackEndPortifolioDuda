package persistence

import (
	"os"
	"path/filepath"
)

type Foto_sql struct {
	Titulo    string  `sql:"Titulo"`
	File_name string  `sql:"path_foto"`
	Descricao string  `sql:"Descricao"`
	Posx      float64 `sql:"posx"`
	Posy      float64 `sql:"posy"`
	Zoom      float64 `sql:"zoom"`
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
			"UPDATE foto SET Titulo =?, path_foto=?, Descricao=?, posx=?, posy=?, zoom=? WHERE foto.id=?;", f.Titulo, f.File_name, f.Descricao,
			f.Posx, f.Posy, f.Zoom, isFotoNew); err != nil {
			return nil, err
		}
	} else {
		if _, err = Db.Exec(
			"INSERT INTO Foto(Titulo, path_foto, Descricao,  posx, posy, zoom) VALUES (?, ?, ?, ?, ?, ?);", f.Titulo, f.File_name, f.Descricao,
			f.Posx, f.Posy, f.Zoom); err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func (f *Foto_sql) UpdateInfo(nova_descrico string, novo_posx float64, novo_posy float64, novo_zoom float64) error {
	_, err := Db.Exec("UPDATE foto SET descricao=?, posx=?, posy=?, zoom=? WHERE foto.titulo=?;", nova_descrico,
		novo_posx, novo_posy, novo_zoom, f.Titulo)
	return err
}

func (f *Foto_sql) UpdateDescricao(nova_descrico string) error {
	_, err := Db.Exec("UPDATE foto SET descricao=? WHERE foto.titulo=?;", nova_descrico, f.Titulo)
	return err
}

func (d *DataBase) GetFotoByTitulo(titulo string) (foto Foto_sql, err error) {
	err = d.db.QueryRow("SELECT Titulo, Descricao, path_foto, posx, posy, zoom FROM Foto WHERE Foto.Titulo=?", titulo).Scan(
		&foto.Titulo, &foto.Descricao, &foto.File_name, &foto.Posx, &foto.Posy, &foto.Zoom)
	if err != nil {
		return foto, err
	}
	return foto, nil
}

func (d *DataBase) GetFotoById(id int) (foto Foto_sql, err error) {

	err = d.db.QueryRow("SELECT Titulo, Descricao, path_foto, posx, posy, zoom FROM Foto WHERE Foto.id=?;", id).Scan(
		&foto.Titulo, &foto.Descricao, &foto.File_name, &foto.Posx, &foto.Posy, &foto.Zoom)
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
