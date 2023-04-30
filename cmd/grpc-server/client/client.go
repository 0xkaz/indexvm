package client

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	proto "github.com/ava-labs/indexvm/cmd/grpc-server/proto"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

var backendAddr = string(os.Getenv("BACKEND_GRPC_NODE"))

func init() {
	flag.Parse()
}

func Query(_method string, _query string, _nocache bool) (*proto.WeaveDBReply, error) {
	if backendAddr == "" {
		backendAddr = "grpc.weavedb-node.xyz:80"
	}
	log.Printf("INFO: Query backendAddr=%v", backendAddr)
	conn, err := grpc.Dial(backendAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("ERROR: did not connect: %v", err)
	}
	defer conn.Close()
	log.Printf("INFO: NewDBClient(conn)")
	c := proto.NewDBClient(conn)

	// Contact the server and print out its response.
	timeout := time.Second * 20
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// log.Printf("INFO: c.Query(%v, %v, %v)", _method, _query, _nocache)

	r, err := c.Query(ctx, &proto.WeaveDBRequest{
		Method: _method,
		Query:  _query,
		// Nocache: _nocache,
		Nocache: true,
	})

	ret := &proto.WeaveDBReply{
		Result: r.GetResult(),
		Err:    r.GetErr(),
	}
	// log.Printf("INFO: ret: %v", ret)

	return ret, err
}
