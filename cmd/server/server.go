package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/AustinAttempts/goBlog"
)

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdOut io.Writer) error {
	mux := http.NewServeMux()

	postReader := goBlog.FileReader{
		Dir: "posts",
	}
	postTemplate := template.Must(template.ParseFiles("post.gohtml"))
	mux.HandleFunc("GET /posts/{slug}", goBlog.PostHandler(postReader, postTemplate))
	indexTemplate := template.Must(template.ParseFiles("index.gohtml"))
	mux.HandleFunc("GET /", goBlog.IndexHandler(postReader, indexTemplate))

	log.Fatal(http.ListenAndServe(":8080", mux))
	return nil
}
