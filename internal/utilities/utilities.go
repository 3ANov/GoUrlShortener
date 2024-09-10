package utilities

import (
	"GoUrlShortener/internal/models"
	"encoding/json"
	"os"
)

func GetConfiguration() (models.Configuration, error) {
	config := models.Configuration{}
	file, err := os.Open("./config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
