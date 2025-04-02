package category

import "time"

// CategoryRequest represents the request payload for creating/updating a category
type CategoryRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CategoryDetailResponse represents the response payload for a single category
type CategoryDetailResponse struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CategoryListResponse represents the response payload for multiple categories
type CategoryListResponse struct {
	Categories []CategoryDetailResponse `json:"data"`
	Page      uint                     `json:"page"`
	Total     uint64                   `json:"total"`
}
