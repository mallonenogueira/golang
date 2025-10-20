package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type IndexPageData struct {
	Title   string
	Content string
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/2", getRoot2)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		os.Exit(1)
	}
}

// func getGeneral(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Teste2")
// 	io.WriteString(w, "This is my site!2")
// }

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("Teste")
	io.WriteString(w, "This is my site!")
}

func getRoot2(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = tmpl.Execute(w, IndexPageData{"Título", "Conteúdo"})
}
