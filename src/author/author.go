package author

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

// SetDB sets the database connection for the Author model
func SetDB(database *gorm.DB) {
	db = database
}

// Author represents the author table in the database
type Author struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName specifies the table name for Author model
func (Author) TableName() string {
	return "authors"
}

// BeforeCreate hook is called before creating a new record
func (a *Author) BeforeCreate(tx *gorm.DB) error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

// Create saves a new Author record to the database
func (a *Author) Create() error {
	return db.Create(a).Error
}

// update author record base on id
func (a *Author) Update() error {
	if a.ID == 0 {
		return errors.New("cannot update author without ID")
	}
	return db.Save(a).Error
}

// FindByID retrieves an Author by ID
func FindByID(id uint) (*Author, error) {
	var author Author
	err := db.First(&author, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}
	return &author, nil
}

// FindAll retrieves all Authors
func FindAll(p uint, limit uint, sortBy string, direction string, authorName string) ([]Author, uint, uint64, error) {
	var authors []Author
	err := db.Order(fmt.Sprintf("%s %s", sortBy, direction)).Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(authorName)+"%").Offset(int((p - 1) * limit)).Limit(int(limit)).Find(&authors).Error
	if err != nil {
		return nil, 0, 0, err
	}
	count, err := GetTotalCount(authorName)
	if err != nil {
		return nil, 0, 0, err
	}

	//Get total pages of author
	pages := uint64(count) / uint64(limit)
	if uint64(count)%uint64(limit) > 0 {
		pages++
	}

	return authors, p, uint64(count), nil
}

// GetTotalCount returns the total count of authors
func GetTotalCount(authorName string) (int64, error) {
	var count int64
	err := db.Model(&Author{}).Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(authorName)+"%").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Delete removes an Author record (soft delete)
func (a *Author) Delete() error {
	if a.ID == 0 {
		return errors.New("cannot delete author without ID")
	}
	return db.Delete(a).Error
}

// Validate checks if the Author data is valid
func (a *Author) Validate() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
