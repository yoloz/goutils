package logger

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestRotatingFile(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	logwriter, err := NewRotatingFileWriter(filepath.Join(homeDir, "logs", "test.log"), 1)
	if err != nil {
		panic(err)
	}
	defer logwriter.Close()
	Loger := log.New(logwriter, "[Prefix]", log.LstdFlags|log.Lshortfile)
	for i := 0; i < 100; i++ {
		Loger.Println("Counter:", i)
	}
}
