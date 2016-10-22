package main

import (
	"time"

	"github.com/blevesearch/bleve/index"
)

// START1 OMIT
type Page struct {
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	Section      string    `json:"section"`
	Content      string    `json:"content"`
	WordCount    float64   `json:"word_count"`
	ReadingTime  float64   `json:"reading_time"`
	Keywords     []string  `json:"keywords"`
	Date         time.Time `json:"date"`
	LastModified time.Time `json:"last_modified"`
	Author       string    `json:"author"`
}

// END1 OMIT

func loop() {
	// START2 OMIT
	for _, p := range site.Pages {
		err = index.Index(p.RelPermalink(), p)
		if err != nil {
			// handle error
		}
	}
	// END2 OMIT
}
