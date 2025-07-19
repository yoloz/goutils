package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type RotatingFileWriter struct {
	mu          sync.Mutex
	absfile     string
	file        *os.File
	maxSize     int64
	currentSize int64
}

func (w *RotatingFileWriter) openNewFile() error {
	dir := filepath.Dir(w.absfile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	if w.file != nil {
		w.file.Close()
		w.file = nil
		bakFileName := fmt.Sprintf("%s_bak", w.absfile)
		err := os.Rename(w.absfile, bakFileName)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(w.absfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	w.file = file
	fi, err := w.file.Stat()
	if err != nil {
		return err
	}
	w.currentSize = fi.Size()
	return nil
}

// Write实现io.Writer接口
func (w *RotatingFileWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.currentSize+int64(len(p)) > w.maxSize {
		err := w.openNewFile()
		if err != nil {
			return 0, err
		}
	}
	n, err = w.file.Write(p)
	w.currentSize += int64(n)
	return n, err
}

// Close 关闭文件
func (w *RotatingFileWriter) Close() error {
	if w.file != nil {
		return w.file.Close()
	}
	return nil
}

func NewRotatingFileWriter(absfile string, maxSize int64) (*RotatingFileWriter, error) {
	w := &RotatingFileWriter{
		absfile: absfile,
		maxSize: maxSize,
	}
	err := w.openNewFile()
	if err != nil {
		return nil, err
	}
	return w, nil
}
