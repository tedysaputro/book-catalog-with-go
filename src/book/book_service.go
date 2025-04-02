package book

import (
	"errors"
	"fmt"

	"github.com/tedysaputro/book-catalog-with-go/src/author"
	"github.com/tedysaputro/book-catalog-with-go/src/publisher"
)

// BookService defines the interface for book operations
type BookService interface {
	CreateBook(request BookRequest) (*BookCreateResponse, error)
	GetBook(id uint) (*BookDetailResponse, error)
	GetBooks(p uint, limit uint, sortBy string, direction string, title string) (*BookListResponse, error)
	UpdateBook(id uint, request BookRequest) (*BookDetailResponse, error)
	DeleteBook(id uint) error
}

type bookServiceImpl struct{}

// NewBookService creates a new instance of BookService
func NewBookService() BookService {
	return &bookServiceImpl{}
}

// CreateBook creates a new book
func (s *bookServiceImpl) CreateBook(request BookRequest) (*BookCreateResponse, error) {
	// Get authors if author IDs are provided
	var authors []author.Author
	if len(request.AuthorIDs) > 0 {
		if err := db.Find(&authors, request.AuthorIDs).Error; err != nil {
			return nil, err
		}
		if len(authors) != len(request.AuthorIDs) {
			return nil, errors.New("one or more authors not found")
		}
	}

	book := Book{
		Title:       request.Title,
		Description: request.Description,
		Pages:       request.Pages,
		Year:        request.Year,
		PublisherID: request.PublisherID,
		Authors:     authors,
	}

	if err := book.Create(); err != nil {
		return nil, err
	}

	// Fetch the book again to get the publisher and author details
	createdBook, err := FindByID(book.ID)
	if err != nil {
		return nil, err
	}

	return &BookCreateResponse{
		ID: createdBook.ID,
	}, nil
}

// GetBook retrieves a book by ID
func (s *bookServiceImpl) GetBook(id uint) (*BookDetailResponse, error) {
	book, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	// Convert Publisher to PublisherDTO
	publisherDTO := publisher.PublisherDTO{
		ID:   fmt.Sprintf("%d", book.Publisher.ID),
		Name: book.Publisher.Name,
	}

	// Convert Authors to AuthorDTOs
	authorDTOs := make([]author.AuthorDTO, len(book.Authors))
	for i, a := range book.Authors {
		authorDTOs[i] = author.AuthorDTO{
			ID:   fmt.Sprintf("%d", a.ID),
			Name: a.Name,
		}
	}

	return &BookDetailResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Pages:       book.Pages,
		Year:        book.Year,
		Publisher:   publisherDTO,
		Authors:     authorDTOs,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}, nil
}

// GetBooks retrieves a list of books with pagination
func (s *bookServiceImpl) GetBooks(p uint, limit uint, sortBy string, direction string, title string) (*BookListResponse, error) {
	books, page, total, err := FindAll(p, limit, sortBy, direction, title)
	if err != nil {
		return nil, err
	}

	var bookDTOs []BookDetailResponse
	for _, book := range books {
		// Convert Publisher to PublisherDTO
		publisherDTO := publisher.PublisherDTO{
			ID:   fmt.Sprintf("%d", book.Publisher.ID),
			Name: book.Publisher.Name,
		}

		// Convert Authors to AuthorDTOs
		authorDTOs := make([]author.AuthorDTO, len(book.Authors))
		for i, a := range book.Authors {
			authorDTOs[i] = author.AuthorDTO{
				ID:   fmt.Sprintf("%d", a.ID),
				Name: a.Name,
			}
		}

		dto := BookDetailResponse{
			ID:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Pages:       book.Pages,
			Year:        book.Year,
			Publisher:   publisherDTO,
			Authors:     authorDTOs,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		}
		bookDTOs = append(bookDTOs, dto)
	}

	return &BookListResponse{
		Books: bookDTOs,
		Page:  page,
		Total: total,
	}, nil
}

// UpdateBook updates a book by ID
func (s *bookServiceImpl) UpdateBook(id uint, request BookRequest) (*BookDetailResponse, error) {
	book, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	// Get authors if author IDs are provided
	var authors []author.Author
	if len(request.AuthorIDs) > 0 {
		if err := db.Find(&authors, request.AuthorIDs).Error; err != nil {
			return nil, err
		}
		if len(authors) != len(request.AuthorIDs) {
			return nil, errors.New("one or more authors not found")
		}
	}

	book.Title = request.Title
	book.Description = request.Description
	book.Pages = request.Pages
	book.Year = request.Year
	book.PublisherID = request.PublisherID
	book.Authors = authors

	if err := book.Update(); err != nil {
		return nil, err
	}

	// Fetch the book again to get the updated publisher and author details
	updatedBook, err := FindByID(book.ID)
	if err != nil {
		return nil, err
	}

	// Convert Publisher to PublisherDTO
	publisherDTO := publisher.PublisherDTO{
		ID:   fmt.Sprintf("%d", updatedBook.Publisher.ID),
		Name: updatedBook.Publisher.Name,
	}

	// Convert Authors to AuthorDTOs
	authorDTOs := make([]author.AuthorDTO, len(updatedBook.Authors))
	for i, a := range updatedBook.Authors {
		authorDTOs[i] = author.AuthorDTO{
			ID:   fmt.Sprintf("%d", a.ID),
			Name: a.Name,
		}
	}

	return &BookDetailResponse{
		ID:          updatedBook.ID,
		Title:       updatedBook.Title,
		Description: updatedBook.Description,
		Pages:       updatedBook.Pages,
		Year:        updatedBook.Year,
		Publisher:   publisherDTO,
		Authors:     authorDTOs,
		CreatedAt:   updatedBook.CreatedAt,
		UpdatedAt:   updatedBook.UpdatedAt,
	}, nil
}

// DeleteBook soft delete a book by ID
func (s *bookServiceImpl) DeleteBook(id uint) error {
	book, err := FindByID(id)
	if err != nil {
		return err
	}

	if err := book.SoftDelete(); err != nil {
		return err
	}

	return nil
}
