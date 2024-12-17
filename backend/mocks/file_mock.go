package mocks

import (
	"bytes"
	"mime/multipart"
)

// MockMultipartFile implements multipart.File interface
type MockMultipartFile struct {
	*bytes.Reader
	closed bool
}

func (f *MockMultipartFile) Close() error {
	f.closed = true
	return nil
}

// MockFileHeader wraps multipart.FileHeader to provide a working Open method
type MockFileHeader struct {
	*multipart.FileHeader
	content []byte
}

// Package-level map to store mock files
var mockFiles = make(map[*multipart.FileHeader][]byte)

// NewMockFileHeader creates a new multipart.FileHeader with mock content
func NewMockFileHeader(filename string, size int64, content []byte) *multipart.FileHeader {
	fh := &multipart.FileHeader{
		Filename: filename,
		Size:     size,
	}
	mockFiles[fh] = content
	return fh
}

// GetMockFile retrieves mock content for a file header
func GetMockFile(fh *multipart.FileHeader) (multipart.File, bool) {
	content, ok := mockFiles[fh]
	if !ok {
		return nil, false
	}
	return &MockMultipartFile{
		Reader: bytes.NewReader(content),
	}, true
}

// ClearMockFiles cleans up all mock files
func ClearMockFiles() {
	mockFiles = make(map[*multipart.FileHeader][]byte)
}

// CreateMockFile creates a mock file with content for testing
func CreateMockFile(content []byte) multipart.File {
	return &MockMultipartFile{
		Reader: bytes.NewReader(content),
	}
}
