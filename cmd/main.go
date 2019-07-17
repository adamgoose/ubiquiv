package main

import (
	"log"
	"net"

	"github.com/adamgoose/ubiquiv"
	"github.com/spf13/cobra"
)

// PORT defines the port on which the TCP server listens
var PORT = "5000"

func main() {
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "ubiquiv",
	Short: "UbiquiV is a flexible key-value store, allowing for embedding or centralization.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kv, err := ubiquiv.OpenFile(args[0], "test")
		if err != nil {
			log.Fatalf("Cannot create local database: %s", err.Error())
		}
		defer kv.Close()

		// Listen for incoming connections.
		l, err := net.Listen("tcp", "0.0.0.0:"+PORT)
		if err != nil {
			log.Fatal("Error listening:", err.Error())
		}
		// Close the listener when the application closes.
		defer l.Close()
		log.Println("Listening on " + "0.0.0.0:" + PORT)
		for {
			// Listen for an incoming connection.
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("Error accepting: ", err.Error())
			}

			// Handle connections in a new goroutine.
			go func(conn net.Conn) {
				req := ubiquiv.RemoteRequest{}
				if err := req.Decode(conn); err != nil {
					conn.Close()
					return
				}

				// consider doing on the correct bucket?
				value, err := kv.Execute(req, kv)

				resp := ubiquiv.RemoteResponse{
					Value: value,
					Error: err,
				}

				resp.Encode(conn)
				conn.Close()
			}(conn)
		}
	},
}
