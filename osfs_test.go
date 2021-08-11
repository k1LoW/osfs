package osfs

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
)

func TestSub(t *testing.T) {
	fsys := New()
	abs, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	subs := []fs.FS{}
	sub, err := fs.Sub(fsys, strings.TrimPrefix(abs, "/"))
	if err != nil {
		t.Fatal(err)
	}
	subs = append(subs, sub)
	sub2, err := fsys.Sub(strings.TrimPrefix(abs, "/"))
	if err != nil {
		t.Fatal(err)
	}
	subs = append(subs, sub2)

	for _, sub := range subs {
		f, err := sub.Open("path/to/foo.txt")
		if err != nil {
			t.Error(err)
		}
		fi, err := f.Stat()
		if err != nil {
			t.Error(err)
		}
		if fi.IsDir() {
			t.Error("path/to/foo.txt is file")
		}
	}
}
