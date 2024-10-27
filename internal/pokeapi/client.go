package pokeapi

import (
  "net/http"
  "time"
  "github.com/Adam-loy/pokedexcli/internal/pokecache"

)

//Client

type Client struct {
  cache      pokecache.Cache
  httpClient http.Client
}

// NewClient

func NewClient(timeout, cachInterval time.Duration) client {
    return Client{
      cache: pokecache.NewCache(cacheInterval),
      httpClient: http.Client{
          Timeout: timeout,
      }
    }

}
