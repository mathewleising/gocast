package rest

import (
  "fmt"
  "io"
  "net/http"
  "github.com/gorilla/mux"
)

const (
  appUrl string = "Application-URL"
  contType string = "Content-Type"
  ddxmlTemplate string = `<?xml version="1.0" encoding="utf-8"?>
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

func (s *Server) StartHttp(){
    r := mux.NewRouter()
    r.NotFoundHandler = http.HandlerFunc(genericHandler)
    r.HandleFunc("/ssdp/device-desc.xml", func(w http.ResponseWriter, r *http.Request) {
              returnDeviceDescriptionXML(w, r, fmt.Sprintf(ddxmlTemplate, s.httpAddr , s.httpPort , "GiveUsAnAPlus", s.uuid ), fmt.Sprintf("http://%s:%d/apps/", s.httpAddr , s.httpPort))
       })
/*
    r.HandleFunc("/setup/{info}", appsHandler)
    r.HandleFunc("/apps", appsHandler)
    r.HandleFunc("/connection", appsHandler)
    r.HandleFunc("/connection/{info}", appsHandler)
    r.HandleFunc("/receiver/{info}", appsHandler)
    r.HandleFunc("/session/{info}", appsHandler)
    r.HandleFunc("/system/control", appsHandler)
*/

    //r.HandleFunc("/products", ProductsHandler)
    //r.HandleFunc("/articles", ArticlesHandler)
    //r.HandleFunc("/products/{key}", ProductHandler)
    //r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
    //r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
    http.Handle("/", r)

    go http.ListenAndServe(fmt.Sprintf("%s:%d", s.httpAddr, s.httpPort), nil)
}

func returnDeviceDescriptionXML(w http.ResponseWriter, req *http.Request, ddxml string, url string) {
    fmt.Println(req)
    w.Header().Set("Content-Type", "application/xml")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Application-URL", url)
    fmt.Printf("Response:\n\n%s\n\n", ddxml)
    io.WriteString(w, ddxml)
}

func genericHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.Error(w, "No Apps running", http.StatusNoContent)
}
