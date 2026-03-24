package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLoggerLevelFilter(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, LevelWarn, "", 0)

	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	output := buf.String()
	if strings.Contains(output, "debug") {
		t.Error("Debug message should not be output when level is WARN")
	}
	if strings.Contains(output, "info") {
		t.Error("Info message should not be output when level is WARN")
	}
	if !strings.Contains(output, "warn") {
		t.Error("Warn message should be output when level is WARN")
	}
	if !strings.Contains(output, "error") {
		t.Error("Error message should be output when level is WARN")
	}
}

func TestLoggerLevelChange(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, LevelInfo, "", 0)

	logger.Debug("debug message")
	output := buf.String()
	if strings.Contains(output, "debug") {
		t.Error("Debug message should not be output when level is INFO")
	}

	buf.Reset()
	logger.SetLevel(LevelDebug)
	logger.Debug("debug message")
	if !strings.Contains(buf.String(), "debug") {
		t.Error("Debug message should be output after level changed to DEBUG")
	}
}

func TestLoggerPrefix(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, LevelInfo, "[PREFIX]", 0)
	logger.Info("test")
	output := buf.String()
	if !strings.Contains(output, "[PREFIX]") {
		t.Error("Output should contain prefix")
	}
}

func TestLoggerFlags(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, LevelInfo, "", 1) // Ldate flag
	logger.Info("test")
	output := buf.String()
	// Check that date appears (format 2006/01/02)
	if !strings.Contains(output, "/") {
		t.Error("Output should contain date when flag Ldate is set")
	}
}

func TestDefaultLogger(t *testing.T) {
	var buf bytes.Buffer
	SetDefaultOutput(&buf)
	SetDefaultLevel(LevelInfo)

	Debug("debug")
	if strings.Contains(buf.String(), "debug") {
		t.Error("Default logger should filter debug messages at INFO level")
	}

	buf.Reset()
	Info("info")
	if !strings.Contains(buf.String(), "info") {
		t.Error("Default logger should output info messages at INFO level")
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected Level
		hasErr   bool
	}{
		{"debug", LevelDebug, false},
		{"DEBUG", LevelDebug, false},
		{"info", LevelInfo, false},
		{"INFO", LevelInfo, false},
		{"warn", LevelWarn, false},
		{"WARN", LevelWarn, false},
		{"warning", LevelWarn, false},
		{"error", LevelError, false},
		{"ERROR", LevelError, false},
		{"unknown", LevelInfo, true},
	}
	for _, tt := range tests {
		level, err := ParseLevel(tt.input)
		if tt.hasErr {
			if err == nil {
				t.Errorf("ParseLevel(%q) expected error, got nil", tt.input)
			}
			continue
		}
		if err != nil {
			t.Errorf("ParseLevel(%q) unexpected error: %v", tt.input, err)
			continue
		}
		if level != tt.expected {
			t.Errorf("ParseLevel(%q) = %v, want %v", tt.input, level, tt.expected)
		}
	}
}

func TestLevelString(t *testing.T) {
	if LevelDebug.String() != "DEBUG" {
		t.Errorf("LevelDebug.String() = %q, want DEBUG", LevelDebug.String())
	}
	if LevelInfo.String() != "INFO" {
		t.Errorf("LevelInfo.String() = %q, want INFO", LevelInfo.String())
	}
	if LevelWarn.String() != "WARN" {
		t.Errorf("LevelWarn.String() = %q, want WARN", LevelWarn.String())
	}
	if LevelError.String() != "ERROR" {
		t.Errorf("LevelError.String() = %q, want ERROR", LevelError.String())
	}
	// Test unknown level
	var l Level = 99
	if l.String() == "" {
		t.Error("String() for unknown level should not return empty string")
	}
}