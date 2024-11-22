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

var createCollectionCmd = &cobra.Command {
  Use: "coll",
  Short: "create collection",
  Run: func(cmd *cobra.Command, args []string) {
    _name, _ := cmd.Flags().GetString("name")
    _size, _ := cmd.Flags().GetInt("size")
    _admin.Connect()
    _admin.CreateCollection(_name, uint64(_size))
  },
}

func init() {
  createCollectionCmd.Flags().StringP("name", "n", "", "collection name")
  createCollectionCmd.Flags().IntP("size", "s", 0, "vector size")
  createCollectionCmd.MarkFlagsRequiredTogether("name","size")
  createCollectionCmd.MarkFlagRequired("name")
  createCollectionCmd.MarkFlagRequired("size")
  createCmd.AddCommand(createCollectionCmd)
  rootCmd.AddCommand(createCmd)
}
