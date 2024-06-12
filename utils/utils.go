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

	if len(arr) == 1 && wildCard == "x" {
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

// http://localhost:8080/js-golang-linux-js-js-golang-linux-js-js-golang-linux-js

// ![Javascript](https://img.icons8.com/color/48/000000/javascript--v1.png)
// ![html](https://img.icons8.com/color/48/000000/html-5--v1.png)
// ![css](https://img.icons8.com/color/48/000000/css3.png)
// ![react](https://img.icons8.com/color/48/000000/react-native.png)
// ![redux](https://img.icons8.com/color/48/000000/redux.png)
// ![node.js](https://img.icons8.com/color/48/000000/nodejs.png)
// ![github](https://img.icons8.com/ios-glyphs/48/000000/github.png)
// ![heroku](https://img.icons8.com/color/48/000000/heroku.png)
// ![postgres](https://img.icons8.com/color/48/000000/postgreesql.png)
// ![webpack](https://img.icons8.com/color/48/000000/webpack.png)

// ![logos for: js-golang-linux-js-js-golang-linux-js-js-golang-linux-js](https://stackimages.xyz/js-golang-linux-js-js-golang-linux-js-js-golang-linux-js)
