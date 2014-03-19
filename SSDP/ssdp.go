package ssdp

import(
    "fmt"
    "net"
    "os"
    "net/http"
    "io"
    "strings"
)

const (
    ssdpListen string = "239.255.255.250:1900"
    //ssdpReplyTemplate string = `HTTP/1.1 200 OK\r\nLOCATION: http://%s:%d/dd.xml\r\nCACHE-CONTROL: max-age=1800\r\nEXT:\r\nBOOTID.UPNP.ORG: 1\r\nSERVER: Linux/2.6 UPnP/1.0 quick_ssdp/1.0\r\nST: urn:dial-multiscreen-org:service:dial:1\r\nUSN: uuid:%s::urn:dial-multiscreen-org:service:dial:1\r\n\r\n`
    ssdpReplyTemplate string =  `HTTP/1.1 200 OK
LOCATION: http://%s:%d/dd.xml
CACHE-CONTROL: max-age=1800
CONFIGID.UPNP.ORG: 7337
BOOTID.UPNP.ORG: 7337
USN: uuid:%s
ST: urn:dial-multiscreen-org:service:dial:1
`

    ddxmlTemplate string = `HTTP/1.1 200 OK
Content-Type: application/xml
Application-URL: http://%s:%d/apps/
<?xml version="1.0" encoding="utf-8"?>
<root xmlns="urn:schemas-upnp-org:device-1-0" xmlns:r="urn:restful-tv-org:schemas:upnp-dd">
    <specVersion>
        <major>1</major>
        <minor>0</minor>
    </specVersion>
    <URLBase>http://%s:%d</URLBase>
    <device>
        <deviceType>urn:schemas-upnp-org:device:dail:1</deviceType>
        <friendlyName>%s</friendlyName>
        <manufacturer>Google Inc.</manufacturer>
        <modelName>Eureka Dongle</modelName>
        <UDN>uuid:%s</UDN>
        <serviceList>
            <service>
                <serviceType>urn:schemas-upnp-org:service:dail:1</serviceType>
                <serviceId>urn:upnp-org:serviceId:dail</serviceId>
                <controlURL>/ssdp/notfound</controlURL>
                <eventSubURL>/ssdp/notfound</eventSubURL>
                <SCPDURL>/ssdp/notfound</SCPDURL>
            </service>
        </serviceList>
    </device>
</root>
    `
)

type Server struct{
    conn *net.UDPConn 
    httpAddr string 
    httpPort int
    uuid string
}

func NewSSDPServer(ip string, port int, uuid string) (*Server, error) {
    s := new(Server)
    s.httpAddr = ip
    s.httpPort = port
    s.uuid = uuid
    //ddXML= fmt.Sprintf(ddxmlTemplate, ip, port, uuid, "GiveUsAnAPlus")
    //ssdpReply= fmt.Sprintf(ssdpReplyTemplate, ip, port, uuid)
    addr, err := net.ResolveUDPAddr("udp4", ssdpListen)
    if err != nil {
          fmt.Println("error from ResolveUDPAddr:", err)
          return nil, err
    }
    socket,err := net.ListenMulticastUDP("udp4", nil, addr)
    if err != nil {
          fmt.Println("error from ListenMulticastUDP:", err)
          return nil, err
    }
    s.conn = socket
    return s, nil
}

func (s *Server) Run(){
    defer func() {
      fmt.Println("Closing connection")
      s.conn.Close()
    }()

    go s.handleMessage()

    http.HandleFunc("/dd.xml", func(w http.ResponseWriter, r *http.Request) {
              returnDeviceDescriptionXML(w, r, fmt.Sprintf(ddxmlTemplate, s.httpAddr , s.httpPort , "GiveUsAnAPlus", s.uuid ))
       })
    http.ListenAndServe(fmt.Sprintf("%s:%d", s.httpAddr, s.httpPort), nil)
}

func (s *Server) handleMessage(){
    for{
    b := make([]byte, 512)

    n, addr, err := s.conn.ReadFromUDP(b)
    if err != nil || n == 0{
        continue
    }

    msg := string(b)
    fmt.Println(msg)

    if strings.Contains(msg, "ST: urn:dial-multiscreen-org:service:dial:1") {
      fmt.Printf("Responding to %s...\n", addr)
      ssdpReply := fmt.Sprintf(ssdpReplyTemplate, s.httpAddr, s.httpPort, s.httpAddr, s.httpPort, s.uuid)
      _,err := s.conn.WriteToUDP([]byte(ssdpReply),addr)
      if err != nil {
        fmt.Println(err) 
      } else {
        fmt.Printf("Responded Successfully.\n")
      }
  }
    }
}

/*func (s *Server) sendMessage(c Client) {
    ddXML= fmt.Sprintf(ddxmlTemplate, ip, port, uuid, "GiveUsAnAPlus")
    for{
        n,err := s.conn.WriteToUDP([]byte(ddXML),c.userAddr)
        fmt.Println(n,err)
    }

}*/

func returnDeviceDescriptionXML(w http.ResponseWriter, req *http.Request, ddxml string) {
    fmt.Println(req)
    io.WriteString(w, ddxml)
}

func checkError(err error){
    if err != nil{
        fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
        os.Exit(1)
    }
}