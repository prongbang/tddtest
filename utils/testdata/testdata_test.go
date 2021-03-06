package testdata

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

// +build integration
func helperReadData(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
