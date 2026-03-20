package connection

import (
	"PortifolioDuda/internal/persistence"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddObra(w http.ResponseWriter, r *http.Request) {
	type obraPayload struct {
		Foto      string `json:"foto"`
		Titulo    string `json:"titulo"`
		Data      string `json:"data"`
		Descricao string `json:"descricao"`
		Tema      string `json:"tema"`
		Link      string `json:"link"`
	}
	var p obraPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	obra_sql := persistence.NewObra(
		p.Titulo, p.Foto, p.Data, p.Descricao, p.Tema, p.Link)
	if obra_sql == nil {
		log.Println("não encontrou tema ou foto")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = obra_sql.Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func UpdateObra(w http.ResponseWriter, r *http.Request) {
	type obraPayload struct {
		TituloAntigo string `json:"titulo_antigo"`
		Foto         string `json:"foto"`
		Titulo       string `json:"titulo"`
		Data         string `json:"data"`
		Descricao    string `json:"descricao"`
		Tema         string `json:"tema"`
		Link         string `json:"link"`
	}
	var p obraPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	obra_sql := persistence.NewObra(
		p.TituloAntigo, p.Foto, p.Data, p.Descricao, p.Tema, p.Link)
	if obra_sql == nil {
		log.Println("não encontrou tema ou foto")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = obra_sql.Update(p.Titulo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func DeleteObra(w http.ResponseWriter, r *http.Request) {
	type delObraPayload struct {
		Titulo string `json:"titulo"`
	}
	var p delObraPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = persistence.Db.DeleteObra(p.Titulo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetObras(w http.ResponseWriter, r *http.Request) {
	type gobraPayload struct {
		Tema string `json:"tema"`
	}
	var p gobraPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tema_id := persistence.Db.GetTemaIdByTitulo(p.Tema)
	obras, err := persistence.Db.GetObrasByTema(tema_id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fotos := []persistence.Foto_sql{}
	for _, o := range obras {
		var f persistence.Foto_sql
		f, err = persistence.Db.GetFotoById(o.Foto)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fotos = append(fotos, f)
	}

	response := struct {
		Obras []persistence.Obra     `json:"obras"`
		Fotos []persistence.Foto_sql `json:"fotos"`
	}{
		Obras: obras,
		Fotos: fotos,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func AddTema(w http.ResponseWriter, r *http.Request) {
	type temaPayload struct {
		Titulo  string `json:"titulo"`
		Foto    string `json:"foto"`
		Periodo string `json:"periodo"`
	}
	var p temaPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t_sql := persistence.NewTema_sql(
		p.Titulo,
		p.Foto,
	)
	err := t_sql.Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
}

func GetTemas(w http.ResponseWriter, r *http.Request) {
	temas := persistence.Db.GetallTemas()
	fotos := []persistence.Foto_sql{}
	w.Header().Set("Content-Type", "application/json")
	for _, t := range temas {
		f, err := persistence.Db.GetFotoById(t.Foto)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fotos = append(fotos, f)
	}
	response := struct {
		Temas []persistence.Tema_sql `json:"temas"`
		Fotos []persistence.Foto_sql `json:"fotos"`
	}{
		Temas: temas,
		Fotos: fotos,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func UpdateTema(w http.ResponseWriter, r *http.Request) {
	type temaPayload struct {
		TituloNovo string `json:"tituloNovo"`
		Titulo     string `json:"titulo"`
		Foto       string `json:"foto"`
		Periodo    string `json:"periodo"`
	}
	var p temaPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t_sql := persistence.NewTema_sql(
		p.Titulo,
		p.Foto,
	)
	if t_sql == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := t_sql.Update(p.TituloNovo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
}

func DeleteTema(w http.ResponseWriter, r *http.Request) {
	type temaPayload struct {
		Tema string `json:"tema"`
	}
	var p temaPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := persistence.Db.DeleteTema(p.Tema)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("deleted tema", p.Tema)

}

func UpdateDescricao(w http.ResponseWriter, r *http.Request) {
	type fotoPayload struct {
		Descricao string `json:"descricao"`
		Titulo    string `json:"titulo_foto"`
	}
	var p fotoPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	foto_sql := persistence.Foto_sql{
		Titulo:    p.Titulo,
		Descricao: p.Descricao,
	}
	err := foto_sql.UpdateDescricao(p.Descricao)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	// Limit file size to 10MB. This line saves you from those accidental 100MB images!
	r.ParseMultipartForm(10 << 20)

	// Retrieve the file from form data
	file, handler, err := r.FormFile("arquivo")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "Uploaded File: %s\n", handler.Filename)
	fmt.Fprintf(w, "File Size: %d\n", handler.Size)
	fmt.Fprintf(w, "MIME Header: %v\n", handler.Header)

	obj_foto := persistence.Foto_sql{
		Titulo:    r.FormValue("Titulo"),
		File_name: r.FormValue("File_name"),
		Descricao: r.FormValue("Descricao"),
	}

	// Now let’s save it locally
	dst, err := obj_foto.Insert()
	defer dst.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the destination file
	if _, err = dst.ReadFrom(file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
	}
}

func GetFoto(w http.ResponseWriter, r *http.Request) {
	foto, err := persistence.Db.GetFotoByTitulo(r.FormValue("Titulo"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error retrieving the foto", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(foto)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error retrieving the foto", http.StatusInternalServerError)
		return
	}
}

func DeleteFoto(w http.ResponseWriter, r *http.Request) {
	type temaPayload struct {
		Titulo string `json:"titulo"`
	}
	var p temaPayload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := persistence.Db.GetFotoID(p.Titulo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = persistence.Db.DeleteFoto(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func OlaMundoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola Mundo!"))
}
