package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hexarch",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use -h to show available commands")
	},
}

func Run() {
	rootCmd.AddCommand(restCmd)

	restCmd.Flags().StringP("config", "c", "config.json", "Config file, example file://config.json")
	restCmd.MarkFlagRequired("config")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
