package controller

import (
	"fmt"
	"net/http"
)

type HelloReply struct {
	Message string `json:"message"`
}

func (h *Handler) Hello(_ *http.Request, _ *struct{}, reply *HelloReply) (err error) {
	reply.Message = "test"
	return nil
}

type TestArgs struct {
	Message string `json:"message"`
}

type TestReply struct {
	Message string `json:"message"`
}

func (h *Handler) Test(req *http.Request, args *TestArgs, reply *TestReply) error {

	reply.Message = fmt.Sprintf("reply: %v", args.Message)
	return nil
}
