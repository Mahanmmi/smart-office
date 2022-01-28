package src

import (
	"github.com/Mahanmmi/fuzzy-lamp/local-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/local-server/db"
)

type LocalServer struct {
	conf       *config.LocalServerConfig
	databases  *db.LocalServerDatabase
}

func NewLocalServer(conf *config.LocalServerConfig, databases *db.LocalServerDatabase) *LocalServer {
	return &LocalServer{
		conf:       conf,
		databases:  databases,
	}
}

func (s *LocalServer) Start() {
	
	// s.httpServer.Start()
}
