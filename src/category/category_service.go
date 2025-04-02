package category

// CategoryService defines the interface for category operations
type CategoryService interface {
	CreateCategory(request CategoryRequest) (*CategoryDetailResponse, error)
	GetCategory(id uint) (*CategoryDetailResponse, error)
	GetCategories(p uint, limit uint, sortBy string, direction string, categoryName string) (*CategoryListResponse, error)
	UpdateCategory(id uint, request CategoryRequest) (*CategoryDetailResponse, error)
	DeleteCategory(id uint) error
}

type categoryServiceImpl struct{}

// NewCategoryService creates a new instance of CategoryService
func NewCategoryService() CategoryService {
	return &categoryServiceImpl{}
}

// CreateCategory creates a new category
func (s *categoryServiceImpl) CreateCategory(request CategoryRequest) (*CategoryDetailResponse, error) {
	category := Category{
		Code:        request.Code,
		Name:        request.Name,
		Description: request.Description,
	}

	if err := category.Create(); err != nil {
		return nil, err
	}

	return &CategoryDetailResponse{
		ID:          category.ID,
		Code:        category.Code,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// GetCategory retrieves a category by ID
func (s *categoryServiceImpl) GetCategory(id uint) (*CategoryDetailResponse, error) {
	category, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	return &CategoryDetailResponse{
		ID:          category.ID,
		Code:        category.Code,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// GetCategories retrieves a list of categories with pagination
func (s *categoryServiceImpl) GetCategories(p uint, limit uint, sortBy string, direction string, categoryName string) (*CategoryListResponse, error) {
	categories, page, total, err := FindAll(p, limit, sortBy, direction, categoryName)
	if err != nil {
		return nil, err
	}

	var categoryDTOs []CategoryDetailResponse
	for _, category := range categories {
		dto := CategoryDetailResponse{
			ID:          category.ID,
			Code:        category.Code,
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt,
			UpdatedAt:   category.UpdatedAt,
		}
		categoryDTOs = append(categoryDTOs, dto)
	}

	return &CategoryListResponse{
		Categories: categoryDTOs,
		Page:      page,
		Total:     total,
	}, nil
}

// UpdateCategory updates a category by ID
func (s *categoryServiceImpl) UpdateCategory(id uint, request CategoryRequest) (*CategoryDetailResponse, error) {
	category, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Code = request.Code
	category.Name = request.Name
	category.Description = request.Description

	if err := category.Update(); err != nil {
		return nil, err
	}

	return &CategoryDetailResponse{
		ID:          category.ID,
		Code:        category.Code,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// DeleteCategory soft delete a category by ID
func (s *categoryServiceImpl) DeleteCategory(id uint) error {
	category, err := FindByID(id)
	if err != nil {
		return err
	}

	if err := category.SoftDelete(); err != nil {
		return err
	}

	return nil
}
