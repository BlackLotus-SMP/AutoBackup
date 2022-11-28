package backup

import (
	"backup/cfg"
	"backup/utils"
	"net/http"
)

type Delete struct {
}

func (d Delete) Delete(name string, conf *cfg.Config) utils.Result {
	server, err := conf.GetServer(name)
	var res utils.Result
	if err != nil {
		res = utils.Result{Code: http.StatusNotFound, Data: err.Error()}
	} else {
		err = conf.DeleteServer(server.Name)
		if err != nil {
			res = utils.Result{Code: http.StatusInternalServerError, Data: err.Error()}
		} else {
			res = utils.Result{
				Code: http.StatusOK,
				Data: "Deleted!",
			}
		}
	}
	return res
}
