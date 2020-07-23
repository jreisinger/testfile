package testfile

import (
	"io/ioutil"
	"log"
	"os"
)

// TestFile represents a temporary file to be used for testing purposes.
type TestFile struct {
	Name string
}

// New creates a TestFile with the content.
func New(content []byte) *TestFile {
	testfile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := testfile.Write([]byte(content)); err != nil {
		log.Fatal(err)
	}
	if err := testfile.Close(); err != nil {
		log.Fatal(err)
	}

	return &TestFile{Name: testfile.Name()}
}

// Remove remove a TestFile from filesystem.
func (f *TestFile) Remove() {
	if err := os.Remove(f.Name); err != nil {
		log.Fatal(err)
	}
}
