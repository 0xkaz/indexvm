package server

import (
	context "context"
	"log"
	"strings"
	"time"

	// helpers "github.com/ava-labs/indexvm/cmd/grpc-server/helpers"

	// "weavedb-gprc-cache-proxy/app/libs/mycache"
	client "github.com/ava-labs/indexvm/cmd/grpc-server/client"
	proto "github.com/ava-labs/indexvm/cmd/grpc-server/proto"
)

type Server struct {
	proto.UnimplementedDBServer
}

const (
	CachePrefix = "weavedb"
)

var ttl = time.Duration(3600)

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

	// _func := _tmp[0]
	// _contractTxId := _tmp[1]
	// _collectionNames, err := helpers.ParseCollectionName(_func, _query)

	// // logParams["func"] = _func
	// // logParams["contractTxId"] = _contractTxId

	// if err != nil || len(_collectionNames) == 0 {
	// 	// call back end server
	// 	log.Printf("INFO: call backend cuz collectionName Err: %v", err)
	// 	return weavedb_client.Query(_method, _query, _nocache)
	// }

	// // log.Printf("INFO: func: %v, contractTxId:%v, collectionName: %v", _func, _contractTxId, _collectionName)
	// var cacheKeyPrefs []string
	// for _, _name := range _collectionNames {
	// 	pref := fmt.Sprintf("%s.%s.%x", CachePrefix, _contractTxId, md5.Sum([]byte(_name)))
	// 	cacheKeyPrefs = append(cacheKeyPrefs, pref)
	// }
	// cacheKey := fmt.Sprintf("%s.%x.%x", cacheKeyPrefs[0], md5.Sum([]byte(_func)), md5.Sum([]byte(_query)))
	// log.Printf("cacheKey: %v", cacheKey)
	// // log.Printf("INFO: cacheKeyPref: %s", cacheKeyPref)
	// // log.Printf("INFO: cacheKey: %s", cacheKey)
	// // log.Printf("nocache: %v", _nocache)
	// // if !_nocache {
	// // 	if _func == "add" ||
	// // 		_func == "set" ||
	// // 		_func == "upsert" ||
	// // 		_func == "update" {
	// // 		// DELETE CACHE
	// // 		log.Printf("INFO: delete cache")
	// // 		if err := mycache.DelCachePref(cacheKeyPrefs[0]); err != nil {
	// // 			log.Printf("ERROR: mycache.DelCachePref: err=%v", err)
	// // 		}

	// // 		// request backend gRPC Server
	// // 		log.Printf("INFO: request backend gRPC Server")
	// // 		return weavedb_client.CallWeaveDBQuery(_method, _query, _nocache)

	// // 	} else if _func == "get" || _func == "cget" {

	// // 		// log.Printf("INFO: try to get cache. cacheKey=%v",cacheKey)
	// // 		// cret,err := GetSerializedCache(cacheKey)
	// // 		cret, err := mycache.GetCache(cacheKey)
	// // 		// log.Printf("INFO: cret=%v", cret)
	// // 		if (err == nil) && cret != "" {
	// // 			log.Printf("INFO: CACHE HIT. cacheKey=%v", cacheKey)
	// // 			return &proto.WeaveDBReply{
	// // 				Result: cret,
	// // 				Err:    "",
	// // 			}, nil
	// // 		}

	// // 		log.Printf("INFO: NO CACHE key=%v", cacheKey)
	// // 		ret, err := weavedb_client.CallWeaveDBQuery(_method, _query, _nocache)
	// // 		if err == nil && ret.GetErr() == "" {
	// // 			// log.Printf("INFO: SAVE CACHE cacheKey=%v", cacheKey)
	// // 			err = mycache.SetCache(cacheKey, ret.GetResult(), ttl)
	// // 			if err != nil {
	// // 				log.Printf("ERROR: SAVE CACHE ERROR. error=%v", err)
	// // 			}
	// // 		}
	// // 		log.Printf("INFO: return")
	// // 		return &proto.WeaveDBReply{
	// // 			Result: ret.GetResult(),
	// // 			Err:    ret.GetErr(),
	// // 		}, nil
	// // 	} else if _func == "setRules" {
	// // 		log.Printf("INFO: delete cache")
	// // 		if err := mycache.DelCachePref(cacheKeyPrefs[0]); err != nil {
	// // 			log.Printf("ERROR: mycache.DelCachePref: err=%v", err)
	// // 		}
	// // 		return weavedb_client.CallWeaveDBQuery(_method, _query, _nocache)

	// // 	} else if _func == "listCollections" {
	// // 		log.Printf("listCollections.....")
	// // 		cret, err := mycache.GetCache(cacheKey)
	// // 		// log.Printf("INFO: cret=%v", cret)
	// // 		if (err == nil) && cret != "" {
	// // 			log.Printf("INFO: CACHE HIT.. cacheKey=%v", cacheKey)
	// // 			return &proto.WeaveDBReply{
	// // 				Result: cret,
	// // 				Err:    "",
	// // 			}, nil
	// // 		}
	// // 		log.Printf("INFO: NO CACHE.. key=%v", cacheKey)
	// // 		ret, err := weavedb_client.CallWeaveDBQuery(_method, _query, _nocache)
	// // 		if err == nil && ret.GetErr() == "" {
	// // 			err = mycache.SetCache(cacheKey, ret.GetResult(), 0)
	// // 			if err != nil {
	// // 				log.Printf("ERROR: SAVE CACHE ERROR. error=%v", err)
	// // 			}
	// // 		}
	// // 		log.Printf("INFO: return")
	// // 		return &proto.WeaveDBReply{
	// // 			Result: ret.GetResult(),
	// // 			Err:    ret.GetErr(),
	// // 		}, nil

	// // 	} else if _func == "batch" {
	// // 		log.Printf("INFO: batch: cacheKeyPrefs: %v", cacheKeyPrefs)
	// // 		for _, pref := range cacheKeyPrefs {
	// // 			// mycache.DelCachePref(pref)
	// // 			if err := mycache.DelCachePref(pref); err != nil {
	// // 				log.Printf("ERROR: mycache.DelCachePref: err=%v", err)
	// // 			}
	// // 		}

	// // 	}

	// // }

	// // other functions

	// request backend gRPC Server
	log.Printf("INFO request backend gRPC Server")
	return client.Query(_method, _query, _nocache)
}
