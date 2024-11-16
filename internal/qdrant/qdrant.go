package qdrant

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func Ping() string {
  return "hello from qdrant"
}

func ListManagedCollection() {
    client, err := qdrant.NewClient(&qdrant.Config{
      Host:   "86115bde-f557-43ad-aa2c-3305fdf9aa11.us-east4-0.gcp.cloud.qdrant.io",
      Port:   6334,
      APIKey: "-DO0kr7HWAM7Ho1BeaREZ-vCVVPzhw4qWqQ_hFstxMZcGo-hWYoOrg",
      UseTLS: true,
    })
    if err != nil {
      panic(err)
    }

    collections, err := client.ListCollections(context.Background())
    if err != nil {
      panic(err)
    }

    fmt.Println(collections)
}
