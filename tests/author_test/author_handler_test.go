package author_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tedysaputro/book-catalog-with-go/src/author"
)

func setupTestDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=book_catalog_test port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrate the schema
	err = db.AutoMigrate(&author.Author{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
}

func cleanupTestDB(db *gorm.DB) {
	db.Exec("DELETE FROM authors")
}

func setupTestApp() (*fiber.App, *gorm.DB) {
	db := setupTestDB()
	author.SetDB(db)

	app := fiber.New()
	authorService := author.NewAuthorService()
	authorHandler := author.NewAuthorHandler(authorService)
	authorHandler.RegisterRoutes(app)
	return app, db
}

func TestCreateAuthor(t *testing.T) {
	app, db := setupTestApp()
	defer cleanupTestDB(db)

	tests := []struct {
		name           string
		payload        author.AuthorRequest
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "Valid Author Creation",
			payload: author.AuthorRequest{
				Name:        "J.K. Rowling",
				Description: "British author, best known for the Harry Potter series",
			},
			expectedStatus: fiber.StatusCreated,
			expectedError:  false,
		},
		{
			name: "Invalid Author - Empty Name",
			payload: author.AuthorRequest{
				Name:        "",
				Description: "Some description",
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/authors", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			if !tt.expectedError {
				var response author.AuthorCreateResponse
				err = json.Unmarshal(body, &response)
				assert.NoError(t, err)
				assert.NotZero(t, response.ID)

				// Get the author details to verify the name
				detailReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/authors/%d", response.ID), nil)
				detailResp, _ := app.Test(detailReq)
				assert.Equal(t, fiber.StatusOK, detailResp.StatusCode)

				detailBody, _ := io.ReadAll(detailResp.Body)
				var detailResponse author.AuthorDetailResponse
				err = json.Unmarshal(detailBody, &detailResponse)
				assert.NoError(t, err)
				assert.Equal(t, tt.payload.Name, detailResponse.Name)
			} else {
				var errorResponse map[string]string
				err = json.Unmarshal(body, &errorResponse)
				assert.NoError(t, err)
				assert.NotNil(t, errorResponse["error"])
			}
		})
	}
}

func TestGetAuthor(t *testing.T) {
	app, db := setupTestApp()
	defer cleanupTestDB(db)

	// First create an author
	createPayload := author.AuthorRequest{
		Name:        "George R.R. Martin",
		Description: "American novelist",
	}
	payload, _ := json.Marshal(createPayload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/authors", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	var createResp author.AuthorCreateResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &createResp)

	tests := []struct {
		name           string
		authorID       uint
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "Get Existing Author",
			authorID:       createResp.ID,
			expectedStatus: fiber.StatusOK,
			expectedError:  false,
		},
		{
			name:           "Get Non-existent Author",
			authorID:       9999,
			expectedStatus: fiber.StatusNotFound,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/authors/%d", tt.authorID), nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if !tt.expectedError {
				body, _ := io.ReadAll(resp.Body)
				var response author.AuthorDetailResponse
				err = json.Unmarshal(body, &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.authorID, response.ID)
				assert.Equal(t, createPayload.Name, response.Name)
			} else {
				var errorResponse map[string]string
				body, _ := io.ReadAll(resp.Body)
				err = json.Unmarshal(body, &errorResponse)
				assert.NoError(t, err)
				assert.NotNil(t, errorResponse["error"])
			}
		})
	}
}

func TestUpdateAuthor(t *testing.T) {
	app, db := setupTestApp()
	defer cleanupTestDB(db)

	// First create an author
	createPayload := author.AuthorRequest{
		Name:        "Original Name",
		Description: "Original Description",
	}
	payload, _ := json.Marshal(createPayload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/authors", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	var createResp author.AuthorCreateResponse
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &createResp)

	tests := []struct {
		name           string
		authorID       uint
		payload        author.AuthorRequest
		expectedStatus int
		expectedError  bool
	}{
		{
			name:     "Valid Author Update",
			authorID: createResp.ID,
			payload: author.AuthorRequest{
				Name:        "Updated Name",
				Description: "Updated Description",
			},
			expectedStatus: fiber.StatusOK,
			expectedError:  false,
		},
		{
			name:     "Update Non-existent Author",
			authorID: 9999,
			payload: author.AuthorRequest{
				Name:        "Test Name",
				Description: "Test Description",
			},
			expectedStatus: fiber.StatusNotFound,
			expectedError:  true,
		},
		{
			name:     "Invalid Update - Empty Name",
			authorID: createResp.ID,
			payload: author.AuthorRequest{
				Name:        "",
				Description: "Test Description",
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/authors/%d", tt.authorID), bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if !tt.expectedError {
				body, _ := io.ReadAll(resp.Body)
				var response author.AuthorDetailResponse
				err = json.Unmarshal(body, &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.authorID, response.ID)
				assert.Equal(t, tt.payload.Name, response.Name)
				assert.Equal(t, tt.payload.Description, response.Description)
			} else {
				var errorResponse map[string]string
				body, _ := io.ReadAll(resp.Body)
				err = json.Unmarshal(body, &errorResponse)
				assert.NoError(t, err)
				assert.NotNil(t, errorResponse["error"])
			}
		})
	}
}

func TestListAuthors(t *testing.T) {
	app, db := setupTestApp()
	defer cleanupTestDB(db)

	// Create multiple authors
	authors := []author.AuthorRequest{
		{Name: "Author 1", Description: "Description 1"},
		{Name: "Author 2", Description: "Description 2"},
		{Name: "Author 3", Description: "Description 3"},
	}

	for _, a := range authors {
		payload, _ := json.Marshal(a)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/authors", bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		app.Test(req)
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		expectedStatus int
		expectedCount  int
	}{
		{
			name: "List All Authors",
			queryParams: map[string]string{
				"p":     "1",
				"limit": "10",
			},
			expectedStatus: fiber.StatusOK,
			expectedCount:  3,
		},
		{
			name: "List Authors with Pagination",
			queryParams: map[string]string{
				"p":     "1",
				"limit": "2",
			},
			expectedStatus: fiber.StatusOK,
			expectedCount:  2,
		},
		{
			name: "Search Author by Name",
			queryParams: map[string]string{
				"authorName": "Author 1",
			},
			expectedStatus: fiber.StatusOK,
			expectedCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build query string
			query := url.Values{}
			for key, value := range tt.queryParams {
				query.Add(key, value)
			}
			path := "/api/v1/authors"
			if len(query) > 0 {
				path += "?" + query.Encode()
			}

			req := httptest.NewRequest(http.MethodGet, path, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			var response author.AuthorListResponse
			err = json.Unmarshal(body, &response)
			assert.NoError(t, err)
			assert.Len(t, response.Result, tt.expectedCount)
		})
	}
}
