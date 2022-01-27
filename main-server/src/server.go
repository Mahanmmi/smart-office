package src

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db"
)

type MainServer struct {
	conf      *config.MainServerConfig
	databases *db.MainServerDatabase
}

func NewMainServer(conf *config.MainServerConfig, databases *db.MainServerDatabase) *MainServer {
	return &MainServer{conf: conf, databases: databases}
}

func (s *MainServer) Start() {
	//TODO start servers
}
