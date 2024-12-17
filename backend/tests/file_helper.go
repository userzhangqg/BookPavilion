package tests

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
)

// CreateTestFileHeader creates a multipart.FileHeader for testing
func CreateTestFileHeader(t testing.TB, filename string, content []byte) *multipart.FileHeader {
	// Create a temporary file
	tmpDir := filepath.Join(os.TempDir(), "bookpavilion_test")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	tmpFile := filepath.Join(tmpDir, filename)
	if err := os.WriteFile(tmpFile, content, 0644); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	// Open the file
	file, err := os.Open(tmpFile)
	if err != nil {
		t.Fatalf("Failed to open temp file: %v", err)
	}
	defer file.Close()

	// Create a buffer to store the file content
	var buff bytes.Buffer
	if _, err := io.Copy(&buff, file); err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}

	// Create the FileHeader
	return &multipart.FileHeader{
		Filename: filename,
		Size:     int64(buff.Len()),
	}
}

// CreateTestFile creates a test file in the specified directory
func CreateTestFile(t testing.TB, filename string, content []byte) string {
	// Create test uploads directory if it doesn't exist
	uploadDir := GetUploadDir()
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		t.Fatalf("Failed to create test upload directory: %v", err)
	}

	// Create the test file
	filepath := filepath.Join(uploadDir, filename)
	if err := os.WriteFile(filepath, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	return filepath
}

// GetUploadDir returns the test upload directory
func GetUploadDir() string {
	return filepath.Join(os.TempDir(), "bookpavilion_test_uploads")
}

// CleanupTestFiles removes all test files
func CleanupTestFiles(t testing.TB) {
	if err := os.RemoveAll(GetUploadDir()); err != nil {
		t.Errorf("Failed to cleanup test files: %v", err)
	}
}
