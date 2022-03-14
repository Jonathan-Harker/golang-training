package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"hash"
	"io"
	"os"
)

type HashValues struct {
	FileName string
	HexValue string
}

func sha256envs(args []string) []HashValues {
	var hashList []HashValues
	hashValue := sha256.New()
	var failures int

	for _, arg := range args {
		hashList, failures = process(hashValue, arg, failures, hashList)
	}

	if failures > 0 {
		os.Exit(1)
	}

	return hashList
}

func process(hash hash.Hash, arg string, failures int, hashList []HashValues) ([]HashValues, int) {
	hash.Reset()

	info, err, failures := getFileDetails(arg, failures)

	if info.IsDir() {
		println("Path is a directory: " + info.Name())
		return nil, failures
	}

	file, failures := openFile(err, arg, failures)

	defer file.Close()

	failures = copyFile(hash, err, file, failures)
	hashValue := buildHashMap(hash, arg)
	hashList = append(hashList, hashValue)
	return hashList, failures
}

func openFile(err error, arg string, failures int) (*os.File, int) {
	file, err := os.Open(arg)
	if err != nil {
		failures++
		println(errors.Wrap(err, "Could not open the file"))
	}
	return file, failures
}

func getFileDetails(arg string, failures int) (os.FileInfo, error, int) {
	info, err := os.Stat(arg)
	if err != nil {
		failures++
		println(errors.Wrap(err, "Path is not a valid file"))
	}
	return info, err, failures
}

func copyFile(hash hash.Hash, err error, file *os.File, failures int) int {
	_, err = io.Copy(hash, file)
	if err != nil {
		failures++
		println(errors.Wrap(err, "Could not copy the file contents"))
	}
	return failures
}

func buildHashMap(hash hash.Hash, arg string) HashValues {
	hashValue := HashValues{
		FileName: arg,
		HexValue: hex.Dump(hash.Sum(nil)),
	}
	return hashValue
}
