package main

import (
	"log"

	"github.com/blevesearch/bleve"
)

func main() {
	// START OMIT

	shard1, err := bleve.Open("shard1.bleve")
	if err != nil {
		log.Fatal(err)
	}

	shard2, err := bleve.Open("shard2.bleve")
	if err != nil {
		log.Fatal(err)
	}

	index := bleve.NewIndexAlias(shard1, shard2)

	q := bleve.NewMatchQuery("application query")
	req := bleve.NewSearchRequest(q)
	res, err := index.Search(req)
	if err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
