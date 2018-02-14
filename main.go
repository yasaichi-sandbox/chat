package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	// All fields are private because their names start with lower case
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join(
			"templates",
			t.filename,
		)))
	})

	t.templ.Execute(w, nil)
}

func main() {
	room := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", room)

	go room.run()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
