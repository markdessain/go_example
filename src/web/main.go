package main

import (
    "net/http"
    "fmt"
     pb "proto"
    "github.com/golang/protobuf/proto"
    "bytes"
    "time"
)

func main() {
    http.HandleFunc("/def", func(w http.ResponseWriter, r *http.Request) {
      event := pb.PageView{
        Uuid:    "1234-1234-1234",
        Vendor:  "com.markkdessain",
        Version: "1-0-0",
        Timestamp: int64(time.Now().UTC().Unix()),
        PageUrl: "test",
      }

      data, err := proto.Marshal(&event)
      if err != nil {
        fmt.Println(err)
        return
      }

      _, err = http.Post("http://event_log/api/log?event_name=PageView", "", bytes.NewBuffer(data))
      if err != nil {
        fmt.Println(err)
        return
      }
    })

    http.ListenAndServe(":80", nil)
}
