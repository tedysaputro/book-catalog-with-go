package book

import (
	"errors"
	"strings"
	"time"

	"github.com/tedysaputro/book-catalog-with-go/src/author"
	"github.com/tedysaputro/book-catalog-with-go/src/publisher"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetDB sets the database instance
func SetDB(database *gorm.DB) {
	db = database
}

// Book represents a book in the catalog
type Book struct {
	ID          uint                `gorm:"primaryKey" json:"id"`
	Title       string              `gorm:"type:varchar(200);not null" json:"title"`
	Description string              `gorm:"type:varchar(1000)" json:"description"`
	Pages       uint                `gorm:"not null" json:"pages"`
	Year        uint                `gorm:"not null" json:"year"`
	PublisherID uint                `gorm:"not null" json:"publisher_id"`
	Publisher   publisher.Publisher `gorm:"foreignKey:PublisherID" json:"publisher"`
	Authors     []author.Author     `gorm:"many2many:book_authors;" json:"authors"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `gorm:"index" json:"deleted_at,omitempty"`
}

// BookQuery represents a projection of book data with publisher and author information
type BookQuery struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	PublisherID   uint   `json:"publisher_id"`
	PublisherName string `json:"publisher_name"`
}

// Create inserts a new Book record
func (b *Book) Create() error {
	if err := b.Validate(); err != nil {
		return err
	}
	return db.Create(b).Error
}

// Update modifies an existing Book record
func (b *Book) Update() error {
	if err := b.Validate(); err != nil {
		return err
	}

	// Start a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Update book details
	if err := tx.Save(b).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update authors relationship
	if err := tx.Model(b).Association("Authors").Replace(b.Authors); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// FindByID retrieves a Book by ID while deleted_at is null
func FindByID(id uint) (*Book, error) {
	var book Book
	err := db.Preload("Publisher").Preload("Authors").First(&book, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

// FindAll retrieves all Books while deleted_at is null
func FindAll(p uint, limit uint, sortBy string, direction string, title string) ([]Book, uint, uint64, error) {
	var books []Book
	var total int64

	// Calculate offset
	offset := (p - 1) * limit

	// Count total records
	query := db
	if title != "" {
		query = query.Where("UPPER(title) LIKE ?", "%"+strings.ToUpper(title)+"%")
	}
	err := query.Model(&Book{}).Count(&total).Error
	if err != nil {
		return nil, p, 0, err
	}

	// Get records with pagination
	query = db.Preload("Publisher").Preload("Authors").Order(sortBy + " " + direction)
	err = query.Where("UPPER(title) LIKE ?", "%"+strings.ToUpper(title)+"%").Offset(int(offset)).Limit(int(limit)).Find(&books).Error
	if err != nil {
		return nil, p, 0, err
	}

	//Get total pages of books
	pages := uint64(total) / uint64(limit)
	if uint64(total)%uint64(limit) > 0 {
		pages++
	}

	return books, p, uint64(total), nil
}

// SoftDelete performs a soft delete on the Book record
func (b *Book) SoftDelete() error {
	return db.Delete(b).Error
}

// AddAuthors adds authors to the book
func (b *Book) AddAuthors(authorIDs []uint) error {
	var authors []author.Author
	if err := db.Find(&authors, authorIDs).Error; err != nil {
		return err
	}

	if len(authors) != len(authorIDs) {
		return errors.New("one or more authors not found")
	}

	return db.Model(b).Association("Authors").Append(authors)
}

// RemoveAuthors removes authors from the book
func (b *Book) RemoveAuthors(authorIDs []uint) error {
	var authors []author.Author
	if err := db.Find(&authors, authorIDs).Error; err != nil {
		return err
	}

	return db.Model(b).Association("Authors").Delete(authors)
}

// Validate checks if the Book data is valid
func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Pages == 0 {
		return errors.New("pages must be greater than 0")
	}
	if b.Year == 0 {
		return errors.New("year is required")
	}
	if b.PublisherID == 0 {
		return errors.New("publisher is required")
	}

	// Validate that publisher exists
	var count int64
	if err := db.Model(&publisher.Publisher{}).Where("id = ?", b.PublisherID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return errors.New("publisher not found")
	}

	return nil
}
