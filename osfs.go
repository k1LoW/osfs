package osfs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var _ fs.FS = (*OSFS)(nil)
var _ fs.SubFS = (*OSFS)(nil)

type OSFS struct {
	root string
}

func New() *OSFS {
	return &OSFS{
		root: "/",
	}
}

func (fsys *OSFS) Open(name string) (fs.File, error) {
	path := filepath.Join(fsys.root, name)
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fsys *OSFS) Sub(dir string) (fs.FS, error) {
	if strings.HasPrefix(dir, "/") {
		return nil, fmt.Errorf("invalid name: %s", dir)
	}
	sub := filepath.Join(fsys.root, dir)
	d, err := os.Stat(sub)
	if err != nil {
		return nil, err
	}
	if !d.IsDir() {
		return nil, fmt.Errorf("%s is not directory", sub)
	}
	return &OSFS{
		root: sub,
	}, nil
}
