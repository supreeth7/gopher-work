package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const port = ":7070"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)

	yamlFile := flag.String("yaml", "", "Yaml file with path:url mapping")
	jsonFile := flag.String("json", "", "JSON file with path:url mapping")

	flag.Parse()

	var (
		data        []byte
		handler     http.HandlerFunc
		pathsToUrls map[string]string
	)

	switch {
	case *yamlFile != "":
		data = getData(*yamlFile)
		handler, _ = yamlHandler(data, mapHandler(pathsToUrls, r))

	case *jsonFile != "":
		data = getData(*jsonFile)
		handler, _ = jsonHandler(data, mapHandler(pathsToUrls, r))

	default:
		// URL:Path map
		pathsToUrls = map[string]string{
			"/github":   "https://github.com/supreeth7",
			"/linkedin": "https://www.linkedin.com/in/supreeth-b",
		}

		handler = mapHandler(pathsToUrls, r)
	}

	fmt.Printf("Starting server at port: %s\n", port)

	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func getData(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}

	return data
}
