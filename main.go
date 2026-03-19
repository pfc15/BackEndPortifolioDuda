package main

import (
	"PortifolioDuda/internal/connection"
	"PortifolioDuda/internal/persistence"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	persistence.StartDataBase()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := "static" + r.URL.Path
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}

		// Check if file exists
		if info, err := os.Stat(path); err == nil && !info.IsDir() {
			http.ServeFile(w, r, path)
			return
		}

		// Fallback to index.html (React Router)
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/api/getObras", connection.GetObras)
	http.HandleFunc("/api/addObra", connection.AddObra)
	http.HandleFunc("/api/updateObra", connection.UpdateObra)
	http.HandleFunc("/api/deleteObra", connection.DeleteObra)

	http.HandleFunc("/api/getTemas", connection.GetTemas)
	http.HandleFunc("/api/addTema", connection.AddTema)
	http.HandleFunc("/api/updateTema", connection.UpdateTema)
	http.HandleFunc("/api/deleteTema", connection.DeleteTema)

	http.HandleFunc("api/getInfoFoto", connection.GetFoto)
	http.HandleFunc("/api/upload", connection.FileUploadHandler)
	http.HandleFunc("/api/deleteFoto", connection.DeleteFoto)
	http.HandleFunc("/api/updateDescricaoFoto", connection.UpdateDescricao)
	http.HandleFunc("/api/olaMundo", connection.OlaMundoHandler)

	log.Println("Listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
