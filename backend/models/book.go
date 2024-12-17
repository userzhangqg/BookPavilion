package models

import (
	"time"
	"gorm.io/gorm"
)

// BookFormat 图书格式枚举
type BookFormat string

const (
	FormatPDF  BookFormat = "pdf"
	FormatEPUB BookFormat = "epub"
	FormatTXT  BookFormat = "txt"
	FormatMOBI BookFormat = "mobi"
)

// Book 图书模型
type Book struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"size:200;not null" json:"title"`
	Author    string         `gorm:"size:100" json:"author"`
	Format    BookFormat     `gorm:"size:10" json:"format"`
	FilePath  string         `gorm:"size:500" json:"file_path"`
	FileSize  int64          `json:"file_size"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (Book) TableName() string {
	return "books"
}

// Validate 验证图书数据
func (b *Book) Validate() error {
	if b.Title == "" {
		return ErrTitleRequired
	}
	if b.Format == "" {
		return ErrFormatRequired
	}
	if b.FilePath == "" {
		return ErrFilePathRequired
	}
	return nil
}
