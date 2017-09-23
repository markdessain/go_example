package main

import "proto/event_log"

import (
    "os"
    "net"
    "fmt"
   "github.com/golang/protobuf/jsonpb"
   "golang.org/x/net/context"
   "google.golang.org/grpc"
)

const LogFile = "/data/%s.log"

func main() {
  	lis, err := net.Listen("tcp", ":80")
  	if err != nil {
  		fmt.Println("failed to listen: %v", err)
  	}

  	srv := grpc.NewServer()
  	event_log.RegisterEventLogServer(srv, &server{})
  	srv.Serve(lis)
}

type server struct {
}

// GetProfiles returns hotel profiles for requested IDs
func (s *server) LogPageView(ctx context.Context, req *event_log.PageView) (*event_log.Result, error) {
	res := new(event_log.Result)
  res.Success = true

  marshaller := jsonpb.Marshaler{}
  data, err := marshaller.MarshalToString(req)
  if err != nil {
    fmt.Println(err)
  }

  writeToFile(data, "PageView")
	return res, nil
}

func writeToFile(data string, event_name string) error {

  f, err := os.OpenFile(fmt.Sprintf(LogFile, event_name), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
          return err
  }
  defer f.Close()

  _, err = f.WriteString(data + "\n")
  if err != nil {
    fmt.Println(err)
  }

  return nil
}
