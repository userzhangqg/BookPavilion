package models

import "errors"

// Define model-related errors
var (
	// Book errors
	ErrTitleRequired    = errors.New("book title is required")
	ErrFormatRequired   = errors.New("book format is required")
	ErrFilePathRequired = errors.New("book file path is required")
	ErrInvalidFormat    = errors.New("unsupported book format")
	ErrFileNotFound     = errors.New("book file not found")
	ErrFileTooLarge     = errors.New("book file exceeds size limit")
)

// IsValidBookFormat checks if the book format is valid
func IsValidBookFormat(format BookFormat) bool {
	switch format {
	case FormatPDF, FormatEPUB, FormatTXT, FormatMOBI:
		return true
	default:
		return false
	}
}
