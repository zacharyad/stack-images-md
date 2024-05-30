package server

import (
	"bytes"
	"encoding/json"
	"image/png"
	"log"
	"net/http"

	gim "github.com/ozankasikci/go-image-merge"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)

	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/images", s.getImages)

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

	// search database for related words, if found they with relate to a directory on server locally.
	pathStart := "/Users/zach-engineering/Documents/Coding/tutorials/golang/projects/stack-images-md/images/"
	//
	grids := []*gim.Grid{
		{ImageFilePath: pathStart + "golang.png"},
		{ImageFilePath: pathStart + "react.png"},
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
