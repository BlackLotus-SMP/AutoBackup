package main

import (
	"backup/cfg"
	"backup/logger"
	"backup/routes"
	"backup/rsync"
	"backup/utils"
	"flag"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) init() {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.Default()
	var router = routes.Loader{}
	for _, route := range router.Load() {
		route.Route(server.Router)
	}
}

// IOValidation IO stuff, create everything related to config and give a sample if not exist as helper for the end user.
func IOValidation(log logger.ColorLogger) bool {
	if !utils.DirExists("config") {
		if !utils.TouchDir("config") {
			log.Warning("Unable to Create the config/ directory!")
			return false
		}
	}

	if !utils.FileExists("config/config.json") {
		if !utils.TouchFile("config/config.json") {
			log.Warning("Unable to Create the config.json file!")
			return false
		}
		cfg.CreateSample("config/config.json")
	}

	if err := cfg.ReadConfig("config/config.json"); err != nil {
		log.Warning("Unable to Read the config.json file!")
		return false
	}

	return true
}

// main starts the rsync scheduler and api router, you need to specify a port, or it will default to 8462.
// EX: ./backup -p 8888
func main() {
	var port = flag.String("p", "8462", "Service port")
	flag.Parse()

	log := logger.NewLogger("AUTO_BACKUP")

	if !IOValidation(log) {
		log.Critical("IO Error!")
		return
	}

	// Start the rsync scheduler.
	rsync.RsyncExecutor = rsync.NewExecutor()
	rsync.RsyncExecutor.Start()

	server := Server{}
	server.init()

	log.Info("Starting 0.0.0.0:%s service", *port)
	// Start the api router.
	err := server.Router.Run("0.0.0.0:" + *port)
	if err != nil {
		log.Critical(err.Error())
	}
}
