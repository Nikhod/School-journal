package configs

import (
	"encoding/json"
	"os"
	"second/pkg/models"
)

const (
	info = "info"
	err  = "error"
)

func InitConfigs() (*models.Configs, error) {
	file, err := os.OpenFile("./internal/configs/configs.json", os.O_RDWR, 0444)
	if err != nil {
		return nil, err
	}

	var configs *models.Configs
	err = json.NewDecoder(file).Decode(&configs)
	if err != nil {
		return nil, err
	}

	return configs, nil
}
