package control

import (
	"embed"
	"fmt"
	"html"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	//go:embed all:ui/out
	embeddedOutput embed.FS
)

func NewUIServer(port string) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	rootEmbeddedOutput, err := fs.Sub(embeddedOutput, "ui/out")
	if err != nil {
		log.Fatal(err)
	}
	nextFs := http.FileServer(http.FS(rootEmbeddedOutput))
	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		ext := filepath.Ext(path)
		// When path has no extension, append the .html extension before serving
		if ext == "" && path != "/" {
			r.URL.Path = fmt.Sprintf("%s.%s", path, "html")
		}

		nextFs.ServeHTTP(w, r)
	})

	return &Server{
		router: r,
		port:   port,
	}
}

func NewAPIServer(port string) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return &Server{
		router: r,
		port:   port,
	}
}

type Server struct {
	router *chi.Mux
	port   string
}

func (s *Server) ListenAndServe() error {
	log.Printf("serving on port 0.0.0.0:%s", s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}
