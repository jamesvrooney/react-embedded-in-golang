package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed my-react-app/build
var embeddedFiles embed.FS

func main() {
	fmt.Println("Starting Server")
	http.Handle("/", http.FileServer(getFileSystem()))
	http.HandleFunc("/products", getProducts)

	http.ListenAndServe(":9000", nil)
}

func getFileSystem() http.FileSystem {

	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "my-react-app/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func getProducts(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)

	audi := &Car{
		ID:   6,
		Name: "Audi A3",
	}

	encoder.Encode(audi)
}

type Car struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
