package qdrant

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
  "github.com/mavcode/qdrant-demo/internal/config"
)

func Ping() string {
  return "hello from qdrant"
}

type Admin struct {
  client *qdrant.Client
  c *config.Config
}

func NewAdmin(c *config.Config) *Admin {
  return &Admin{c:c}
}

func (a *Admin) ListCollection() {
    collections, err := a.client.ListCollections(context.Background())
    if err != nil {
      panic(err)
    }

    fmt.Println(collections)
}

func (a *Admin) GetConfig() *config.Config {
  return a.c
}

func (a *Admin) Connect() {
    client, err := qdrant.NewClient(&qdrant.Config{
      Host:   a.c.Host,
      Port:   6334,
      APIKey: a.c.APIKey,
      UseTLS: a.c.UseTLS,
    })
    if err != nil {
      panic(err)
    }

    a.client = client
 }

 func (a *Admin) CreateCollection(name string, size uint64) {
   err := a.client.CreateCollection(context.Background(), &qdrant.CreateCollection{
     CollectionName: name,
     VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
       Size: size,
       Distance: qdrant.Distance_Cosine,
     }),
   })

   if err != nil {
     panic(err)
   }
 }

 func (a *Admin) QueryCollection(name string, vector string) {
   searchResult, err := a.client.Query(context.Background(), &qdrant.QueryPoints{
   CollectionName: name,
   Query: qdrant.NewQuery(0.2, 0.1, 0.9, 0.7),
  })

  if err != nil {
    panic(err)
  }
  fmt.Println(searchResult)
 }

