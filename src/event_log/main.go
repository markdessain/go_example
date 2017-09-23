package main

import pb "proto"

import (
    "os"
    // "bufio"
    "github.com/golang/protobuf/proto"
    "net/http"
    "fmt"
    "io/ioutil"
    // "encoding/binary"
    // "bytes"
    // "strconv"
   "github.com/golang/protobuf/jsonpb"
)

const LogFile = "/data/%s.log"

var EventNames =  [2]string{"PageView", "Click"}

func main() {
    http.HandleFunc("/api/log", func(w http.ResponseWriter, r *http.Request) {
      loadEvent(r)
    })
    //
    // http.HandleFunc("/api/v1.0/log", func(w http.ResponseWriter, r *http.Request) {
    //   go check(messages)
    // })
    //
    // http.HandleFunc("/api/v1.1/log", func(w http.ResponseWriter, r *http.Request) {
    //   go check(messages)
    // })
    //
    // http.HandleFunc("/api/v2.0/log", func(w http.ResponseWriter, r *http.Request) {
    //   go check(messages)
    // })

    http.ListenAndServe(":80", nil)
}


func loadEvent(r *http.Request) {
  event_name := r.URL.Query().Get("event_name")

  data, err := ioutil.ReadAll(r.Body)
  if err != nil {
      fmt.Println(err)
  }

  err = writeToFile(data, event_name)
  if err != nil {
      fmt.Println(err)
  }
}

func writeToFile(data []byte, event_name string) error {

  f, err := os.OpenFile(fmt.Sprintf(LogFile, event_name), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
          return err
  }
  defer f.Close()

  fmt.Println("Start")

  marshaller := jsonpb.Marshaler{}

  if event_name == "PageView" {
    event := pb.PageView{}
    if err := proto.Unmarshal(data, &event); err != nil {
        fmt.Println(err)
    }
    err = marshaller.Marshal(f, &event)
    if err != nil {
      fmt.Println(err)
    }

    fmt.Println(event)
    
  } else {
    fmt.Println("Nothing")
  }

  _, err = f.WriteString("\n")
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println("end")
  return nil
}
