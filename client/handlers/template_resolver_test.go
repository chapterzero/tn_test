package handlers

import (
	"os"
	"testing"
)

func TestGetTemplatePath(t *testing.T) {
	cwd, _ := os.Getwd()

	expectedPath := cwd + "/somedir/tplfile.gohtml"
	actual := GetTemplatePath("./somedir/tplfile.gohtml")

	if actual != expectedPath {
		t.Errorf("Expected %v, got %v", expectedPath, actual)
	}
}
