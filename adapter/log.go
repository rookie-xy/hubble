package adapter

import (
	"github.com/rookie-xy/hubble/log"
	"github.com/rookie-xy/hubble/log/level"
)

type LevelLog interface {
	log.Log
	Get() level.Level
	Copy(logger *log.Logger)
}

func ToLevelLog(log log.Log) LevelLog {
	return log.(LevelLog)
}
