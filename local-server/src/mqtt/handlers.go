package mqtt

import (
	"encoding/json"
	"log"
    "fmt"
	pahomqtt "github.com/eclipse/paho.mqtt.golang"
)

func (s *Server) connectHandler(client pahomqtt.Client) {
	log.Println("Connected to MQTT Broker.")

	s.subscribe("checkin")
	s.subscribe("checkout")
	s.subscribe("connect")
}
func (s *Server) messageHandler(client pahomqtt.Client, msg pahomqtt.Message) {
	log.Printf("%s>: %s\n", msg.Topic(), msg.Payload())
	if msg.Topic() == "checkin" {
		s.checkinHandler(string(msg.Payload()))
	} else if msg.Topic() == "checkout" {
		s.checkoutHandler(string(msg.Payload()))
	} else if msg.Topic() == "connect" {
		s.officeConnectionHandler(string(msg.Payload()))
	}
}

func (s *Server) connectionLostHandler(client pahomqtt.Client, err error) {
	log.Printf("MQTT Connection lost: %v", err)
}

func (s *Server) publish(topic string, message string) {
	token := s.client.Publish(topic, 0, false, message)
	token.Wait()
}

func (s *Server) subscribe(topic string) {
	token := s.client.Subscribe(topic, 0, nil)
	token.Wait()
	log.Printf("Subscribed to topic: %s\n", topic)
}
func (s *Server) checkinHandler(cardId string) {
	resp := s.checkin(cardId)
	var roomSettings map[string]interface{}

	json.Unmarshal([]byte(resp), &roomSettings)
	s.publish("lightintensity", fmt.Sprintf("%v", roomSettings["light"]))
}
func (s *Server) checkoutHandler(cardId string) {
    s.checkout(cardId)
	s.publish("closeroom", "")
}
func (s *Server) officeConnectionHandler(clientId string) {
	if clientId == "ESP8266-f0400d130b6029d61f030b81fdc33a3d28872c70" {
		log.Println("Office is known")
		lightSchedule := s.getOfficeLightSchedule()
		log.Println()
		s.publish("lightschedule", lightSchedule)
	}
}