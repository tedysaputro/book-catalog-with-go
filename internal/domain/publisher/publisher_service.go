package publisher

import "strconv"

// PublisherService defines the interface for publisher operations

type PublisherService interface {
	createPublisher(request PublisherRequest) (*PublisherCreateResponse, error)
	GetPublisher(id uint) (*PublisherDetailResponse, error)
	GetPublishers(p uint, limit uint, sortBy string, direction string, publisherName string) (*PublisherListResponse, error)
	UpdatePublisher(id uint, request PublisherRequest) (*PublisherDetailResponse, error)
	DeletePublisher(id uint) error
}

type publisherServiceImpl struct{}

// NewPublisherService creates a new instance of PublisherService
func NewPublisherService() PublisherService {
	return &publisherServiceImpl{}
}

// createPublisher creates a new publisher
func (s *publisherServiceImpl) createPublisher(request PublisherRequest) (*PublisherCreateResponse, error) {
	publisher := &Publisher{
		Name:        request.Name,
		Description: request.Description,
	}

	if err := publisher.Validate(); err != nil {
		return nil, err
	}

	if err := publisher.Create(); err != nil {
		return nil, err
	}

	dto := &PublisherCreateResponse{
		ID: publisher.ID,
	}

	return dto, nil
}

// GetPublisher retrieves a publisher by ID
func (s *publisherServiceImpl) GetPublisher(id uint) (*PublisherDetailResponse, error) {
	publisher, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := &PublisherDetailResponse{
		ID:          publisher.ID,
		Name:        publisher.Name,
		Description: publisher.Description,
	}

	return dto, nil
}

// GetPublishers retrieves all publishers
func (s *publisherServiceImpl) GetPublishers(p uint, limit uint, sortBy string, direction string, publisherName string) (*PublisherListResponse, error) {
	publishers, p, el, err := FindAll(p, limit, sortBy, direction, publisherName)
	if err != nil {
		return nil, err
	}

	dtos := make([]PublisherDTO, len(publishers))
	for i, publisher := range publishers {
		dtos[i] = PublisherDTO{
			ID:   strconv.FormatUint(uint64(publisher.ID), 10),
			Name: publisher.Name,
		}
	}

	return &PublisherListResponse{
		Result:   dtos,
		Pages:    p,
		Elements: el,
	}, nil
}

// UpdatePublisher updates a publisher by ID
func (s *publisherServiceImpl) UpdatePublisher(id uint, request PublisherRequest) (*PublisherDetailResponse, error) {
	publisher, err := FindByID(id)
	if err != nil {
		return nil, err
	}

	publisher.Name = request.Name
	publisher.Description = request.Description

	if err := publisher.Update(); err != nil {
		return nil, err
	}

	dto := &PublisherDetailResponse{
		ID:          publisher.ID,
		Name:        publisher.Name,
		Description: publisher.Description,
	}

	return dto, nil
}

// DeletePublisher soft delete a publisher by ID
func (s *publisherServiceImpl) DeletePublisher(id uint) error {
	publisher, err := FindByID(id)
	if err != nil {
		return err
	}

	if err := publisher.SoftDelete(); err != nil {
		return err
	}

	return nil
}
