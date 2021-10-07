package service

import (
	"bufio"
	"os"
)

type FileService struct {
	filePath  string
	Delimiter string
	Lines     []string
}

func NewFileService(filePath string) *FileService {
	fileService := &FileService{
		filePath:  filePath,
		Delimiter: " ",
	}

	return fileService
}

func (f *FileService) Open() (*FileService, error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	defer file.Close()

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	f.Lines = lines
	return f, nil

}
