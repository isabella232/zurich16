package main

import (
	"context"

	"github.com/blevesearch/bleve/document"
)

// START OMIT
type Index interface {
	Index(id string, data interface{}) error

	Delete(id string) error

	NewBatch() *Batch
	Batch(b *Batch) error

	Document(id string) (*document.Document, error)
	DocCount() (uint64, error)

	Search(req *SearchRequest) (*SearchResult, error)
	SearchInContext(ctx context.Context, req *SearchRequest) (*SearchResult, error)

	Close() error
}

// END OMIT
