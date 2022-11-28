package cfg

import (
	"backup/logger"
	"backup/utils"
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type Config struct {
	configPath string
	servers    []Server
	logger     logger.ColorLogger
}

type Server struct {
	Name          string `json:"name"`            // Unique name ID.
	SSHRemotePath string `json:"ssh_remote_path"` // SSH path to download where the world/ dir is located.
	SSHUser       string `json:"ssh_user"`        // SSH user.
	SSHPass       string `json:"ssh_pass"`        // SSH pass.
	LocalPath     string `json:"local_path"`      // Path where we want the copy to be stored.
	NBackups      int    `json:"n_backups"`       // Number of maximum backups that can be stored.
}

func NewConfig(configPath string, logger logger.ColorLogger) (*Config, error) {
	cfg := &Config{
		configPath: configPath,
		servers:    []Server{},
		logger:     logger,
	}
	if !cfg.validatePaths() {
		return cfg, errors.New("error on IO validation")
	}
	return cfg, nil
}

func (cfg *Config) GetServers() []Server {
	return cfg.servers
}

func (cfg *Config) validatePaths() bool {
	path := strings.Split(cfg.configPath, "/")
	if len(path) == 0 {
		cfg.logger.Error("Invalid config path!")
		return false
	}

	file := path[len(path)-1]
	dirs := path[0 : len(path)-1]
	dirPath := strings.Join(dirs, "/")

	if !utils.DirExists(dirPath) && !utils.TouchDir(dirPath) {
		cfg.logger.Warning("Unable to create %s!", dirPath)
		return false
	}

	if !utils.FileExists(cfg.configPath) {
		if !utils.TouchFile(cfg.configPath) {
			cfg.logger.Warning("Unable to Create the %s file!", file)
			return false
		}
		cfg.createSample(cfg.configPath)
	}

	if err := cfg.ReadConfig(); err != nil {
		cfg.logger.Warning("Unable to Read the %s file!", file)
		return false
	}

	return true
}

func (cfg *Config) ReadConfig() error {
	data, err := os.ReadFile(cfg.configPath)
	if err != nil {
		return err
	}

	var servers []Server
	err = json.Unmarshal(data, &servers)
	if err != nil {
		return err
	}

	cfg.servers = servers
	return nil
}

// GetServer iterates the Server array to get the credentials of the server based on endpoint name.
func (cfg *Config) GetServer(name string) (Server, error) {
	for _, server := range cfg.servers {
		if server.Name == name {
			return server, nil
		}
	}
	return Server{}, errors.New("server not found")
}

// CreateSample if the config json file does not exist, it will create a default one.
func (cfg *Config) createSample(path string) {
	cfg.servers = []Server{
		{
			Name:          "test",
			SSHRemotePath: "1.2.3.4:/home/test/bck/",
			SSHUser:       "user",
			SSHPass:       "pass",
			LocalPath:     "/home/bck/",
			NBackups:      5,
		},
	}
	jsonString, err := json.MarshalIndent(cfg.servers, "", "    ")
	if err != nil {
		return
	}
	err = os.WriteFile(path, jsonString, 0644)
	if err != nil {
		return
	}
}
