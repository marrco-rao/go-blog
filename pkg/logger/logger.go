package logger

import (
	"go-blog/internal/config"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 统一日志，使用zap
// 后续调整记录：配置日志输出字段？

var Log *zap.Logger

// 日志写文件,stdout 同步输出
func Init() {
	cfg := config.Cfg.Log

	// ---------- 1. 日志级别 ----------
	level := zapcore.InfoLevel
	_ = level.Set(strings.ToLower(cfg.Level))

	// ---------- 2. Encoder 配置 ----------
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	// ---------- 3. lumberjack 文件输出 ----------
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename,   // 使用 log.filename
		MaxSize:    cfg.MaxSize,    // MB
		MaxBackups: cfg.MaxBackups, // 文件数量
		MaxAge:     cfg.MaxAge,     // 天
		Compress:   true,           // 是否压缩
	})

	// ---------- 4. stdout 输出 ----------
	consoleWriter := zapcore.AddSync(os.Stdout)

	// ---------- 5. 多路输出 ----------
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, consoleWriter, level),
		zapcore.NewCore(encoder, fileWriter, level),
	)

	// ---------- 6. 构建 Logger ----------
	Log = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
}

// 日志全部写入到 stdout（标准输出）
func InitScene() {
	cfg := config.Cfg.Log

	level := zapcore.InfoLevel
	_ = level.Set(strings.ToLower(cfg.Level))

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Sync() {
	_ = Log.Sync()
}
