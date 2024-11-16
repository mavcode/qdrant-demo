package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/mavcode/qdrant-demo/internal/qdrant"
)

var createCmd = &cobra.Command{
  Use: "create",
  Short: "hello from create",
  Run: func(cmd *cobra.Command, args []string){
    fmt.Print(qdrant.Ping())
  },
}

func init() {
  rootCmd.AddCommand(createCmd)
}
