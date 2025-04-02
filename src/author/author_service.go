package author

import "strconv"

// AuthorService defines the interface for author operations
type AuthorService interface {
	createAuthor(request AuthorRequest) (*AuthorCreateResponse, error)
	GetAuthor(id uint) (*AuthorDetailResponse, error)
	GetAuthors(p uint, limit uint, sortBy string, direction string, authorName string) (*AuthorListResponse, error)
	UpdateAuthor(id uint, request AuthorRequest) (*AuthorDetailResponse, error)
}

type authorServiceImpl struct{}

// NewAuthorService creates a new instance of AuthorService
func NewAuthorService() AuthorService {
	return &authorServiceImpl{}
}

// createAuthor creates a new author
func (s *authorServiceImpl) createAuthor(request AuthorRequest) (*AuthorCreateResponse, error) {
	author := &Author{
		Name:        request.Name,
		Description: request.Description,
	}

	if err := author.Validate(); err != nil {
		return nil, err
	}

	if err := author.Create(); err != nil {
		return nil, err
	}

	dto := &AuthorCreateResponse{
		ID: author.ID,
	}

	return dto, nil
}

// Update Author by ID
func (s *authorServiceImpl) UpdateAuthor(id uint, request AuthorRequest) (*AuthorDetailResponse, error) {
	author, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	author.Name = request.Name
	author.Description = request.Description

	if err := author.Update(); err != nil {
		return nil, err
	}

	dto := &AuthorDetailResponse{
		ID:          author.ID,
		Name:        author.Name,
		Description: author.Description,
	}

	return dto, nil
}

// GetAuthor retrieves an author by ID
func (s *authorServiceImpl) GetAuthor(id uint) (*AuthorDetailResponse, error) {
	author, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := &AuthorDetailResponse{
		ID:          author.ID,
		Name:        author.Name,
		Description: author.Description,
	}

	return dto, nil
}

// GetAuthors retrieves all authors
func (s *authorServiceImpl) GetAuthors(p uint, limit uint, sortBy string, direction string, authorName string) (*AuthorListResponse, error) {
	authors, p, el, err := FindAll(p, limit, sortBy, direction, authorName)
	if err != nil {
		return nil, err
	}

	dtos := make([]AuthorDTO, len(authors))
	for i, author := range authors {
		dtos[i] = AuthorDTO{
			ID:   strconv.FormatUint(uint64(author.ID), 10),
			Name: author.Name,
		}
	}

	return &AuthorListResponse{
		Result:   dtos,
		Pages:    p,
		Elements: el,
	}, nil
}
