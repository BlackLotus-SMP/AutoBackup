package backup

import (
	"backup/cfg"
	"backup/utils"
	"net/http"
)

type Reload struct {
}

func (r Reload) Reload() utils.Result {
	err := cfg.ReadConfig("config/config.json")
	var res utils.Result
	if err != nil {
		res = utils.Result{Code: http.StatusInternalServerError, Data: "Unable to reload file! " + err.Error()}
	} else {
		res = utils.Result{Code: http.StatusOK, Data: "Config file reloaded!"}
	}
	return res
}
