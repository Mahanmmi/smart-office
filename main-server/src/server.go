package src

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/src/http"
)

type MainServer struct {
	conf       *config.MainServerConfig
	databases  *db.MainServerDatabase
	httpServer *http.Server
}

func NewMainServer(conf *config.MainServerConfig, databases *db.MainServerDatabase) *MainServer {
	return &MainServer{
		conf:       conf,
		databases:  databases,
		httpServer: http.NewServer(conf, databases),
	}
}

func (s *MainServer) Start() {
	s.httpServer.Start()
}
