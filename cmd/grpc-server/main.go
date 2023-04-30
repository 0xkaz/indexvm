package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	// gutils "github.com/ava-labs/indexvm/cmd/grpc-server/utils"
	proto "github.com/ava-labs/indexvm/cmd/grpc-server/proto"
	server "github.com/ava-labs/indexvm/cmd/grpc-server/server"
	gutils "github.com/ava-labs/indexvm/cmd/grpc-server/utils"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc/reflection"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	"google.golang.org/grpc"
)

func main() {
	//
	buildtime := gutils.GetBuildTime()
	log.Printf("build: %v", buildtime)

	// prepare listeners with cmux
	lis, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Fatalf("ERROR: failed to listen: %v", err)
	}
	m := cmux.New(lis)
	grpcListener := m.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"),
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc+proto"),
	)
	anyListener := m.Match(cmux.Any()) // http & grpc-web & ANY

	// // gRPC server with gRPC-web Wrapper
	opts := GetGrpcServerOpts()
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	proto.RegisterDBServer(grpcServer, &server.Server{})
	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithWebsockets(true))
	// ///////////////////////////////////

	// Serve gRPC
	go func(l net.Listener, s *grpc.Server) {
		log.Printf("INFO: grpc start ")
		// s.Serve(l)
		if err = s.Serve(l); err != nil {
			log.Fatalf("ERROR: failed to grpc listen : %v", err)
		}
	}(grpcListener, grpcServer)

	// Serve HTTP and gRPC-Web and ANY
	go func(l net.Listener, wrappedServer http.Handler) {

		////////////////////////
		httpapp := fiber.New(fiber.Config{
			ServerHeader: "Fiber",
		})

		httpapp.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))

		httpapp.Use(cors.New(cors.Config{
			// AllowOrigins: "https://gofiber.io, https://gofiber.net",
			AllowOrigins: "*",
			AllowHeaders: "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message, Accept, Accept-Language, ",
			AllowMethods: "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT",
			// AllowHeaders: "*",
		}))

		// For gRPC-web
		httpapp.Use(adaptor.HTTPMiddleware(func(next http.Handler) http.Handler {

			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// log.Println("log middleware")
				if strings.Contains(r.Header.Get("Content-Type"), "application/grpc-web") ||
					strings.Contains(r.Header.Get("Access-Control-Request-Headers"), "x-grpc-web") ||
					r.Header.Get("x-grpc-web") == "1" ||
					r.Header.Get("X-Grpc-Web") == "1" ||
					false {

					log.Printf("grpc-web")

					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
					w.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
					w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
					wrappedServer.ServeHTTP(w, r)
					return
				}
				next.ServeHTTP(w, r)
			})

		}))

		httpapp.Get("/", func(c *fiber.Ctx) error {
			buildtime := gutils.GetBuildTime()
			sincebuild := gutils.GetHowOldInSec()
			return c.SendString(fmt.Sprintf("grpc-server (build: %v; %d sec ago)", buildtime, sincebuild))
		})
		log.Printf("INFO: http/grpc-web/any start ")
		if err := httpapp.Listener(l); err != nil {
			log.Fatalf("ERROR: failed to gin-http listen: %v", err)
		}
	}(anyListener, wrappedServer)

	if err = m.Serve(); err != nil {
		log.Fatalf("ERROR: failed to cmux listen: %v", err)
		log.Fatal(err)
	}
}
