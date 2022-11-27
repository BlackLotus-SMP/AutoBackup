package main

import (
	"backup/cfg"
	"backup/routes"
	"backup/rsync"
	"backup/utils"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
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

// main starts the rsync scheduler and api router, you need to specify a port, or it will default to 8462.
// EX: ./backup -p 8888
func main() {
	var port = flag.String("p", "8462", "Service port")
	flag.Parse()

	// IO stuff, create everything related to config and give a sample if not exist as helper for the end user.
	if !utils.DirExists("config") {
		if !utils.TouchDir("config") {
			return
		}
	}

	if !utils.FileExists("config/config.json") {
		if !utils.TouchFile("config/config.json") {
			return
		}
		cfg.CreateSample("config/config.json")
	}
	err := cfg.ReadConfig("config/config.json")
	if err != nil {
		return
	}

	// Start the rsync scheduler.
	rsync.RsyncExecutor = rsync.NewExecutor()
	rsync.RsyncExecutor.Start()

	server := Server{}
	server.init()
	// Start the api router.
	err = server.Router.Run("0.0.0.0:" + *port)
	if err != nil {
		log.Fatal(err)
	}
}
