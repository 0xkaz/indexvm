package server

import (
	context "context"
	"log"
	"strings"

	gprc_client "github.com/ava-labs/indexvm/cmd/grpc-server/client"
	proto "github.com/ava-labs/indexvm/cmd/grpc-server/proto"
)

type Server struct {
	proto.UnimplementedDBServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	log.Printf("===")
	log.Printf("RECV: %v", in.GetName())
	message := "Hello, " + in.GetName() + "!"
	log.Printf("SEND: %v", message)
	return &proto.HelloReply{Message: message}, nil
}

func (s *Server) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PingReply, error) {

	log.Printf("RECV: PING")
	message := "PONG"
	log.Printf("SEND: PONG")
	return &proto.PingReply{Message: message}, nil
}

func (s *Server) Query(ctx context.Context, in *proto.WeaveDBRequest) (*proto.WeaveDBReply, error) {

	_method := in.GetMethod()
	_query := in.GetQuery()
	_nocache := in.GetNocache()

	_tmp := strings.Split(_method, "@")

	if len(_tmp) != 2 {
		// ERROR
		log.Printf("ERROR: method parse error")
		return &proto.WeaveDBReply{
			Result: "method parse error",
			Err:    "method parse error",
		}, nil
	}

	// request backend gRPC Server
	log.Printf("INFO request backend gRPC Server")
	return gprc_client.Query(_method, _query, _nocache)
}
