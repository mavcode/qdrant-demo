package cmd

import (
  "os"
  "github.com/spf13/cobra"
  "github.com/mavcode/qdrant-demo/internal/config"
  "github.com/mavcode/qdrant-demo/internal/qdrant"
)

var _admin *qdrant.Admin
var _c *config.Config

var rootCmd = &cobra.Command{
  Use: "qdrant-demo",
  Short: "foo",
}

func Execute() {
  err := rootCmd.Execute()
  if err != nil {
    os.Exit(1)
  }
}

func init() {
  c := config.NewConfig()
  rootCmd.PersistentFlags().StringVarP(&c.Host,"host" ,"i", "localhost", "qdrant db endpoint")
  rootCmd.PersistentFlags().StringVarP(&c.APIKey,"key", "k", "", "api key")
  rootCmd.PersistentFlags().BoolVarP(&c.UseTLS ,"tls", "t", false, "tls enabled")
  _c = c
  admin := qdrant.NewAdmin(_c)
  _admin = admin
}
