-- Create test database if not exists
CREATE DATABASE IF NOT EXISTS bookpavilion_test;

-- Use test database
USE bookpavilion_test;

-- Create books table
CREATE TABLE IF NOT EXISTS books (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    author VARCHAR(100),
    format VARCHAR(10) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    file_size BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create indexes
CREATE INDEX idx_books_title ON books(title);
CREATE INDEX idx_books_author ON books(author);
CREATE INDEX idx_books_created_at ON books(created_at);
CREATE INDEX idx_books_deleted_at ON books(deleted_at);

-- Grant privileges to test database
GRANT ALL PRIVILEGES ON bookpavilion_test.* TO 'root'@'localhost';
FLUSH PRIVILEGES;
