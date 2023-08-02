package logger

import (
	"github.com/zldongly/log"
	"go.uber.org/zap"
)

type zapLog struct {
	*zap.SugaredLogger
}

func NewZapSuger(l *zap.SugaredLogger) log.Log {
	return &zapLog{
		SugaredLogger: l,
	}
}

func (l *zapLog) With(key string, val interface{}) log.Log {
	return &zapLog{
		SugaredLogger: l.SugaredLogger.With(key, val),
	}
}
