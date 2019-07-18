package main

import (
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
}
