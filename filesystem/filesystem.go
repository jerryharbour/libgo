// Package filesystem provide types and methods for the filesystem, as an abstraction layer

package filesystem

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type FileSystem struct {
	Fser
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (f FileSystem) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ReadDir(f.Fser, dirname)
}

func ReadDir(fs Fser, dirname string) ([]os.FileInfo, error) {
	f, err := fs.Open(dirname)
	if err != nil {
		return nil, err
	}

	finfos, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	return finfos, nil
}

// ReadFile reads context of the file identified filename.
// returns err == nil if successful, or err == EOF. Because ReadFile
// reads the whole file, do not try reading very large file
func (f FileSystem) ReadFile(filename string) ([]byte, error) {
	return ReadFile(f.Fser, filename)
}

func ReadFile(fs Fser, filename string) ([]byte, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var fsize int64
	if fstat, err := f.Stat(); err == nil {
		// do not preallocate a huge buffer
		if size := fstat.Size(); size < 1e9 {
			fsize = size
		} else {
			return nil, fmt.Errorf("the size[%d] of file is too big", size)
		}
	}

	return readAll(f, fsize*bytes.MinRead)
}

// readAll reads data from r until an error or EOF
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))

	defer func() {
		e := recover()
		if e == nil {
			return
		}

		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()

	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

// WriteFile writes data to teh file identified filename.
// If the file is not existed, will create a file named filename with permission perm
func (f FileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return WriteFile(f, filename, data, perm)
}

func WriteFile(fs Fser, filename string, data []byte, perm os.FileMode) error {
	f, err := fs.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}

	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}

	if errClose := f.Close(); err == nil {
		err = errClose
	}

	return err
}
