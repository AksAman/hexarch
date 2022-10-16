package utils

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type loggersFactory map[string]*zap.SugaredLogger

func (lFactory *loggersFactory) GetLogger(name string) *zap.SugaredLogger {
	if logger, ok := (*lFactory)[name]; ok {
		return logger
	}
	return nil
}

func (lFactory *loggersFactory) RegisterLogger(name string, logger *zap.SugaredLogger) {
	(*lFactory)[name] = logger
}

// deregister
func (lFactory *loggersFactory) DeregisterLogger(name string) {
	delete(*lFactory, name)
}

var (
	factory loggersFactory = make(map[string]*zap.SugaredLogger)
)

// InitializeLogger initializes the zap.SugaredLogger with the given logFilename
// If logFilename is empty, it will log to stdout
func InitializeLogger(logFilename string) *zap.SugaredLogger {
	if !strings.HasSuffix(logFilename, ".log") {
		logFilename = logFilename + ".log"
	}
	logger := factory.GetLogger(logFilename)
	if logger != nil {
		return logger
	}

	var core zapcore.Core
	if logFilename == "" {
		core = zapcore.NewCore(getConsoleEncoder(), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(getConsoleEncoder(), zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(getJSONEncoder(), getLogWriter(logFilename), zapcore.DebugLevel),
		)
	}
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.FatalLevel)).Sugar()
	factory.RegisterLogger(logFilename, logger)
	defer logger.Sync()
	return logger
}

func getEncoderConfig() zapcore.EncoderConfig {
	baseConfig := zap.NewProductionEncoderConfig()
	baseConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	baseConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	baseConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return baseConfig
}

func getJSONEncoder() zapcore.Encoder {
	config := getEncoderConfig()
	return zapcore.NewJSONEncoder(config)
}

func getConsoleEncoder() zapcore.Encoder {
	config := getEncoderConfig()
	e := &ColoredConsoleEncoder{cfg: config, Encoder: zapcore.NewConsoleEncoder(config)}
	return e
}

func getLogWriter(logFilename string) zapcore.WriteSyncer {
	logsPath := "logs"
	if _, err := os.Stat(logsPath); os.IsNotExist(err) {
		os.Mkdir(logsPath, os.ModePerm)
	}
	file, _ := os.OpenFile(filepath.Join(logsPath, logFilename), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	return zapcore.AddSync(file)
}

// custom zap encoder
type ColoredConsoleEncoder struct {
	zapcore.Encoder
	cfg zapcore.EncoderConfig
}

func (e *ColoredConsoleEncoder) Clone() zapcore.Encoder {
	return &ColoredConsoleEncoder{
		// cloning the encoder with the base config
		Encoder: zapcore.NewConsoleEncoder(e.cfg),
		cfg:     e.cfg,
	}
}

// EncodeEntry implementing only EncodeEntry
func (e *ColoredConsoleEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	entry.Message = e.ColoredMessage(entry.Level, entry.Message)

	// calling the embedded encoder's EncodeEntry to keep the original encoding format
	consolebuf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	return consolebuf, nil
}

const HEADER = "\033[95m"
const OKBLUE = "\033[94m"
const OKCYAN = "\033[96m"
const OKGREEN = "\033[92m"
const WARNING = "\033[93m"
const FAIL = "\033[91m"
const ENDC = "\033[0m"
const BOLD = "\033[1m"
const UNDERLINE = "\033[4m"
const GREY = "\x1b[5m\x1b[29m"

var levelToColorMapping = map[zapcore.Level][]string{
	zapcore.DebugLevel:  {OKBLUE},
	zapcore.InfoLevel:   {OKGREEN},
	zapcore.WarnLevel:   {WARNING},
	zapcore.ErrorLevel:  {FAIL},
	zapcore.DPanicLevel: {FAIL, BOLD, UNDERLINE},
	zapcore.PanicLevel:  {FAIL, BOLD, UNDERLINE},
	zapcore.FatalLevel:  {FAIL, BOLD, UNDERLINE},
}

// some mapper function
func (e *ColoredConsoleEncoder) ColoredMessage(lvl zapcore.Level, message string) string {
	colors, ok := levelToColorMapping[lvl]
	if !ok {
		message = " <> " + message + " <> "
		return message
	}
	prefix := strings.Join(colors, "")
	message = prefix + message + ENDC
	return message
}
