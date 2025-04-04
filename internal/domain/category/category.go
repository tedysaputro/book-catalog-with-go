package category

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

// SetDB sets the database instance
func SetDB(database *gorm.DB) {
	db = database
}

// Category represents a book category
type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Code        string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"code"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:varchar(500)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Create inserts a new Category record
func (c *Category) Create() error {
	if err := c.Validate(); err != nil {
		return err
	}
	return db.Create(c).Error
}

// Update modifies an existing Category record
func (c *Category) Update() error {
	if err := c.Validate(); err != nil {
		return err
	}
	return db.Save(c).Error
}

// FindByID retrieves a Category by ID while deleted_at is null
func FindByID(id uint) (*Category, error) {
	var category Category
	err := db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

// FindAll retrieves all Categories while deleted_at is null
func FindAll(p uint, limit uint, sortBy string, direction string, categoryName string) ([]Category, uint, uint64, error) {
	var categories []Category
	var total int64

	// Calculate offset
	offset := (p - 1) * limit

	// Count total records
	query := db
	if categoryName != "" {
		query = query.Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(categoryName)+"%")
	}
	err := query.Model(&Category{}).Count(&total).Error
	if err != nil {
		return nil, p, 0, err
	}

	// Get records with pagination
	query = db.Order(sortBy + " " + direction)
	err = query.Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(categoryName)+"%").Offset(int(offset)).Limit(int(limit)).Find(&categories).Error
	if err != nil {
		return nil, p, 0, err
	}

	//Get total pages of category
	pages := uint64(total) / uint64(limit)
	if uint64(total)%uint64(limit) > 0 {
		pages++
	}

	return categories, p, uint64(total), nil
}

// SoftDelete performs a soft delete on the Category record
func (c *Category) SoftDelete() error {
	return db.Delete(c).Error
}

// Validate checks if the Category data is valid
func (c *Category) Validate() error {
	if c.Code == "" {
		return errors.New("code is required")
	}
	if c.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
