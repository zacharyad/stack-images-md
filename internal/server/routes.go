package server

import (
	"bytes"
	"encoding/json"
	gim "github.com/ozankasikci/go-image-merge"
	"image/png"
	"log"
	"net/http"
	"path/filepath"
	util "stack-images-md/utils"
	"strconv"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleHomepage)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	mux.HandleFunc("/images-list", s.getImagesList)
	mux.HandleFunc("/l/{logos}", s.handleGetImages)
	mux.HandleFunc("/l/{gridRowCol}/{logos}", s.handleGetImagesWithOpts)
	return mux
}

func (s *Server) handleHomepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("public", "index.html"))

}

func (s *Server) getImagesList(w http.ResponseWriter, r *http.Request) {
	images, err := util.GetDirNames("./images")

	if err != nil {
		log.Println("Could not get all dir names")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func (s *Server) handleGetImages(w http.ResponseWriter, r *http.Request) {
	optsArr, errGettingOpts := util.WildCardToStringSlice(r.PathValue("logos"), "-")
	if errGettingOpts != nil {
		return
	}

	grids, err := util.CreateGrid(optsArr, ".png")

	if err != nil {
		log.Printf("Issue creating grids slice. Err: %v", err)
	}

	rgba, err := gim.New(grids, len(grids), 1).Merge()

	if err != nil {
		log.Printf("error Creating grid image. Err: %v", err)
	}

	buf := new(bytes.Buffer)

	if err := png.Encode(buf, rgba); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "max-age=3600")

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func (s *Server) handleGetImagesWithOpts(w http.ResponseWriter, r *http.Request) {
	optsArr, errOpts := util.WildCardToStringSlice(r.PathValue("logos"), "-")
	gridRowCol, errRowCol := util.WildCardToStringSlice(r.PathValue("gridRowCol"), "x")

	if errRowCol != nil {

		w.Write([]byte("Please use proper size options. e.g. /3x2/ for a 3 row, 2 column grid. Any whitespace is due to not enough stack images to fit the grid size specified."))
		return
	}

	if errOpts != nil {
		w.Write([]byte("Issue getting logos. Please make sure your url matches the required syntax."))
		return
	}

	cols, errCols := strconv.Atoi(gridRowCol[1])
	rows, errRows := strconv.Atoi(gridRowCol[0])

	if errCols != nil || errRows != nil {
		log.Println("Error converting x and y from grid params")
		return
	}

	grids, err := util.CreateGrid(optsArr, ".png")

	if err != nil {
		log.Printf("Issue creating grids slice. Err: %v", err)

	}

	image, err := gim.New(grids, cols, rows, gim.OptGridSize(512, 512)).Merge()

	if err != nil {
		log.Printf("ERROR Creating grid image. Err: %v", err)
	}

	buf := new(bytes.Buffer)

	if err := png.Encode(buf, image); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")

	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Println("unable to write image.")
	}

}
