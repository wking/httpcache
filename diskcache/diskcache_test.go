package diskcache

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gregjones/httpcache/test"
)

func TestDiskCache(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "httpcache")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	test.StreamingCache(t, New(tempDir))
}
