package iocsv

import (
	"encoding/csv"
	"os"
	"os/user"

	utils "github.com/dakong/gofeed/utils"
	color "github.com/fatih/color"
)

const (
	cStoreFeed string = "There was an error storing the RSS feed"
	cIO        string = "There was an error storing or reading the file data"
)

var ioERR = color.RedString(cIO)
var storeFeedERR = color.RedString(cStoreFeed)

// File ...
type File struct {
	filePath string
}

// NewFile ...
func NewFile(name string) *File {
	usr, err := user.Current()
	utils.HandleError(err, ioERR, true)

	filePath := usr.HomeDir + "/" + name
	return &File{filePath: filePath}
}

// Read ...
func (f File) Read() (records [][]string) {
	file, readErr := os.OpenFile(f.filePath, os.O_RDWR|os.O_CREATE, 0777)
	utils.HandleError(readErr, storeFeedERR, true)
	defer file.Close()
	csvReader := csv.NewReader(file)

	records, csvReadErr := csvReader.ReadAll()
	utils.HandleError(csvReadErr, storeFeedERR, true)

	return records
}

// Store ...
func (f File) Store(record []string) bool {
	// Open csv file
	file, readErr := os.OpenFile(f.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	utils.HandleError(readErr, storeFeedERR, true)
	defer file.Close()

	fileInfo, statErr := file.Stat()
	utils.HandleError(statErr, storeFeedERR, true)

	char := make([]byte, 1)

	// Check if the last character is \n, if not then add it to the end of the file
	file.ReadAt(char, fileInfo.Size()-1)
	if char[0] != '\n' {
		file.Write([]byte{'\n'})
	}

	csvWriter := csv.NewWriter(file)

	writeErr := csvWriter.Write(record)
	utils.HandleError(writeErr, storeFeedERR, true)

	csvWriter.Flush()
	utils.HandleError(csvWriter.Error(), storeFeedERR, true)

	return true
}
