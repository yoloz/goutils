# goutils

utils for personal use,etc.

## logger包 - 带等级过滤的日志记录器

`logger` 包提供了带级别过滤的日志记录功能，支持文件轮转和灵活的日志级别控制。

### 特性

- **等级过滤**: 支持 `DEBUG`, `INFO`, `WARN`, `ERROR` 四个级别
- **文件轮转**: 自动管理日志文件大小和备份数量
- **线程安全**: 所有操作都是线程安全的
- **灵活配置**: 可设置前缀、格式标志和输出目标
- **默认实例**: 提供包级默认 Logger 和快捷函数

### 安装

```bash
go get github.com/yoloz/goutils/logger
```

### 使用示例

#### 基本使用

```go
package main

import (
    "github.com/yoloz/goutils/logger"
    "log"
)

func main() {
    // 创建旋转文件写入器（最大1MB，保留5个备份）
    writer, err := logger.NewRotatingFileWriter("app.log", 1024*1024, 5)
    if err != nil {
        log.Fatal(err)
    }
    defer writer.Close()

    // 创建带级别过滤的 Logger
    logger := logger.NewLogger(writer, logger.LevelInfo, "[APP]", log.LstdFlags)

    logger.Info("Application started")
    logger.Debug("This will not be logged") // 级别为 Info，Debug 被过滤

    // 修改级别
    logger.SetLevel(logger.LevelDebug)
    logger.Debug("Now debug messages are visible")

    // 使用不同的日志级别
    logger.Info("Information message")
    logger.Warn("Warning message")
    logger.Error("Error message")

    // 格式化输出
    logger.Infof("User %s logged in", "alice")
    logger.Errorf("Failed to process request: %v", err)
}
```

#### 包级默认 Logger

```go
package main

import "github.com/yoloz/goutils/logger"

func main() {
    // 设置默认 Logger 的输出和级别
    writer, _ := logger.NewRotatingFileWriter("app.log", 1024*1024, 5)
    defer writer.Close()

    logger.SetDefaultOutput(writer)
    logger.SetDefaultLevel(logger.LevelWarn)

    // 使用包级快捷函数
    logger.Warn("This warning will be logged")
    logger.Info("This info will be filtered") // 被过滤

    logger.Debug("Debug message")
    logger.Error("Error message")

    // 格式化输出
    logger.Infof("User: %s", "bob")
    logger.Errorf("Error: %v", err)
}
```

#### 级别解析和字符串表示

```go
package main

import (
    "fmt"
    "github.com/yoloz/goutils/logger"
)

func main() {
    // 从字符串解析级别
    level, err := logger.ParseLevel("debug")
    if err != nil {
        fmt.Printf("Error parsing level: %v\n", err)
    } else {
        fmt.Printf("Parsed level: %v\n", level.String())
    }

    // 所有可用级别
    fmt.Printf("Debug: %v\n", logger.LevelDebug.String())
    fmt.Printf("Info: %v\n", logger.LevelInfo.String())
    fmt.Printf("Warn: %v\n", logger.LevelWarn.String())
    fmt.Printf("Error: %v\n", logger.LevelError.String())
}
```

### API 参考

#### 类型

```go
type Level int
```

日志级别常量:
- `LevelDebug` - 调试信息
- `LevelInfo`  - 一般信息
- `LevelWarn`  - 警告信息
- `LevelError` - 错误信息

#### 函数

```go
// 创建新的 RotatingFileWriter
func NewRotatingFileWriter(absfile string, maxSize int64, maxBackups ...int) (*RotatingFileWriter, error)

// 创建新的 Logger
func NewLogger(writer io.Writer, level Level, prefix string, flags int) *Logger

// 从字符串解析 Level
func ParseLevel(s string) (Level, error)

// 包级默认 Logger 设置
func SetDefaultOutput(w io.Writer)
func SetDefaultLevel(level Level)
func SetDefaultPrefix(prefix string)
func SetDefaultFlags(flags int)

// 包级快捷函数
func Debug(v ...any)
func Debugf(format string, v ...any)
func Debugln(v ...any)
func Info(v ...any)
func Infof(format string, v ...any)
func Infoln(v ...any)
func Warn(v ...any)
func Warnf(format string, v ...any)
func Warnln(v ...any)
func Error(v ...any)
func Errorf(format string, v ...any)
func Errorln(v ...any)
```

#### Logger 方法

```go
// 设置日志级别
func (l *Logger) SetLevel(level Level)

// 设置输出写入器
func (l *Logger) SetOutput(w io.Writer)

// 设置日志前缀
func (l *Logger) SetPrefix(prefix string)

// 设置格式标志
func (l *Logger) SetFlags(flags int)

// 日志输出方法（每个级别都有对应的 f 和 ln 版本）
func (l *Logger) Debug(v ...any)
func (l *Logger) Info(v ...any)
func (l *Logger) Warn(v ...any)
func (l *Logger) Error(v ...any)
```

### 级别过滤规则

日志级别按照以下顺序排列（从低到高）：
1. `DEBUG` - 最详细，用于调试
2. `INFO`  - 一般信息
3. `WARN`  - 警告信息
4. `ERROR` - 错误信息

设置某个级别后，只有该级别及更高级别的日志会被输出。例如：
- 设置为 `INFO` 级别：`INFO`, `WARN`, `ERROR` 会被输出，`DEBUG` 被过滤
- 设置为 `WARN` 级别：`WARN`, `ERROR` 会被输出，`DEBUG`, `INFO` 被过滤

### 文件轮转配置

`NewRotatingFileWriter` 函数参数：
- `absfile`: 日志文件的绝对路径
- `maxSize`: 单个日志文件的最大大小（字节）
- `maxBackups`: 可选参数，保留的备份文件数量（默认为1）

当当前日志文件超过 `maxSize` 时，会自动创建新文件并按照配置保留指定数量的备份文件。
