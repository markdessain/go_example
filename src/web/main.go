package main

import (
    // "github.com/golang/protobuf/proto"
    "net/http"
    "fmt"
    "time"
    // "io/ioutil"
)

func main() {
    http.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {

      messages := make(chan string)

      go check(messages)
      go check(messages)
      time.Sleep(time.Second)
      go check(messages)

      msg := <-messages
      fmt.Println(msg)

      msg2 := <-messages
      fmt.Println(msg2)

      fmt.Println("Done")
    })

    http.ListenAndServe(":80", nil)
}

func check(messages chan string) {

  time.Sleep(time.Second)
  resp, err := http.Get("http://account/def")

  if err != nil {
     fmt.Println(err)
     return
  }

  fmt.Println(resp)

  messages <- "ping"
}
