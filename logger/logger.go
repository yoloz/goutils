package logger

import (
	"fmt"
	"io"
	"sync"
	"strings"
	"time"
)

// Level 表示日志级别
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var levelNames = map[Level]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
}

// ParseLevel 将字符串解析为 Level
func ParseLevel(s string) (Level, error) {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return LevelDebug, nil
	case "INFO":
		return LevelInfo, nil
	case "WARN", "WARNING":
		return LevelWarn, nil
	case "ERROR":
		return LevelError, nil
	default:
		return LevelInfo, fmt.Errorf("invalid level: %s", s)
	}
}

// String 返回 Level 的字符串表示
func (l Level) String() string {
	if name, ok := levelNames[l]; ok {
		return name
	}
	return fmt.Sprintf("Level(%d)", l)
}

// Logger 提供带级别过滤的日志记录器
type Logger struct {
	mu     sync.Mutex
	level  Level
	writer io.Writer
	prefix string
	flags  int
}

// NewLogger 创建一个新的 Logger
func NewLogger(writer io.Writer, level Level, prefix string, flags int) *Logger {
	return &Logger{
		writer: writer,
		level:  level,
		prefix: prefix,
		flags:  flags,
	}
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// SetOutput 设置输出写入器
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.writer = w
}

// SetPrefix 设置日志前缀
func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

// SetFlags 设置格式标志
func (l *Logger) SetFlags(flags int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.flags = flags
}

// formatHeader 格式化日志头
func (l *Logger) formatHeader(level Level, t time.Time) string {
	var buf []byte
	if l.prefix != "" {
		buf = append(buf, l.prefix...)
	}
	if l.flags&(1<<0) != 0 { // Ldate
		buf = append(buf, t.Format("2006/01/02")...)
		buf = append(buf, ' ')
	}
	if l.flags&(1<<1) != 0 { // Ltime
		buf = append(buf, t.Format("15:04:05")...)
		buf = append(buf, ' ')
	}
	if l.flags&(1<<2) != 0 { // Lmicroseconds
		buf = append(buf, t.Format("15:04:05.000000")...)
		buf = append(buf, ' ')
	}
	if levelNames[level] != "" {
		buf = append(buf, '[')
		buf = append(buf, levelNames[level]...)
		buf = append(buf, ']')
		buf = append(buf, ' ')
	}
	return string(buf)
}

// output 写入日志行
func (l *Logger) output(level Level, s string) error {
	if level < l.level {
		return nil
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.writer == nil {
		return nil
	}
	t := time.Now()
	header := l.formatHeader(level, t)
	line := header + s + "\n"
	_, err := l.writer.Write([]byte(line))
	return err
}

// Debug 输出调试日志
func (l *Logger) Debug(v ...any) {
	l.output(LevelDebug, fmt.Sprint(v...))
}

// Debugf 格式化输出调试日志
func (l *Logger) Debugf(format string, v ...any) {
	l.output(LevelDebug, fmt.Sprintf(format, v...))
}

// Debugln 输出调试日志并换行
func (l *Logger) Debugln(v ...any) {
	l.output(LevelDebug, fmt.Sprintln(v...))
}

// Info 输出信息日志
func (l *Logger) Info(v ...any) {
	l.output(LevelInfo, fmt.Sprint(v...))
}

// Infof 格式化输出信息日志
func (l *Logger) Infof(format string, v ...any) {
	l.output(LevelInfo, fmt.Sprintf(format, v...))
}

// Infoln 输出信息日志并换行
func (l *Logger) Infoln(v ...any) {
	l.output(LevelInfo, fmt.Sprintln(v...))
}

// Warn 输出警告日志
func (l *Logger) Warn(v ...any) {
	l.output(LevelWarn, fmt.Sprint(v...))
}

// Warnf 格式化输出警告日志
func (l *Logger) Warnf(format string, v ...any) {
	l.output(LevelWarn, fmt.Sprintf(format, v...))
}

// Warnln 输出警告日志并换行
func (l *Logger) Warnln(v ...any) {
	l.output(LevelWarn, fmt.Sprintln(v...))
}

// Error 输出错误日志
func (l *Logger) Error(v ...any) {
	l.output(LevelError, fmt.Sprint(v...))
}

// Errorf 格式化输出错误日志
func (l *Logger) Errorf(format string, v ...any) {
	l.output(LevelError, fmt.Sprintf(format, v...))
}

// Errorln 输出错误日志并换行
func (l *Logger) Errorln(v ...any) {
	l.output(LevelError, fmt.Sprintln(v...))
}

// Default logger instance
var (
	defaultLogger *Logger
	initOnce      sync.Once
)

// Default 返回默认 Logger 实例
func Default() *Logger {
	initOnce.Do(func() {
		defaultLogger = NewLogger(nil, LevelInfo, "", 0)
	})
	return defaultLogger
}

// SetDefaultLevel 设置默认 Logger 的级别
func SetDefaultLevel(level Level) {
	Default().SetLevel(level)
}

// SetDefaultOutput 设置默认 Logger 的输出
func SetDefaultOutput(w io.Writer) {
	Default().SetOutput(w)
}

// SetDefaultPrefix 设置默认 Logger 的前缀
func SetDefaultPrefix(prefix string) {
	Default().SetPrefix(prefix)
}

// SetDefaultFlags 设置默认 Logger 的标志
func SetDefaultFlags(flags int) {
	Default().SetFlags(flags)
}

// 包级快捷函数
func Debug(v ...any) {
	Default().Debug(v...)
}

func Debugf(format string, v ...any) {
	Default().Debugf(format, v...)
}

func Debugln(v ...any) {
	Default().Debugln(v...)
}

func Info(v ...any) {
	Default().Info(v...)
}

func Infof(format string, v ...any) {
	Default().Infof(format, v...)
}

func Infoln(v ...any) {
	Default().Infoln(v...)
}

func Warn(v ...any) {
	Default().Warn(v...)
}

func Warnf(format string, v ...any) {
	Default().Warnf(format, v...)
}

func Warnln(v ...any) {
	Default().Warnln(v...)
}

func Error(v ...any) {
	Default().Error(v...)
}

func Errorf(format string, v ...any) {
	Default().Errorf(format, v...)
}

func Errorln(v ...any) {
	Default().Errorln(v...)
}