package mqtt

import(
    "log"
    "net/http"
    "io/ioutil"
)

func (s *Server) checkin(cardId string){
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://localhost:8080/api/office/checkin", nil)
    q := req.URL.Query()
    q.Set("cardid", cardId)
    q.Set("in", "true")
    req.URL.RawQuery = q.Encode()
    req.Header.Add("Authorization", s.conf.OfficeAPIKey)
    log.Println(req)

    resp, err := client.Do(req)

    if err != nil {
        log.Fatalln(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    //Convert the body to type string
    log.Printf(resp.Status)   
    sb := string(body)
    log.Printf(sb)
}
func (s *Server) checkout(cardId string){
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://localhost:8080/api/office/checkin", nil)
    q := req.URL.Query()
    q.Set("cardid", cardId)
    q.Set("in", "true")
    req.URL.RawQuery = q.Encode()
    req.Header.Add("Authorization", s.conf.OfficeAPIKey)
    log.Println(req)

    resp, err := client.Do(req)

    if err != nil {
        log.Fatalln(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    //Convert the body to type string
    sb := string(body)
    log.Printf(resp.Status)   
    log.Printf(sb)   
}
func (s *Server) getOfficeLightSchedule() string{
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://localhost:8080/api/office/lights", nil)
    req.Header.Add("Authorization", s.conf.OfficeAPIKey)
    log.Println(req)

    resp, err := client.Do(req)

    if err != nil {
        log.Fatalln(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    //Convert the body to type string
    sb := string(body)
    log.Printf(resp.Status)   
    log.Printf(sb)   

	return sb
}