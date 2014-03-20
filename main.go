package main

import (
    "strings"
    "github.com/mathewleising/gocast/SSDP"
    "os/exec"
    "net"
    "fmt"
    "flag"
)

var (
    address string
    port int
    uuid string
)

func init() {
    flag.StringVar(&address, "address", getLocalIP(), "IPv4 address for device description, default is local IP")
    flag.IntVar(&port, "port", 1234, "UPnP port for device description, Default is 1234, which is also my default password, shhh")
    flag.StringVar(&uuid, "uuid", getUUID(), "UUID, it's used for something?... Default uses OS UUID, hope you don't mind :)")
}


func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Please forgived me, I panicked :|\n", r)
        }
    }()

    s,err := ssdp.NewSSDPServer(address, port, uuid)
    if err != nil {
        panic(fmt.Sprintf("Oops NewSSDPServer Error: %v\n", err))
    }

    s.Run()
}

func getLocalIP() string {
    tt, err := net.Interfaces() 
    if err != nil { 
        panic(fmt.Sprintf("Oops uuidgen Error: %v\n", err))
    } 

    for _, t := range tt { 
        aa, err := t.Addrs() 
        if err != nil { 
            panic(fmt.Sprintf("Oops uuidgen Error: %v\n", err))
        } 
        for _, a := range aa { 
            ipnet, ok := a.(*net.IPNet) 
            if !ok { 
                continue 
            } 
            v4 := ipnet.IP.To4() 
            if v4 == nil || v4[0] == 127 { // loopback address 
                continue 
            }
            return string(v4)
        } 
    } 
    panic(fmt.Sprintf("Well this is embarrassing.. You started this program without having any IP Addresses open.. Come back when you are ready\n"))
}

func getUUID() string {
    uuid, err := exec.Command("uuidgen").Output()
    if err != nil {
        panic(fmt.Sprintf("Oops uuidgen Error: %v\n", err))
    }

    return strings.Replace(string(uuid), "\n", "", -1)
}
