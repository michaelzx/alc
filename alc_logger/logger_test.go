package alc_logger

import (
	"github.com/michaelzx/alc/alc_config"
	"testing"
)

func TestInit(t *testing.T) {
	Init(alc_config.LoggerConfig{Mode: "dev"})
	Debug("xxxx")
	Info("xxxx")
	Warn("xxxx")
	Error("xxxx")
	DPanic("xxxx")
	Panic("xxxx")
	Fatal("xxxx")
}
