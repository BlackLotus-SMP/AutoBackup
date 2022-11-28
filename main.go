package main

import (
	"backup/cfg"
	"backup/logger"
	"backup/routes"
	"backup/rsync"
	"flag"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router        *gin.Engine
	conf          *cfg.Config
	rsyncExecutor *rsync.Executor
}

func (server *Server) init() {
	gin.SetMode(gin.ReleaseMode)
	server.router = gin.Default()
	var router = routes.Loader{}
	for _, route := range router.Load(server.conf, server.rsyncExecutor) {
		route.Route(server.router)
	}
}

// main starts the rsync scheduler and api router, you need to specify a port, or it will default to 8462.
// EX: ./backup -p 8888
func main() {
	var port = flag.String("p", "8462", "Service port")
	var config = flag.String("c", "config/config.json", "Config file path")
	flag.Parse()

	log := logger.NewLogger("AUTO_BACKUP")
	conf, err := cfg.NewConfig(*config, log)

	if err != nil {
		log.Critical("IO Error!")
		log.Critical(err.Error())
		return
	}

	// Start the rsync scheduler.
	rsyncExecutor := rsync.NewExecutor()
	rsyncExecutor.Start()

	server := Server{
		conf:          conf,
		rsyncExecutor: rsyncExecutor,
	}
	server.init()

	log.Info("Starting 0.0.0.0:%s http service", *port)
	// Start the api router.
	err = server.router.Run("0.0.0.0:" + *port)
	if err != nil {
		log.Critical(err.Error())
	}
}
