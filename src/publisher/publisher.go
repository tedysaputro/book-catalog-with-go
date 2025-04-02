package publisher

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

// SetDB sets the database connection for the Publisher model
func SetDB(database *gorm.DB) {
	db = database
}

// Publisher represents the publisher table in the database
type Publisher struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for Publisher model
func (Publisher) TableName() string {
	return "publishers"
}

// BeforeCreate hook is called before creating a new record
func (p *Publisher) BeforeCreate(tx *gorm.DB) error {
	if err := p.Validate(); err != nil {
		return err
	}
	return nil
}

// Create saves a new Publisher record to the database
func (p *Publisher) Create() error {
	return db.Create(p).Error
}

// Update updates a Publisher record in the database
func (p *Publisher) Update() error {
	if err := p.Validate(); err != nil {
		return err
	}
	return db.Save(p).Error
}

// FindByID retrieves a Publisher by ID while deleted_at is null
func FindByID(id uint) (*Publisher, error) {
	var publisher Publisher
	err := db.First(&publisher, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("publisher not found")
		}
		return nil, err
	}
	return &publisher, nil
}

// FindAll retrieves all Publishers while deleted_at is null
func FindAll(p uint, limit uint, sortBy string, direction string, publisherName string) ([]Publisher, uint, uint64, error) {
	var publishers []Publisher
	var total int64

	// Get total count
	if err := db.Model(&Publisher{}).Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(publisherName)+"%").Count(&total).Error; err != nil {
		return nil, p, 0, err
	}

	// Calculate offset
	offset := (p - 1) * limit

	// Get records with pagination
	query := db.Model(&Publisher{})
	if sortBy != "" && direction != "" {
		query = query.Order(sortBy + " " + direction)
	}

	err := query.Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(publisherName)+"%").Offset(int(offset)).Limit(int(limit)).Find(&publishers).Error
	if err != nil {
		return nil, p, 0, err
	}

	//Get total pages of publisher
	pages := uint64(total) / uint64(limit)
	if uint64(total)%uint64(limit) > 0 {
		pages++
	}

	return publishers, p, uint64(total), nil
}

// Delete soft deletes a Publisher record
func (p *Publisher) SoftDelete() error {
	return db.Delete(p).Error
}

// Validate checks if the Publisher data is valid
func (p *Publisher) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
