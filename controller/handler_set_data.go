package controller

import (
	"fmt"
	"net/http"
)

type SetDataArgs struct {
	Message string `json:"message"`
	// TxID    ids.ID `json:"txId"`
}

type SetDataReply struct {
	Message string `json:"message"`
}

func (h *Handler) SetData(req *http.Request, args *SetDataArgs, reply *SetDataReply) (err error) {
	// reply.Message = "test"
	// ctx, span := h.c.inner.Tracer().Start(req.Context(), "Handler.Balance")
	// defer span.End()

	// reply.Message = fmt.Sprintf("set-data-reply: %v", args.Message)
	h.c.genesis.GlobalString = args.Message
	// storage.SetDataContent(ctx, h.c.metaDB, args.TxID, )
	reply.Message = fmt.Sprintf("SetDataReply: %v", args.Message)
	return nil
}
