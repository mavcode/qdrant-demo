package cmd


import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/mavcode/qdrant-demo/internal/qdrant"
)


var uploadCmd = &cobra.Command{
  Use: "upload",
  Short: "upload to collection",
  Run: func(cmd *cobra.Command, args []string){
    fmt.Print(qdrant.Ping())
  },
}

func init() {
  rootCmd.AddCommand(uploadCmd)
}
