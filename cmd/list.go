package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd= &cobra.Command{
  Use: "list",
  Short: "list",
}

var listCollectionCmd = &cobra.Command{
  Use: "coll",
  Short: "list collection",
  Run: func(cmd *cobra.Command, args []string) {
    _admin.Connect()
    _admin.ListCollection()
  },
}

func init() {
  listCmd.AddCommand(listCollectionCmd)
  rootCmd.AddCommand(listCmd)
}
