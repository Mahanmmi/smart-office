package mqtt

import(
    pahomqtt "github.com/eclipse/paho.mqtt.golang"
	"fmt"
    "log"
    "net/http"
    "io/ioutil"
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
func (s *Server) checkin(cardId string){
    resp, err := http.Get("http://localhost:8080/api/office/checkin?cardid=id&in=true")
    if err != nil {
        log.Fatalln(err)
    }
    //We Read the response body on the line below.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    //Convert the body to type string
    sb := string(body)
    log.Printf(sb)
}
func (s *Server) checkout(cardId string){
    resp, err := http.Get("http://localhost:8080/api/office/checkin?cardid=id&in=false")
    if err != nil {
        log.Fatalln(err)
    }
    //We Read the response body on the line below.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    //Convert the body to type string
    sb := string(body)
    log.Printf(sb)
}