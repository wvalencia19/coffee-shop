package dtos

import (
	"encoding/json"
	"os"
)

type Config struct {
	NumBaristas int       `json:"numBaristas"`
	Grinders    []Grinder `json:"grinders"`
	Brewers     []Brewer  `json:"brewers"`
	Orders      []Order   `json:"orders"`
}

type Grinder struct {
	GramsPerSecond int `json:"gramsPerSecond"`
}

type Brewer struct {
	OuncesWaterPerSecond int `json:"ouncesWaterPerSecond"`
}

type Order struct {
	OuncesOfCoffeeWanted int `json:"ouncesOfCoffeeWanted"`
}

func ReadConfig(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
