package src

import (
	"github.com/Mahanmmi/fuzzy-lamp/local-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/local-server/db"
	mqtt "github.com/Mahanmmi/fuzzy-lamp/local-server/src/mqtt"
)

type LocalServer struct {
	conf       *config.LocalServerConfig
	databases  *db.LocalServerDatabase
	mqttServer *mqtt.Server
}

func NewLocalServer(conf *config.LocalServerConfig, databases *db.LocalServerDatabase) *LocalServer {
	return &LocalServer{
		conf:       conf,
		databases:  databases,
		mqttServer: mqtt.NewServer(conf, databases),
	}
}

func (s *LocalServer) Start() {
	s.mqttServer.Start()
}