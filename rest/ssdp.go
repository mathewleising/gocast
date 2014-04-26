package rest

import(
    "fmt"
    "strings"
)

const (
    ssdpReplyTemplate string = `HTTP/1.1 200 OK
LOCATION: http://%s:%d/ssdp/device-desc.xml
CACHE-CONTROL: max-age=1800
CONFIGID.UPNP.ORG: 7337
BOOTID.UPNP.ORG: 7337
USN: uuid:%s
ST: urn:dial-multiscreen-org:service:dial:1
`
)

func (s *Server) StartSSDP(){
    defer func() {
      fmt.Println("Closing connection")
      s.conn.Close()
    }()

    s.handleMessage()
}

func (s *Server) handleMessage(){
    for{
    b := make([]byte, 512)

    n, addr, err := s.conn.ReadFromUDP(b)
    if err != nil || n == 0{
        fmt.Printf("Oops ReadFromUDP Error: %v\n", err)
        continue
    }

    msg := string(b)
    fmt.Println(msg)

    if strings.Contains(msg, "ST: urn:dial-multiscreen-org:service:dial:1") {
      fmt.Printf("Responding to %s...\n", addr)
      ssdpReply := fmt.Sprintf(ssdpReplyTemplate, s.httpAddr, s.httpPort, s.uuid)
      fmt.Printf("Reply:\n\n%s\n\n", ssdpReply)
      _,err := s.conn.WriteToUDP([]byte(ssdpReply),addr)
      if err != nil {
        fmt.Printf("Oops WriteToUDP Error: %v\n", err)
      } else {
        fmt.Printf("Responded Successfully.\n")
      }
  }
    }
}
