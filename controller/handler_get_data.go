package controller

import (
	"net/http"
)

// type GetDataArgs struct {
// 	// Message string `json:"message"`
// 	// TxID    ids.ID `json:"txId"`
// }

type GetDataReply struct {
	Message string `json:"message"`
}

func (h *Handler) GetData(_ *http.Request, _ *interface{}, reply *GetDataReply) (err error) {
	// reply.Message = "test"
	// ctx, span := h.c.inner.Tracer().Start(req.Context(), "Handler.Balance")
	// defer span.End()

	// reply.Message = fmt.Sprintf("set-data-reply: %v", args.Message)
	reply.Message = h.c.genesis.GlobalString
	// storage.SetDataContent(ctx, h.c.metaDB, args.TxID, )
	// reply.Message = fmt.Sprintf("SetDataReply: %v", args.Message)
	return nil
}
