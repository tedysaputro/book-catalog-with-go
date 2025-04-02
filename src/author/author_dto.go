package author

// AuthorRequest represents the request body for creating an author
type AuthorRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type AuthorCreateResponse struct {
	ID uint `json:"id"`
}

type AuthorDetailResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AuthorResponse represents the response for author endpoints
type AuthorListResponse struct {
	Result   []AuthorDTO `json:"result"`
	Pages    uint        `json:"pages"`
	Elements uint64      `json:"elements"`
}

type AuthorDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
