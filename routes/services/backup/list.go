package backup

import (
	"backup/cfg"
	"backup/utils"
)

type ParsedServers struct {
	Name     string `json:"name"`
	NBackups int    `json:"n_backups"`
}

type List struct {
}

func (l List) List(conf *cfg.Config) utils.Result {
	var parsedServers []ParsedServers
	for _, server := range conf.GetServers() {
		parsedServers = append(parsedServers, ParsedServers{
			Name:     server.Name,
			NBackups: server.NBackups,
		})
	}
	return utils.Result{
		Code: 200,
		Data: parsedServers,
	}
}
