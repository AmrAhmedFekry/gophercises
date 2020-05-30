package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/AmrAhmedFekry/urlshort/urlshort"
)

func main() {
	flagYamlFileName := flag.String("yml", "urls.yaml", "Yaml file path")
	flagJsonFileName := flag.String("json", "urls.json", "Json file path")

	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	ymlFile, err := os.Open(*flagYamlFileName)
	if err != nil {
		fmt.Printf("failed to open %q: %v\n", *flagYamlFileName, err)
		return
	}
	defer ymlFile.Close()
	yaml, err := ioutil.ReadAll(ymlFile)
	if err != nil {
		fmt.Printf("failed to read %q: %v\n", *flagYamlFileName, err)
		return
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonFile, err := os.Open(*flagJsonFileName)
	if err != nil {
		fmt.Printf("failed to open %q: %v\n", *flagJsonFileName, err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read %q: %v\n", *flagJsonFileName, err)
		return
	}
	jsonHandler, err := urlshort.JSONHandler([]byte(jsonData), yamlHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Starting the server on :8000")
	http.ListenAndServe(":3000", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
