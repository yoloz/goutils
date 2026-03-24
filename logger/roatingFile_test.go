package logger

import (
	"fmt"
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

func TestRotatingFileWithBackups(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "test.log")
	// Create writer with maxSize=10 bytes and maxBackups=3
	logwriter, err := NewRotatingFileWriter(logPath, 10, 3)
	if err != nil {
		t.Fatal(err)
	}
	defer logwriter.Close()
	// Write enough data to trigger rotation multiple times
	data := []byte("1234567890\n") // 11 bytes including newline
	for i := 0; i < 10; i++ {
		_, err := logwriter.Write(data)
		if err != nil {
			t.Fatal(err)
		}
	}
	// Check that backup files exist
	for i := 1; i <= 3; i++ {
		backup := fmt.Sprintf("%s.%d", logPath, i)
		if _, err := os.Stat(backup); os.IsNotExist(err) {
			t.Errorf("backup file %s does not exist", backup)
		}
	}
	// Ensure no extra backup files beyond maxBackups
	extraBackup := fmt.Sprintf("%s.%d", logPath, 4)
	if _, err := os.Stat(extraBackup); err == nil {
		t.Errorf("extra backup file %s exists beyond maxBackups", extraBackup)
	}
}
