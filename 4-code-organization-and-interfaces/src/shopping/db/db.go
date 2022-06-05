package db

import "shopping/models"

func LoadItem(id int) *models.Item {
	return &models.Item{
		Name:  "mochi",
		Price: 6.10,
	}
}
