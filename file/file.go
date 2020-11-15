package file

import (
	"bufio"
	"io/ioutil"
	"os"
)

// File .
type File struct {
	file    *os.File
	err     error
	scanner *bufio.Scanner
}

// CheckFileIsExist .
func CheckFileIsExist(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Open .
func Open(filename string) *File {
	f, err := os.Open(filename)
	file := File{
		file: f,
		err:  err,
	}
	return &file
}

// ReadAll .
func (f *File) ReadAll() ([]byte, error) {
	if f.err != nil {
		return nil, f.err
	}
	bytes, err := ioutil.ReadAll(f.file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// ReadLine .
func (f *File) ReadLine() string {
	if f.scanner == nil {
		f.scanner = bufio.NewScanner(f.file)
	}
	if f.scanner.Scan() {
		return f.scanner.Text()
	}
	return ""

}

// ReadLines .
func (f *File) ReadLines() []string {
	var lines []string
	line := f.ReadLine()
	for {
		if line != "" {
			lines = append(lines, line)
		} else {
			break
		}
	}
	return lines
}
