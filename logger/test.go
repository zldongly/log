package logger

import (
	"fmt"
	"github.com/zldongly/log"
	"strings"
	"testing"
)

type testLog struct {
	*testing.T
	withs []string
}

func NewTest(l *testing.T) log.Log {
	return &testLog{
		T:     l,
		withs: []string{},
	}
}

func (l *testLog) With(key string, val interface{}) log.Log {
	w := fmt.Sprintf("%s: %+v", key, val)
	l.withs = append(l.withs, w)

	return l
}

func (l *testLog) with() string {
	if len(l.withs) == 0 {
		return ""
	}
	return "with: {" + strings.Join(l.withs, ",") + "} "
}

func (l *testLog) sprint(args ...interface{}) string {
	if w := l.with(); w != "" {
		args = append([]interface{}{w}, args...)
	}
	return fmt.Sprint(args...)
}

func (l *testLog) sprintf(format string, args ...interface{}) string {
	if w := l.with(); w != "" {
		return w + fmt.Sprintf(format, args...)
	}
	return fmt.Sprintf(format, args...)
}

func (l *testLog) Debug(args ...interface{}) {
	l.T.Log(l.sprint(args...))
}

func (l *testLog) Debugf(format string, args ...interface{}) {
	l.T.Logf(l.with() + fmt.Sprintf(format, args...))
}

func (l *testLog) Info(args ...interface{}) {
	l.T.Log(l.sprint(args...))
}

func (l *testLog) Infof(format string, args ...interface{}) {
	l.T.Log(l.sprintf(format, args...))
}

func (l *testLog) Warn(args ...interface{}) {
	l.T.Log(l.sprint(args...))
}

func (l *testLog) Warnf(format string, args ...interface{}) {
	l.T.Logf(l.sprintf(format, args...))
}

func (l *testLog) Error(args ...interface{}) {
	l.T.Error(l.sprint(args...))
}

func (l *testLog) Errorf(format string, args ...interface{}) {
	l.T.Errorf(l.sprintf(format, args...))
}

func (l *testLog) Fatal(args ...interface{}) {
	l.T.Fatal(l.sprint(args...))
}

func (l *testLog) Fatalf(format string, args ...interface{}) {
	l.T.Fatalf(l.sprintf(format, args...))
}
