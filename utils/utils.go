package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	gim "github.com/ozankasikci/go-image-merge"
)

func WildCardToStringSlice(wildCard string, delim string) ([]string, error) {
	arr := strings.Split(wildCard, delim)

	if len(arr) == 1 && delim == "x" {
		return nil, errors.New("wildcard slice is empty and/or issue with syntax of url")
	}

	return arr, nil
}

var logonames []string

func GetDirNames(path string) ([]string, error) {
	if len(logonames) == 0 {
		fmt.Println("expensive")
		filename, err := os.ReadDir(path)

		if err != nil {
			return nil, err
		}

		for _, e := range filename {
			logonames = append(logonames, strings.Split(e.Name(), ".")[0])
		}

		return logonames, nil
	}

	fmt.Println("cheap")
	return logonames, nil
}

func CreateGrid(optionsArr []string, filetype string) ([]*gim.Grid, error) {
	pathStart := "./images/"
	filenames, err := GetDirNames("./images/")

	if err != nil {
		return nil, err
	}

	grids := []*gim.Grid{}

	for _, optionString := range optionsArr {
		if optionString == "" {
			continue
		}

		if optionString == "js" {
			optionString = "javascript"
		}
		stackLogo := fuzzy.Find(optionString, filenames)[0]

		newI := gim.Grid{}

		if len(stackLogo) == 0 {
			newI.ImageFilePath = pathStart + "404.png"
		} else {

			newI.ImageFilePath = pathStart + stackLogo + filetype
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
