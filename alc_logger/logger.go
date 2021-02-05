package alc_logger

import (
	"alchemy/alc/alc_config"
	"alchemy/alc/alc_fs"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
)

var (
	zapLogger *zap.Logger
	// logDir    = "./app.log"
)

func New(loggerConfig alc_config.LoggerConfig, skip int) (*zap.Logger, error) {
	cfg := getZapConfig(loggerConfig.Mode)
	// 因为我们做了一层包装，所以需要跳过一层caller
	// 否则，日志的caller位置，始终显示的是当前logger包中的位置
	// TODO 到底要条几层？
	callerOption := zap.AddCallerSkip(skip)
	logger, err := cfg.Build(callerOption)
	if err != nil {
		return nil,err
	}
	// TODO 初始化的时候，需要做这个吗？
	defer func() {
		logger.Sync()
	}()
	return logger, nil
}
func Init(loggerConfig alc_config.LoggerConfig) {
	zapCfg := getZapConfig(loggerConfig.Mode)
	// 因为我们做了一层包装，所以需要跳过一层caller
	// 否则，日志的caller位置，始终显示的是当前logger包中的位置
	// TODO 到底要条几层？
	callerOption := zap.AddCallerSkip(1)
	var err error
	zapLogger, err = zapCfg.Build(callerOption)
	if err != nil {
		panic(err)
	}
	// TODO 初始化的时候，需要做这个吗？
	defer func() {
		zapLogger.Sync()
	}()
}
func DoSync() {
	zapLogger.Sync()
}

func getZapConfig(mode string) zap.Config {
	var loggingLevel zapcore.Level
	var OutputPaths []string
	var ErrorOutputPaths []string
	var Encoding string
	var EncodeLevel zapcore.LevelEncoder
	var Development bool
	switch mode {
	case "prod": // 生产模式 TODO 自动切割文件
		logPath := filepath.Join(alc_fs.AppPath, "logs")
		alc_fs.CreateIfNotExist(logPath)
		Development = false
		loggingLevel = zap.InfoLevel
		Encoding = "json"
		EncodeLevel = zapcore.CapitalLevelEncoder
		OutputPaths = []string{filepath.Join(logPath, "output.log")}
		ErrorOutputPaths = []string{filepath.Join(logPath, "errors.log")}
	case "dev": // 开发模式
		Development = true
		loggingLevel = zap.DebugLevel
		Encoding = "console"
		EncodeLevel = zapcore.CapitalColorLevelEncoder
		OutputPaths = []string{"stdout"}
		ErrorOutputPaths = []string{"stderr"}
	default:
		panic(errors.New("日志模式不支持：" + mode))
	}
	return zap.Config{
		Level:             zap.NewAtomicLevelAt(loggingLevel),
		Development:       Development,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          Encoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:       "message",
			LevelKey:         "level",
			TimeKey:          "time",
			NameKey:          "name",
			CallerKey:        "caller",
			FunctionKey:      "func",
			StacktraceKey:    "stacks",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      EncodeLevel, // zapcore.CapitalColorLevelEncoder,
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     customCallerEncoder,
			EncodeName:       zapcore.FullNameEncoder,
			ConsoleSeparator: "",
		},
		OutputPaths:      OutputPaths,
		ErrorOutputPaths: ErrorOutputPaths,
		InitialFields:    nil,
	}
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
}
