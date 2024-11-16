package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/mavcode/qdrant-demo/internal/server"
)

var serverCmd = &cobra.Command{
  Use: "server",
  Short: "short info",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Print(server.Ping())
  },
}

func init() {
  rootCmd.AddCommand(serverCmd)
}

