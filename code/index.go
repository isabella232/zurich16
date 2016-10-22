package main

import (
	"fmt"
	"log"
	"os"
)

// START OMIT
import (
	"github.com/blevesearch/bleve" // HLIMPORT
)

type WikiArticle struct { // HLMODEL
	Title string // HLMODEL
	Body  string // HLMODEL
} // HLMODEL

func main() {
	mapping := bleve.NewIndexMapping()             // HLMAPPING
	index, err := bleve.New("wiki.bleve", mapping) // HLNEW
	if err != nil {
		log.Fatal(err)
	}
	a := WikiArticle{Title: "Go", Body: "Go (often referred to as golang) is a Free and open source[12] programming language created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson."} // HLINDEX
	err = index.Index(a.Title, a)                                                                                                                                                                         // HLINDEX
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Indexed Document")
}

// END OMIT

func init() {
	os.RemoveAll("wiki.bleve")
}
