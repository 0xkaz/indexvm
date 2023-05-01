package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ava-labs/indexvm/grpc"
)

type TestQueryArgs struct {
	Method  string `json:"method"`
	Query   string `json:"query"`
	Nocache bool   `json:"nocache"`
}

type TestQueryReply struct {
	Message string `json:"message"`
	Result  string `json:"result"`
	Err     string `json:"err"`
}

func (h *Handler) TestQuery(req *http.Request, args *TestQueryArgs, reply *TestQueryReply) error {

	r, err := grpc.Query(args.Method, args.Query, args.Nocache)
	if err != nil {
		log.Printf("TestQuery ERROR: %v", err)
	}
	log.Printf("TestQuery: %v", r.Result)
	reply.Result = r.Result
	reply.Err = r.Err
	reply.Message = fmt.Sprintf("query: %v", args.Method)
	return nil
}
