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

var queryCollectionCmd = &cobra.Command{
  Use: "coll",
  Short: "query collection",
  Run: func(cmd *cobra.Command, args []string){
    _name, _ := cmd.Flags().GetString("name")
    _vector, _ := cmd.Flags().GetString("vector")
    _admin.Connect()
    _admin.QueryCollection(_name,_vector)

  },
}

func init() {
  queryCollectionCmd.Flags().StringP("name", "n", "", "foo")
  queryCollectionCmd.Flags().StringP("vector", "v", "", "foo")
  queryCollectionCmd.MarkFlagRequired("name")
  queryCollectionCmd.MarkFlagRequired("vector")
  queryCmd.AddCommand(queryCollectionCmd)
  rootCmd.AddCommand(queryCmd)
}
