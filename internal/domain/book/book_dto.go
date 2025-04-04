package book

import (
	"github.com/tedysaputro/book-catalog-with-go/internal/domain/author"
	"github.com/tedysaputro/book-catalog-with-go/internal/domain/publisher"
)

// BookRequest represents the request payload for creating/updating a book
type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Pages       uint   `json:"pages"`
	Year        uint   `json:"year"`
	PublisherID uint   `json:"publisher_id"`
	AuthorIDs   []uint `json:"author_ids"`
}

type BookCreateResponse struct {
	ID uint `json:"id"`
}

// BookDetailResponse represents the response payload for a single book
type BookDetailResponse struct {
	ID          uint                   `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Pages       uint                   `json:"pages"`
	Year        uint                   `json:"year"`
	Publisher   publisher.PublisherDTO `json:"publisher"`
	Authors     []author.AuthorDTO     `json:"authors"`
}

// BookListResponse represents the response payload for multiple books
type BookListResponse struct {
	Books []BookDetailResponse `json:"data"`
	Page  uint                 `json:"page"`
	Total uint64               `json:"total"`
}
