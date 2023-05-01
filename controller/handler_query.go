package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ava-labs/indexvm/grpc"
)

type QueryArgs struct {
	Method  string `json:"method"`
	Query   string `json:"query"`
	Nocache bool   `json:"nocache"`
}

type QueryReply struct {
	Message string `json:"message"`
	Result  string `json:"result"`
	Err     string `json:"err"`
}

func (h *Handler) Query(req *http.Request, args *QueryArgs, reply *QueryReply) error {
	r, err := grpc.Query(args.Method, args.Query, args.Nocache)
	if err != nil {
		log.Printf("Query ERROR: %v", err)
	}
	reply.Result = r.Result
	reply.Err = r.Err
	reply.Message = fmt.Sprintf("query: %v", args.Method)
	return nil
}
