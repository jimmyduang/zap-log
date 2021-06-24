package zap_log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	logger *zap.SugaredLogger
}

func LogInit(path string, isConsole bool, e string) (*Log, error) {
	env := Env__Dev
	err := env.Set(e)
	if err != nil {
		return nil, err
	}
	// env 与 level 关系对应
	level := zapcore.Level(env)

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder.FunctionKey = "func"

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(os.Stdout), level)

	if !isConsole {
		logger := &lumberjack.Logger{
			Filename:   path,
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   false,
			LocalTime:  true,
		}

		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(logger), level)
	}
	return &Log{logger: zap.New(core, zap.AddCallerSkip(1)).Sugar()}, err
}

func (log Log) Debug(arg ...interface{}) {
	log.logger.Debug(arg)
}

func (log Log) Debugf(template string, arg ...interface{}) {
	log.logger.Debugf(template, arg)
}

func (log Log) Info(arg ...interface{}) {
	log.logger.Info(arg)
}

func (log Log) Infof(template string, arg ...interface{}) {
	log.logger.Infof(template, arg)
}
func (log Log) Warn(arg ...interface{}) {
	log.logger.Warn(arg)
}

func (log Log) Warnf(template string, arg ...interface{}) {
	log.logger.Warnf(template, arg)
}
func (log Log) Error(arg ...interface{}) {
	log.logger.Error(arg)
}

func (log Log) Errorf(template string, arg ...interface{}) {
	log.logger.Errorf(template, arg)
}
