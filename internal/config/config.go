package config

type Config struct {
  Host string
  APIKey string
  UseTLS bool
}

func NewConfig() *Config {
  return &Config{
  }
}
