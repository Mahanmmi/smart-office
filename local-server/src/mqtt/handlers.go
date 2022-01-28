package mqtt

import (
	"fmt"
	"log"

	pahomqtt "github.com/eclipse/paho.mqtt.golang"
)

func (s *Server) connectHandler(client pahomqtt.Client) {
    fmt.Println("Connected to MQTT Broker.")

    fmt.Println("Subscribing to topics")
    s.subscribe("test1")
    fmt.Println()

}
func (s *Server) messageHandler(client pahomqtt.Client, msg pahomqtt.Message){
    fmt.Printf("%s>: %s\n", msg.Topic(),msg.Payload())
	if msg.Topic() == "checkin"{
        s.checkin(string(msg.Payload()))
	}else if msg.Topic() == "checkout" {
        s.checkout(string(msg.Payload())) 
    }else if msg.Topic() == "connect"{
        s.officeConnectionHandler(string(msg.Payload()))
    }
}

func (s *Server) connectionLostHandler(client pahomqtt.Client, err error) {
    fmt.Printf("MQTT Connection lost: %v", err)
}

func (s *Server) publish(topic string, message string) {
	token := s.client.Publish(topic, 0, false, message)
	token.Wait()
}

func (s *Server) subscribe(topic string) {
    token := s.client.Subscribe(topic, 0, nil)
    token.Wait()
  	fmt.Printf("Subscribed to topic: %s", topic)
}
func (s *Server) officeConnectionHandler(clientId string){
    lightSchedule := s.getOfficeLightSchedule()
    log.Println()
    s.publish("lightschedule", lightSchedule)
}