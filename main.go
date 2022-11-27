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
	Router *gin.Engine
	conf   *cfg.Config
}

func (server *Server) init() {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.Default()
	var router = routes.Loader{}
	for _, route := range router.Load(server.conf) {
		route.Route(server.Router)
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
	rsync.RsyncExecutor = rsync.NewExecutor()
	rsync.RsyncExecutor.Start()

	server := Server{
		conf: conf,
	}
	server.init()

	log.Info("Starting 0.0.0.0:%s service", *port)
	// Start the api router.
	err = server.Router.Run("0.0.0.0:" + *port)
	if err != nil {
		log.Critical(err.Error())
	}
}
