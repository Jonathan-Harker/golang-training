package sha256

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func TestSha256ReturnsHashList(t *testing.T) {
	hashListExample := []HashValues{{"", ""}}

	dir, _ := os.Getwd()
	args := []string{filepath.Join(dir, "sha256.go")}
	details := sha256envs(args)

	if !IsInstanceOf(details, hashListExample) {
		t.Error("sha256envs returned wrong type")
	}

	if len(details) != 1 {
		t.Error("Expected 1 result")
	}
}
