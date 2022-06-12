package conf

import (
	"encoding/json"
	"os"

	"github.com/google/wire"
)

var Provider =wire.NewSet(New)

type Config struct {
	Db db `json:"database"`
}

type db struct {
	Dsn string `json:"dsn"`
}

func New() (*Config, error) {
	f, err := os.Open("./app.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
