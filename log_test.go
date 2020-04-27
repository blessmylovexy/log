package log

import (
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	Config.SetLevel("debug")
	Debug("Debug")
	Debugf("%s", "Debuf")
	Debugw("Debugw", zap.String("Debugw", "Debugw"))
	Info("Info")
	Infof("%s", "Infof")
	Infow("Infow", zap.String("Infow", "Infow"))
	Warn("Warn")
	Warnf("%s", "Warnf")
	Warnw("Warnw", zap.String("Warnw", "Warnw"))
	Error("Error")
	Errorf("%s", "Errorf")
	Errorw("Errorw", zap.String("Errorw", "Errorw"))
	Config.DisableJSONFormat()
	Config.EnableConsoleOut()
	Config.SetProjectName("SetProjectName")
	ApplyConfig()
	Info("SetProjectName")
}
