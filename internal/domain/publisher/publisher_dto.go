package publisher

type PublisherRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type PublisherCreateResponse struct {
	ID uint `json:"id"`
}

type PublisherDetailResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PublisherListResponse struct {
	Result   []PublisherDTO `json:"result"`
	Pages    uint           `json:"pages"`
	Elements uint64         `json:"elements"`
}

type PublisherDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
