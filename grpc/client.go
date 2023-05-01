package grpc

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

// var BACKEND_ADDR = string(os.Getenv("BACKEND_GRPC_NODE"))

func init() {
	flag.Parse()
}

const (
	// BACKEND_ADDR = "localhost:8080"
	BACKEND_ADDR = "grpc.weavedb-node.xyz:80"
)

func Query(_method string, _query string, _nocache bool) (*WeaveDBReply, error) {
	log.Printf("INFO: Query, _method=%v, _query=%v, _nocache=%v", _method, _query, _nocache)
	log.Printf("INFO: Query BACKEND_ADDR=%v", BACKEND_ADDR)
	// if backendAddr == "" {
	// 	backendAddr = "grpc.weavedb-node.xyz:80"
	// }

	conn, err := grpc.Dial(BACKEND_ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("ERROR: did not connect: %v", err)
	}
	defer conn.Close()
	log.Printf("INFO: NewDBClient(conn)")
	c := NewDBClient(conn)

	// Contact the server and print out its response.
	timeout := time.Second * 20
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// log.Printf("INFO: c.Query(%v, %v, %v)", _method, _query, _nocache)

	r, err := c.Query(ctx, &WeaveDBRequest{
		Method:  _method,
		Query:   _query,
		Nocache: _nocache,
		// Nocache: true,
	})

	ret := &WeaveDBReply{
		Result: r.GetResult(),
		Err:    r.GetErr(),
	}

	return ret, err
}
