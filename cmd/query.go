package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/mavcode/qdrant-demo/internal/qdrant"
)

var queryCmd = &cobra.Command{
  Use: "query",
  Short: "hello from query",
  Run: func(cmd *cobra.Command, args []string){
    fmt.Print(qdrant.Ping())
  },
}

func init() {
  rootCmd.AddCommand(queryCmd)
}
