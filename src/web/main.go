package main

import (
    "net/http"
    "fmt"
    "proto/event_log"
    "time"
    "utils"
    "github.com/satori/go.uuid"
)

func main() {
	srv := &server{
		eventLogClient: event_log.NewEventLogClient(utils.MustDial("event_log:80")),
	}

	http.HandleFunc("/", srv.defaultHandler)
	http.ListenAndServe(":80", nil)
}

type server struct {
	eventLogClient event_log.EventLogClient
}

func (s *server) defaultHandler(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()

  event := event_log.PageView{
    Uuid:    uuid.NewV4().String(),
    Vendor:  "com.markkdessain",
    Version: "1-0-0",
    Timestamp: int64(time.Now().UTC().Unix()),
    PageUrl: r.URL.String(),
  }

  nearby, err := s.eventLogClient.LogPageView(ctx, &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  fmt.Println(nearby)
}
