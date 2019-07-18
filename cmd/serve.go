package main

import (
	"fmt"
	"log"
	"net"

	"github.com/adamgoose/ubivolt"

	proto "github.com/adamgoose/ubivolt/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serveCommand = &cobra.Command{
	Use:  "serve",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kv, err := ubivolt.OpenFile(args[0], "test")
		if err != nil {
			log.Fatalf("Cannot create local database: %s", err.Error())
		}
		defer kv.Close()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()

		proto.RegisterKeyValueServer(grpcServer, ubivolt.GRPCServer{
			DB: kv.(*ubivolt.LocalInterface).DB,
		})

		reflection.Register(grpcServer)
		log.Println("Listening on :5000")
		grpcServer.Serve(lis)
	},
}

func init() {
	rootCmd.AddCommand(serveCommand)
}
