package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// EnsureDirectoryExists, verilen yoldaki dizinin var olduğundan emin olur
func EnsureDirectoryExists(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	return nil
}

// CreateFileWithTimestamp, zaman damgalı bir dosya oluşturur
func CreateFileWithTimestamp(basePath string) (*os.File, error) {
	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("logs_%s.txt", timestamp)
	fullPath := filepath.Join(basePath, fileName)

	if err := EnsureDirectoryExists(fullPath); err != nil {
		return nil, err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file %s: %w", fullPath, err)
	}

	return file, nil
}

// WriteToFile, verilen içeriği dosyaya yazar
func WriteToFile(file *os.File, content string) error {
	writer := bufio.NewWriter(file)
	_, err := writer.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}
	return nil
}

// SanitizeFileName, dosya adını güvenli hale getirir
func SanitizeFileName(fileName string) string {
	// Geçersiz karakterleri kaldır veya değiştir
	sanitized := strings.Map(func(r rune) rune {
		if r == '/' || r == '\\' || r == ':' || r == '*' || r == '?' || r == '"' || r == '<' || r == '>' || r == '|' {
			return '_'
		}
		return r
	}, fileName)

	// Başındaki ve sonundaki boşlukları kaldır
	return strings.TrimSpace(sanitized)
}

// TruncateString, bir stringi belirli bir uzunluğa kısaltır
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength-3] + "..."
}

// IsFileEmpty, bir dosyanın boş olup olmadığını kontrol eder
func IsFileEmpty(fileName string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return false, fmt.Errorf("failed to open file %s: %w", fileName, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return false, fmt.Errorf("failed to get file stats for %s: %w", fileName, err)
	}

	return stat.Size() == 0, nil
}