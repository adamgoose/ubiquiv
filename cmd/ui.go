package main

import (
	"context"
	"log"
	"net/http"

	"github.com/fullstorydev/grpcui/standalone"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var uiCommand = &cobra.Command{
	Use:  "ui",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		server := "localhost:5000"
		if len(args) == 1 {
			server = args[0]
		}

		ctx := context.Background()
		cc, err := grpc.DialContext(ctx, server, grpc.WithBlock(), grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		h, err := standalone.HandlerViaReflection(ctx, cc, server)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Listening on :8080")
		http.ListenAndServe(":8080", h)
	},
}

func init() {
	rootCmd.AddCommand(uiCommand)
}
