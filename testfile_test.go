package testfile

import (
	"os"
	"testing"
)

const content = `line1
line2
line3`

func TestNew(t *testing.T) {
	// Create a new TestFile.
	file := New(content)
	defer file.Remove()

	if file.Name == "" {
		t.Errorf("test file has no name")
	}

	f, err := os.Open(file.Name)
	if err != nil {
		t.Errorf("can't open test file %s: %v", file.Name, err)
	}
	fi, err := f.Stat()
	if err != nil {
		t.Errorf("can't stat test file %s: %v", file.Name, err)
	}
	if fi.Size() != int64(len(content)) {
		t.Errorf("test file size is %d should be %d", fi.Size(), len(content))
	}
}
