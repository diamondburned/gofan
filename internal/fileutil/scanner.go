package fileutil

import (
	"fmt"
	"os"
)

// Scanner is a file scanner using fmt.Scan.
type Scanner struct {
	file *os.File
	path string
}

// NewScanner creates a new file scanner.
func NewScanner(path string) *Scanner {
	return &Scanner{path: path}
}

// Scan scans the file from its head into the given values.
func (s *Scanner) Scan(v ...interface{}) error {
	if s.file == nil {
		f, err := os.Open(s.path)
		if err != nil {
			return err
		}
		s.file = f
	}

	// seek to start
	if _, err := s.file.Seek(0, 0); err != nil {
		return fmt.Errorf("%s: seek error: %w", s.path, err)
	}

	if _, err := fmt.Fscan(s.file, v...); err != nil {
		return fmt.Errorf("%s: %w", s.path, err)
	}

	return nil
}

// Close closes the scanner.
func (s *Scanner) Close() error {
	if s.file == nil {
		return os.ErrClosed
	}
	err := s.file.Close()
	s.file = nil
	return err
}
