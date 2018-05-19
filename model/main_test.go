package model

import (
	"testing"
)

func TestCanProcessArgs(t *testing.T) {
	args := []string{"-i=test.csv", "-o=test.json"}

	_, err := FromArgs(args)

	if err != nil {
		t.Error("Failed to process arguments")
	}
}

func TestCanNotProcessInputFile(t *testing.T) {
	args := []string{"-i=bad_file.csv", "-o=test.json"}

	_, err := FromArgs(args)

	if err == nil {
		t.Error("File does not exist and should error")
	}
}

func TestCanNotProcessOutputFile(t *testing.T) {
	args := []string{"-i=test.csv", "-o=/bad_root.json"}

	_, err := FromArgs(args)

	if err == nil {
		t.Error("Permissions should disable creation and should error")
	}
}
