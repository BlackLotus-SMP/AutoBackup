package controller

import (
	"backup/cfg"
	"backup/rsync"
)

type Backup struct {
	Conf          *cfg.Config
	RSyncExecutor *rsync.Executor
}
