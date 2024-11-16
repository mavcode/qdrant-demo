package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/mavcode/qdrant-demo/internal/version"
)

var versionCmd = &cobra.Command{
  Use: "version",
  Short: "version info",
  Run: func(cmd *cobra.Command, args[] string) {
    v, _ := cmd.Flags().GetBool("verbose")
    if v {
      fmt.Println(version.GetAll())
    } else {
      fmt.Println(version.Get())
    }
  },
}

func init() {
  versionCmd.Flags().BoolP("verbose", "v", false, "verbose output")
  rootCmd.AddCommand(versionCmd)
}
