package main

import (
	"backup/cfg"
	"backup/logger"
	"backup/routes"
	"backup/rsync"
	"backup/utils"
	"flag"
	"github.com/gin-gonic/gin"
	"strings"
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
func IOValidation(log logger.ColorLogger, conf *string) bool {
	path := strings.Split(*conf, "/")
	if len(path) == 0 {
		log.Error("Invalid config path!")
		return false
	}

	file := path[len(path)-1]
	dirs := path[0 : len(path)-1]

	for i, _ := range dirs {
		tempPath := strings.Join(dirs[0:i+1], "/")
		if !utils.DirExists(tempPath) {
			if !utils.TouchDir(tempPath) {
				log.Warning("Unable to Create the %s directory!", tempPath)
				return false
			}
		}
	}

	if !utils.FileExists(*conf) {
		if !utils.TouchFile(*conf) {
			log.Warning("Unable to Create the %s file!", file)
			return false
		}
		cfg.CreateSample(*conf)
	}

	if err := cfg.ReadConfig(*conf); err != nil {
		log.Warning("Unable to Read the %s file!", file)
		return false
	}

	return true
}

// main starts the rsync scheduler and api router, you need to specify a port, or it will default to 8462.
// EX: ./backup -p 8888
func main() {
	var port = flag.String("p", "8462", "Service port")
	var config = flag.String("c", "config/config.json", "Config file path")
	flag.Parse()

	log := logger.NewLogger("AUTO_BACKUP")

	if !IOValidation(log, config) {
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
