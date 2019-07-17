package main

import (
	"log"

	"github.com/adamgoose/ubiquiv"
	"github.com/spf13/cobra"
)

// PORT defines the port on which the TCP server listens
var PORT = "5000"

func main() {
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:   "ubiquiv {dbpath}",
	Short: "UbiquiV is a flexible key-value store, allowing for embedding or centralization.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kv, err := ubiquiv.OpenFile(args[0], "test")
		if err != nil {
			log.Fatalf("Cannot create local database: %s", err.Error())
		}
		defer kv.Close()

		li := kv.(*ubiquiv.LocalInterface)
		li.TCP(":5000")
	},
}
