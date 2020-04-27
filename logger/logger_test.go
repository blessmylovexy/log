package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	wlog := New()
	wlog.Config.SetLevel("debug")
	wlog.ApplyConfig()
	wlog.Debug("Debug")
	wlog.Debugf("%s", "Debuf")
	wlog.Debugw("Debugw", zap.String("Debugw", "Debugw"))
	wlog.Info("Info")
	wlog.Infof("%s", "Infof")
	wlog.Infow("Infow", zap.String("Infow", "Infow"))
	wlog.Warn("Warn")
	wlog.Warnf("%s", "Warnf")
	wlog.Warnw("Warnw", zap.String("Warnw", "Warnw"))
	wlog.Error("Error")
	wlog.Errorf("%s", "Errorf")
	wlog.Errorw("Errorw", zap.String("Errorw", "Errorw"))
	wlog.Config.DisableJSONFormat()
	wlog.Config.EnableConsoleOut()
	wlog.Config.SetProjectName("SetProjectName")
	wlog.Config.SetStacktraceLevel("info")
	wlog.ApplyConfig()
	wlog.Info("SetProjectName")
}
