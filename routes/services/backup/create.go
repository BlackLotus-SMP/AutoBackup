package backup

import (
	"backup/cfg"
	"backup/rsync"
	"backup/utils"
	"net/http"
)

type Create struct {
}

func (r Create) Create(name string) utils.Result {
	server, err := cfg.GetServer(name)
	var res utils.Result
	if err != nil {
		res = utils.Result{Code: http.StatusInternalServerError, Data: err.Error()}
	} else {
		res = utils.Result{Code: http.StatusOK, Data: "Starting backup!"}
		backupInstance := rsync.NewRsyncInstance(server)
		rsync.RsyncExecutor.StartInstance(backupInstance)
	}
	return res
}
