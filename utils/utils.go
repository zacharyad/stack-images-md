package utils

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	gim "github.com/ozankasikci/go-image-merge"
)

func WildCardToStringSlice(wildCard string, delim string, r *http.Request) ([]string, error) {
	arr := strings.Split(r.PathValue(wildCard), delim)

	if len(arr) == 1 {
		return nil, errors.New("wildcard slice is empty and/or issue with syntax of url")
	}

	return arr, nil
}

func CreateGrid(optionsArr []string) ([]*gim.Grid, error) {
	pathStart := "./images/"
	filenames, err := os.ReadDir("./images")

	if err != nil {
		return nil, err
	}

	words := []string{}

	for _, word := range filenames {
		words = append(words, word.Name())
	}

	grids := []*gim.Grid{}

	for _, optionString := range optionsArr {
		if optionString == "" {
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

	return grids, nil
}
