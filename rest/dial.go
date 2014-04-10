package rest
/*
import(
    "fmt"
    "net"
    "os"
    "net/http"
    "io"
)

type locale struct{
    display_string string
}

type timezone struct{
    display_string string
    offset int
}

type detail struct{
    certificate locale
    timezone timezone
}

type dataSign struct{
    locale string
    nonce string
    signed_data string
}

type ColorGroup struct {
        build_version string
        connected bool
        detail detail
        has_update bool
        hdmi_control bool
        hotspot_bssid string
        locale string
        mac_address string
        name string
        noise_level int
        opt_in:{crash:true,device_id:false,stats:true},
        public_key string
        release_track string
        setup_state int
        sign_data dataSign
        signal_level int
        ssdp_udn:82c5cb87-27b4-2a9a-d4e1-5811f2b1992c,
        ssid:{{ friendlyName }},
        timezone:America/Los_Angeles,
        uptime:0.0,
        version:4,
        wpa_configured:true,
        wpa_state:10
}

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
}*/
