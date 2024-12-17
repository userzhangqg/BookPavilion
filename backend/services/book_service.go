package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/zven/bookpavilion/config"
	"github.com/zven/bookpavilion/mocks"
	"github.com/zven/bookpavilion/models"
	"gorm.io/gorm"
)

// BookService defines the interface for book operations
type BookService interface {
	CreateBook(title, author string, file *multipart.FileHeader) (*models.Book, error)
	GetBook(id uint) (*models.Book, error)
	ListBooks(page, pageSize int) ([]models.Book, int64, error)
	DeleteBook(id uint) error
	GetBookContent(id uint) (string, error) // Add this line
}

// bookService implements BookService interface
type bookService struct {
	db *gorm.DB
}

// NewBookService creates a new instance of BookService
func NewBookService(db *gorm.DB) BookService {
	return &bookService{
		db: db,
	}
}

// CreateBook implements BookService.CreateBook
func (s *bookService) CreateBook(title, author string, file *multipart.FileHeader) (*models.Book, error) {
	// Validate title
	if title == "" {
		return nil, models.ErrTitleRequired
	}

	// Get file extension and validate format
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		return nil, models.ErrInvalidFormat
	}
	format := models.BookFormat(ext[1:]) // Remove the dot from extension
	if !models.IsValidBookFormat(format) {
		return nil, models.ErrInvalidFormat
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), file.Filename)
	filepath := filepath.Join(config.GetUploadDir(), filename)

	// Save file
	var src multipart.File
	var err error

	// Try to get mock file first
	if mockFile, ok := mocks.GetMockFile(file); ok {
		src = mockFile
	} else {
		src, err = file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open uploaded file: %v", err)
		}
	}
	defer src.Close()

	dst, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		os.Remove(filepath) // Clean up on error
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	// Create book record
	book := &models.Book{
		Title:    title,
		Author:   author,
		Format:   format,
		FilePath: filename,
		FileSize: file.Size,
	}

	// Save to database
	if err := s.db.Create(book).Error; err != nil {
		os.Remove(filepath) // Clean up file if database save fails
		return nil, fmt.Errorf("failed to save book to database: %v", err)
	}

	return book, nil
}

// GetBook implements BookService.GetBook
func (s *bookService) GetBook(id uint) (*models.Book, error) {
	var book models.Book
	if err := s.db.First(&book, id).Error; err != nil {
		return nil, fmt.Errorf("book not found: %v", err)
	}
	return &book, nil
}

// ListBooks implements BookService.ListBooks
func (s *bookService) ListBooks(page, pageSize int) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64

	// Get total count
	if err := s.db.Model(&models.Book{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count books: %v", err)
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	// Get books with pagination
	if err := s.db.Offset(offset).Limit(pageSize).Find(&books).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch books: %v", err)
	}

	return books, total, nil
}

// GetBookContent retrieves the content of a book by its ID
func (s *bookService) GetBookContent(id uint) (string, error) {
	var book models.Book
	if err := s.db.First(&book, id).Error; err != nil {
		return "", fmt.Errorf("book not found: %v", err)
	}

	// Read the content from the file based on format
	filePath := filepath.Join(config.GetUploadDir(), book.FilePath)
	switch book.Format {
	case "pdf":
		// Handle PDF content loading
		return "PDF content loading not implemented", nil
	case "epub":
		// // Handle EPUB content loading
		// epub, err := ebooklib.Open(filePath)
		// if err != nil {
		// 	return "", fmt.Errorf("failed to open EPUB file: %v", err)
		// }
		// // Extract text content from the EPUB
		// var content string
		// for _, item := range epub.Items {
		// 	if item.MediaType == "application/xhtml+xml" {
		// 		text, err := item.Content()
		// 		if err != nil {
		// 			return "", fmt.Errorf("failed to read EPUB content: %v", err)
		// 		}
		// 		content += string(text)
		// 	}
		// }
		// return content, nil
		return "EPUB content loading not implemented", nil
	case "txt":
		content, err := os.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to read book content: %v", err)
		}
		return string(content), nil
	default:
		return "", fmt.Errorf("unsupported book format: %s", book.Format)
	}
}

// DeleteBook implements BookService.DeleteBook
func (s *bookService) DeleteBook(id uint) error {
	var book models.Book

	// Get book first to get the file path
	if err := s.db.First(&book, id).Error; err != nil {
		return fmt.Errorf("book not found: %v", err)
	}

	// Delete from database
	if err := s.db.Delete(&book).Error; err != nil {
		return fmt.Errorf("failed to delete book from database: %v", err)
	}

	// Delete file
	filepath := filepath.Join(config.GetUploadDir(), book.FilePath)
	if err := os.Remove(filepath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete book file: %v", err)
	}

	return nil
}
