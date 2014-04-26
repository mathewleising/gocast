package rest
/*
import(
    "fmt"
    "net"
)

type DialServer struct{
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
    //fmt.Println(ddXML)
    //fmt.Println(ssdpReply)
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
*/
