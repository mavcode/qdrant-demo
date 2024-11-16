package cmd


import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/mavcode/qdrant-demo/internal/qdrant"
)


var adminCmd = &cobra.Command{
  Use: "admin",
  Short: "short from admin",
  Run: func(cmd *cobra.Command, args []string){
    fmt.Print(qdrant.Ping())
    qdrant.ListManagedCollection()
  },
}

func init() {
  rootCmd.AddCommand(adminCmd)
}
