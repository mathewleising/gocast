package dialserver

import(
    "fmt"
    "net"
    "os"
    "net/http"
    "io"
)

const(
    ssdpListen string = "239.255.255.250:1900"
    ddxmlTemplate string = "HTTP/1.1 200 OK\r\nLOCATION: http://%s:%d/dd.xml\r\nCACHE-CONTROL: max-age=1800\r\nEXT:\r\nBOOTID.UPNP.ORG: 1\r\nSERVER: Linux/2.6 UPnP/1.0 quick_ssdp/1.0\r\nST: urn:dial-multiscreen-org:service:dial:1\r\nUSN: uuid:%s::urn:dial-multiscreen-org:service:dial:1\r\n\r\n"
    ssdpReplyTemplate string = `<?xml version="1.0" encoding="utf-8"?>
    <root xmlns="urn:schemas-upnp-org:device-1-0" xmlns:r="urn:restful-tv-org:schemas:upnp-dd">
        <specVersion>
        <major>1</major>
        <minor>0</minor>
        </specVersion>
        <URLBase>{{ path }}</URLBase>
        <device>
            <deviceType>urn:schemas-upnp-org:device:dail:1</deviceType>
            <friendlyName>{{ friendlyName }}</friendlyName>
            <manufacturer>Google Inc.</manufacturer>
            <modelName>Eureka Dongle</modelName>
            <UDN>uuid:{{ uuid }}</UDN>
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
    </root>`
)

var(
    ddXML string
    ssdpReply string
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
    ddXML= fmt.Sprintf(ddxmlTemplate, ip, port, uuid, "GiveUsAnAPlus")
    ssdpReply= fmt.Sprintf(ssdpReplyTemplate, ip, port, uuid)
    fmt.Println(ddXML)
    fmt.Println(ssdpReply)
    addr, err := net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
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

    for{
        s.handleMessage()
    }
    http.HandleFunc("/dd.xml", returnDeviceDescriptionXML)
    http.ListenAndServe(fmt.Sprintf("%s:%d", s.httpAddr, s.httpPort), nil)
}

func (s *Server) handleMessage(){
    b := make([]byte, 512)

    n, addr, err := s.conn.ReadFromUDP(b)
    if err != nil || n == 0{
        return
    }

    fmt.Println(addr)
    msg := string(b)
    fmt.Println(msg)
}

  

func returnDeviceDescriptionXML(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, ssdpReplyTemplate)
}

func checkError(err error){
    if err != nil{
        fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
        os.Exit(1)
    }
}