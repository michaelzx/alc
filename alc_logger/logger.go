package alc_logger

import (
  "errors"
  "fmt"
  "github.com/michaelzx/alc/alc_config"
  "github.com/michaelzx/alc/alc_fs"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "log"
  "path/filepath"
  "strconv"
  "strings"
  "time"
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
    return nil, err
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
    log.Println("log will be written to", logPath)
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
    // EncodeLevel = zapcore.LowercaseColorLevelEncoder
    EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
      s, ok := _levelToLowercaseColorString[level]
      if !ok {
        s = _unknownLevelColor.Add(level.String())
      }
      enc.AppendString(s)
    }
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
      MessageKey:    "message",
      LevelKey:      "level",
      TimeKey:       "time",
      NameKey:       "name",
      CallerKey:     "caller",
      FunctionKey:   "",
      StacktraceKey: "stacks",
      LineEnding:    zapcore.DefaultLineEnding,
      // EncodeLevel:   zapcore.CapitalColorLevelEncoder,
      EncodeLevel: EncodeLevel,
      // EncodeTime:    zapcore.ISO8601TimeEncoder,
      EncodeTime: func(time time.Time, enc zapcore.PrimitiveArrayEncoder) {
        enc.AppendString(time.Format("06-01-02T15:04:05.000"))
      },
      EncodeDuration: zapcore.StringDurationEncoder,
      EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
        enc.AppendString(cutStr(caller.FullPath(), 50))
        funcFullName := caller.Function
        start := strings.LastIndex(funcFullName, "/")
        enc.AppendString(cutStr(funcFullName[start+1:], 30))
      },
      EncodeName:       zapcore.FullNameEncoder,
      ConsoleSeparator: "  ",
    },
    OutputPaths:      OutputPaths,
    ErrorOutputPaths: ErrorOutputPaths,
    InitialFields:    nil,
  }
}
func cutStr(fullPath string, maxLen int) string {
  fullPathLen := len(fullPath)
  switch {
  case fullPathLen == maxLen:
    return fullPath
  case fullPathLen < maxLen:
    lenStr := strconv.Itoa(maxLen)
    return fmt.Sprintf("%"+lenStr+"s", fullPath)
  case fullPathLen > maxLen:
    start := fullPathLen - (maxLen - 2)
    if start < 0 {
      start = 0
    }
    return ".." + fullPath[start:]
  }
  return ""
}
