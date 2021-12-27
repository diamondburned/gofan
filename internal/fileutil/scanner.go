package fileutil

import (
	"fmt"
	"os"
	"strconv"
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

// seek opens the file if it isn't already or seeks to start.
func (s *Scanner) seek() error {
	if s.file == nil {
		f, err := os.Open(s.path)
		if err != nil {
			return err
		}
		s.file = f
	} else {
		// seek to start
		if _, err := s.file.Seek(0, 0); err != nil {
			return fmt.Errorf("%s: seek error: %w", s.path, err)
		}
	}

	return nil
}

// Scan scans the file from its head into the given values.
func (s *Scanner) Scan(v ...interface{}) error {
	if err := s.seek(); err != nil {
		return err
	}
	if _, err := fmt.Fscan(s.file, v...); err != nil {
		return fmt.Errorf("%s: %w", s.path, err)
	}
	return nil
}

const maxIntLen = 19 // len(strconv.Itoa(math.MaxInt)) on 64-bit

// ScanInt scans an integer.
func (s *Scanner) ScanInt() (int, error) {
	if err := s.seek(); err != nil {
		return 0, err
	}

	var buf [maxIntLen]byte
	n, err := s.file.Read(buf[:])
	if err != nil {
		return 0, fmt.Errorf("cannot read int: %w", err)
	}

	return strconv.Atoi(string(buf[:n]))
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
