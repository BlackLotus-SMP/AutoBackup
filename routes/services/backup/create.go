package backup

import (
	"backup/cfg"
	"backup/rsync"
	"backup/utils"
	"net/http"
)

type Create struct {
}

func (r Create) Create(name string, conf *cfg.Config, rsyncExecutor *rsync.Executor) utils.Result {
	server, err := conf.GetServer(name)
	var res utils.Result
	if err != nil {
		res = utils.Result{Code: http.StatusNotFound, Data: err.Error()}
	} else {
		res = utils.Result{Code: http.StatusOK, Data: "Starting backup!"}
		backupInstance := rsync.NewRsyncInstance(server)
		rsyncExecutor.StartInstance(backupInstance)
	}
	return res
}
