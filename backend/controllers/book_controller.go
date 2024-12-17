package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zven/bookpavilion/models"
	"github.com/zven/bookpavilion/services"
)

// BookController handles HTTP requests for books
type BookController struct {
	bookService services.BookService
}

// NewBookController creates a new instance of BookController
func NewBookController(bookService services.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

// CreateBook handles book creation request
func (c *BookController) CreateBook(ctx *gin.Context) {
	// Get form data
	title := ctx.PostForm("title")
	author := ctx.PostForm("author")

	// Get the uploaded file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No file uploaded",
		})
		return
	}

	// Create book using service
	book, err := c.bookService.CreateBook(title, author, file)
	if err != nil {
		switch err {
		case models.ErrInvalidFormat:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case models.ErrTitleRequired:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

// GetBook handles single book retrieval request
func (c *BookController) GetBook(ctx *gin.Context) {
	// Parse book ID from URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Get book using service
	book, err := c.bookService.GetBook(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// ListBooks handles book list request
func (c *BookController) ListBooks(ctx *gin.Context) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	// Get books using service
	books, total, err := c.bookService.ListBooks(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

func (c *BookController) GetBookContent(ctx *gin.Context) {
    // Parse book ID from URL
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    // Get book content using service
    content, err := c.bookService.GetBookContent(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Book content not found"})
        return
    }

    ctx.JSON(http.StatusOK, content)
}
func (c *BookController) DeleteBook(ctx *gin.Context) {
	// Parse book ID from URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Delete book using service
	if err := c.bookService.DeleteBook(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
