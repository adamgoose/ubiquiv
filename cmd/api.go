package main

import (
	"fmt"

	"github.com/adamgoose/ubivolt"
	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get {key}",
	Short: "Gets a key's value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := ubivolt.OpenServer("localhost:5000", "test")
		defer c.Close()

		fmt.Printf("%s", c.Get([]byte(args[0])))
	},
}

var putCommand = &cobra.Command{
	Use:   "put {key}",
	Short: "Puts a key's value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := ubivolt.OpenServer("localhost:5000", "test")
		defer c.Close()

		if err := c.Put([]byte(args[0]), []byte(args[1])); err != nil {
			fmt.Println(err.Error())
		}
	},
}

var deleteCommand = &cobra.Command{
	Use:   "delete {key}",
	Short: "Deletes a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := ubivolt.OpenServer("localhost:5000", "test")
		defer c.Close()

		if err := c.Delete([]byte(args[0])); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(getCommand)
	rootCmd.AddCommand(putCommand)
	rootCmd.AddCommand(deleteCommand)
}
