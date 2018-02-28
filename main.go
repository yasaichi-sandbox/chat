package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/yasaichi-sandbox/trace"
	"log"
	"net/http"
	"os"
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

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join(
			"templates",
			t.filename,
		)))
	})

	t.templ.Execute(w, r)
}

func main() {
	addr := flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()

	// TODO: Setup the other providers: Facebook and GitHub
	gomniauth.SetSecurityKey(os.Getenv("OMNIAUTH_SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			fmt.Sprintf("http://localhost%s/auth/callback/google", *addr),
		),
	)

	room := newRoom()
	room.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", room)

	go room.run()

	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
