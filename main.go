package main

import (
    "log"
    "strings"
    "github.com/mathewleising/gocast/SSDP"
    "os/exec"
    "net"
    "fmt"
    "os"
)

func main() {

name, err := os.Hostname()
if err != nil {
fmt.Printf("Oops: %v\n", err)
return
}
addrs, err := net.LookupHost(name)
if err != nil {
fmt.Printf("Oops: %v\n", err)
return
}


    uuid, err := exec.Command("uuidgen").Output()
    if err != nil {
        log.Fatal(err)
    }
    strUuid := strings.Replace(string(uuid), "\n", "", -1)
    s,err := ssdp.NewSSDPServer(addrs[0], 1234, strUuid)
    if err != nil {
        log.Fatal(err)
    }
/*    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }*/
    s.Run()


 /*     for{
        s.handleMessage()
    }

    s.RunDIAL()*/
/*    dialserver.Init()
    dialserver.StartServer()*/
}
