package controller

import (
	"net/http"
)

type HelloReply struct {
	Message string `json:"message"`
}

func (h *Handler) Hello(_ *http.Request, _ *struct{}, reply *HelloReply) (err error) {
	reply.Message = "test"
	return nil
}
