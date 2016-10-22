package main

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

type Person struct {
	Name string
}

// START OMIT
func main() {
	index, err := bleve.Open("wiki.bleve") // HLOPEN
	if err != nil {
		log.Fatal(err)
	}

	query := bleve.NewQueryStringQuery("go program") // HLQUERY
	request := bleve.NewSearchRequest(query)         // HLREQ
	result, err := index.Search(request)             // HLSEARCH
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

// END OMIT
