package logger_test

import (
	"github.com/zldongly/log/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func newZap() (*zap.SugaredLogger, error) {
	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05"))
	}
	logCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		//Sampling *SamplingConfig
		Encoding:         "console",
		EncoderConfig:    encCfg,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		//InitialFields map[string]interface{}
	}

	log, err := logCfg.Build()
	if err != nil {
		return nil, err
	}

	return log.Sugar(), nil
}

func TestZapSuger(t *testing.T) {
	sug, err := newZap()
	if err != nil {
		t.Fatal(err)
	}
	log := logger.NewZapSuger(sug)

	log.Warnf("warn log")

	log = log.With("model", "test").
		With("version", "v1.0.0")
	log = log.With("srv_name", "agriculture")

	log.Error("error log")
}

func TestTestLog(t *testing.T) {
	log := logger.NewTest(t)
	log.Debug("debug log")
	log.Debugf("debugf log")

	log = log.With("version", "v0.1.0").
		With("model", "test")
	log.Info("info log")
	log.Infof("infof log")
}
