package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"

	fuzzy "github.com/lithammer/fuzzysearch/fuzzy"
	gim "github.com/ozankasikci/go-image-merge"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)

	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/images/{options}", s.getImages)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) getImages(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("options")
	optionsArr := strings.Split(params, "-")
	fmt.Println("options from url: ", optionsArr)
	// search database for related words, if found they with relate to a directory on server locally.
	pathStart := "/Users/zach-engineering/Documents/Coding/tutorials/golang/projects/stack-images-md/images/"
	filenames, err := os.ReadDir("./images")

	if err != nil {
		panic("Everything is buring")
	}

	words := []string{}

	for _, word := range filenames {
		words = append(words, word.Name())
	}

	grids := []*gim.Grid{}

	for _, optionString := range optionsArr {
		if optionString == "" {
			fmt.Println("FOUND", optionString)
			continue
		}

		stackLogo := fuzzy.Find(optionString, words)
		newI := gim.Grid{}
		if len(stackLogo) == 0 {
			newI.ImageFilePath = pathStart + "404.png"
		} else {
			newI.ImageFilePath = pathStart + stackLogo[0]
		}

		grids = append(grids, &newI)
	}

	if len(grids) == 0 {
		newI := gim.Grid{}
		newI.ImageFilePath = pathStart + "404.png"
		grids = append(grids, &newI)
	}

	rgba, err := gim.New(grids, len(grids), 1).Merge()

	if err != nil {
		log.Fatalf("error Creating grid image. Err: %v", err)
	}

	if err != nil {
		log.Fatalf("error Creating temp image file. Err: %v", err)
	}

	buffer := new(bytes.Buffer)

	if err := png.Encode(buffer, rgba); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
