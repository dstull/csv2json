package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"testing"

	"github.com/dstull/csv2json/model"
)

func TestCanConverFile(t *testing.T) {
	args := []string{"-i=testdata/test.csv", "-o=testdata/tmp/test.json"}

	_, err := model.FromArgs(args)

	if err != nil {
		t.Error("Failed to process arguments")
	}
}

func TestMainProgram(t *testing.T) {
	goldenFile := "testdata/test.json"
	outputFile := "testdata/tmp/test_main.json"

	os.Args = []string{"bogus",
		"-i=testdata/test.csv",
		"-o=testdata/tmp/test_main.json"}
	main()

	goldenFileMD5, err := hash_file_md5(goldenFile)
	if err != nil {
		t.Errorf("Failed to get md5 sum of %v", goldenFile)
	}

	outputFileMD5, err := hash_file_md5(outputFile)
	if err != nil {
		t.Errorf("Failed to get md5 sum of %v", outputFile)
	}

	if goldenFileMD5 != outputFileMD5 {
		t.Errorf("Checksum of %v does not match %v", outputFile, goldenFile)
	}
}

func hash_file_md5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil

}
