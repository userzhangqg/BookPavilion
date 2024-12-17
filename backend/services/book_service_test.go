package services

import (
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zven/bookpavilion/mocks"
	"github.com/zven/bookpavilion/models"
	"github.com/zven/bookpavilion/tests"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	os.Exit(code)
}

func setupTest(t *testing.T) (*bookService, sqlmock.Sqlmock, func()) {
	// Create mock database
	db, mock, err := mocks.NewMockDB()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	// Create service instance
	service := NewBookService(db).(*bookService)

	// Return cleanup function
	cleanup := func() {
		tests.CleanupTestFiles(t)
		mocks.ClearMockFiles()
	}

	return service, mock, cleanup
}

func TestCreateBook(t *testing.T) {
	service, mock, cleanup := setupTest(t)
	defer cleanup()

	testCases := []struct {
		name        string
		title       string
		author      string
		filename    string
		content     []byte
		mockSetup   func(sqlmock.Sqlmock)
		expectError bool
		errorType   error
	}{
		{
			name:     "Valid PDF Book",
			title:    "Test Book",
			author:   "Test Author",
			filename: "test.pdf",
			content:  []byte("test content"),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `books`").
					WithArgs(
						"Test Book",      // title
						"Test Author",    // author
						"pdf",            // format
						sqlmock.AnyArg(), // file_path
						int64(12),        // file_size
						sqlmock.AnyArg(), // created_at
						sqlmock.AnyArg(), // updated_at
						nil,              // deleted_at
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectError: false,
		},
		{
			name:     "Empty Title",
			title:    "",
			author:   "Test Author",
			filename: "test.pdf",
			content:  []byte("test content"),
			mockSetup: func(mock sqlmock.Sqlmock) {
				// No database expectations needed as validation fails
			},
			expectError: true,
			errorType:   models.ErrTitleRequired,
		},
		{
			name:     "Invalid Format",
			title:    "Test Book",
			author:   "Test Author",
			filename: "test.doc",
			content:  []byte("test content"),
			mockSetup: func(mock sqlmock.Sqlmock) {
				// No database expectations needed as validation fails
			},
			expectError: true,
			errorType:   models.ErrInvalidFormat,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up mock expectations
			tc.mockSetup(mock)

			// Create test file header with mock content
			file := mocks.NewMockFileHeader(tc.filename, int64(len(tc.content)), tc.content)

			// Test CreateBook
			book, err := service.CreateBook(tc.title, tc.author, file)

			if tc.expectError {
				if err == nil {
					t.Error("expected error but got none")
				} else if err != tc.errorType {
					t.Errorf("expected error %v but got %v", tc.errorType, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Verify book data
			if book.Title != tc.title {
				t.Errorf("expected title %s but got %s", tc.title, book.Title)
			}
			if book.Author != tc.author {
				t.Errorf("expected author %s but got %s", tc.author, book.Author)
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	service, mock, cleanup := setupTest(t)
	defer cleanup()

	t.Run("Get Existing Book", func(t *testing.T) {
		mock.ExpectQuery("SELECT.*FROM.*books.*WHERE.*id.*=.*").
			WithArgs(uint(1)).
			WillReturnRows(sqlmock.NewRows(mocks.BookColumns()).
				AddRow(1, "Test Book", "Test Author", "pdf", "test.pdf", 1024,
					time.Now(), time.Now(), nil))

		book, err := service.GetBook(1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if book.ID != 1 {
			t.Errorf("expected book ID 1 but got %d", book.ID)
		}
	})

	t.Run("Get Non-existent Book", func(t *testing.T) {
		mock.ExpectQuery("SELECT.*FROM.*books.*WHERE.*id.*=.*").
			WithArgs(uint(9999)).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := service.GetBook(9999)
		if err == nil {
			t.Error("expected error for non-existent book but got none")
		}
	})
}

func TestListBooks(t *testing.T) {
	service, mock, cleanup := setupTest(t)
	defer cleanup()

	t.Run("List Books with Pagination", func(t *testing.T) {
		// Set up expectations for count query
		mock.ExpectQuery("SELECT count.*FROM.*books.*").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(15))

		// Set up expectations for select query
		rows := sqlmock.NewRows(mocks.BookColumns())
		for i := 1; i <= 10; i++ {
			rows.AddRow(i, "Test Book", "Test Author", "pdf", "test.pdf", 1024,
				time.Now(), time.Now(), nil)
		}
		mock.ExpectQuery("SELECT.*FROM.*books.*").
			WillReturnRows(rows)

		books, total, err := service.ListBooks(1, 10)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if len(books) != 10 {
			t.Errorf("expected 10 books but got %d", len(books))
		}

		if total != 15 {
			t.Errorf("expected total 15 but got %d", total)
		}
	})
}

func TestDeleteBook(t *testing.T) {
	service, mock, cleanup := setupTest(t)
	defer cleanup()

	t.Run("Delete Existing Book", func(t *testing.T) {
		// Set up expectations for getting the book
		mock.ExpectQuery("SELECT.*FROM.*books.*WHERE.*id.*=.*").
			WithArgs(uint(1)).
			WillReturnRows(sqlmock.NewRows(mocks.BookColumns()).
				AddRow(1, "Test Book", "Test Author", "pdf", "test.pdf", 1024,
					time.Now(), time.Now(), nil))

		// Set up expectations for deleting the book
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE.*books.*SET.*deleted_at.*WHERE.*").
			WithArgs(sqlmock.AnyArg(), uint(1)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Create a test file
		content := []byte("test content")
		tests.CreateTestFile(t, "test.pdf", content)

		// Delete the book
		err := service.DeleteBook(1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Delete Non-existent Book", func(t *testing.T) {
		mock.ExpectQuery("SELECT.*FROM.*books.*WHERE.*id.*=.*").
			WithArgs(uint(9999)).
			WillReturnError(gorm.ErrRecordNotFound)

		err := service.DeleteBook(9999)
		if err == nil {
			t.Error("expected error for non-existent book but got none")
		}
	})
}
