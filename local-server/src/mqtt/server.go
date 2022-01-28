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
	mqtt *pahomqtt.ClientOptions
	client pahomqtt.Client
}

func NewServer(conf *config.LocalServerConfig, databases *db.LocalServerDatabase) *Server {
	server := Server{
		conf:       conf,
		databases:  databases,
		mqtt: pahomqtt.NewClientOptions(),
		client: nil,
	}
    server.mqtt.AddBroker(fmt.Sprintf("tcp://%s:%d", server.conf.MQTT.Broker, server.conf.MQTT.Port))
    server.mqtt.SetClientID(server.conf.MQTT.ClientID)
	server.mqtt.SetOnConnectHandler(server.connectHandler)
    server.mqtt.SetDefaultPublishHandler(server.messageHandler)
	server.mqtt.SetConnectionLostHandler(server.connectionLostHandler)
    server.client = pahomqtt.NewClient(server.mqtt)

	return &server
}

func (s *Server) Start() {

	getAuthenticatedUsers()	
	//Start MQTT
	if token := s.client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
	
	//Dear server don't die!
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() 
}
func getAuthenticatedUsers(){
	
}