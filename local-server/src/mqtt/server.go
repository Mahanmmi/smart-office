package mqtt

import (
	"time"

	"github.com/Mahanmmi/fuzzy-lamp/local-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/local-server/db"

	"fmt"
	"sync"

	pahomqtt "github.com/eclipse/paho.mqtt.golang"
)

type UserEntity struct{
	Light 			int `json:"light"`
	Room			int `json:"room"`
	UserId			int `json:"user_id"`
	registeredAt	time.Time
}
type Server struct {
	conf       	*config.LocalServerConfig
	databases  	*db.LocalServerDatabase
	client 		pahomqtt.Client
	users 		map[string]UserEntity
}

func NewServer(conf *config.LocalServerConfig, databases *db.LocalServerDatabase) *Server {
	server := Server{
		conf:       conf,
		databases:  databases,
		client: nil,
		users: make(map[string]UserEntity),
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

	//Dear server don't die!
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() 
}