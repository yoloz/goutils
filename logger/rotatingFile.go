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
	maxBackups  int
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
		// Remove legacy _bak file
		bakFile := fmt.Sprintf("%s_bak", w.absfile)
		os.Remove(bakFile) // ignore error
		if w.maxBackups > 0 {
			// Remove the oldest backup if it exists
			oldest := fmt.Sprintf("%s.%d", w.absfile, w.maxBackups)
			os.Remove(oldest) // ignore error
			// Shift backups
			for i := w.maxBackups - 1; i >= 1; i-- {
				old := fmt.Sprintf("%s.%d", w.absfile, i)
				new := fmt.Sprintf("%s.%d", w.absfile, i+1)
				os.Rename(old, new) // ignore error if old doesn't exist
			}
			// Rename current file to backup 1
			backup1 := fmt.Sprintf("%s.1", w.absfile)
			os.Rename(w.absfile, backup1) // ignore error if current file doesn't exist
		} else {
			// No backups, delete current file
			os.Remove(w.absfile) // ignore error
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

func NewRotatingFileWriter(absfile string, maxSize int64, maxBackups ...int) (*RotatingFileWriter, error) {
	backups := 1
	if len(maxBackups) > 0 {
		backups = maxBackups[0]
		if backups < 0 {
			backups = 0
		}
	}
	w := &RotatingFileWriter{
		absfile:    absfile,
		maxSize:    maxSize,
		maxBackups: backups,
	}
	err := w.openNewFile()
	if err != nil {
		return nil, err
	}
	return w, nil
}
