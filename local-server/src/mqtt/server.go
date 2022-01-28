package mqtt

import (
	"github.com/Mahanmmi/fuzzy-lamp/local-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/local-server/db"

    pahomqtt "github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"sync"
)

type Server struct {
	conf       *config.LocalServerConfig
	databases  *db.LocalServerDatabase
	client pahomqtt.Client
}

func NewServer(conf *config.LocalServerConfig, databases *db.LocalServerDatabase) *Server {
	server := Server{
		conf:       conf,
		databases:  databases,
		client: nil,
	}
	opts := pahomqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", server.conf.MQTT.Broker, server.conf.MQTT.Port))
    opts.SetClientID(server.conf.MQTT.ClientID)
	opts.SetOnConnectHandler(server.connectHandler)
    opts.SetDefaultPublishHandler(server.messageHandler)
	opts.SetConnectionLostHandler(server.connectionLostHandler)
    server.client = pahomqtt.NewClient(opts)

	return &server
}

func (s *Server) Start() {
	if token := s.client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
	s.subscribe("checkin")
	s.subscribe("checkout")
	s.subscribe("connect")
	
	//Dear server don't die!
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() 
}