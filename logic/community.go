package logic

import (
	"lelForum/database/postgres"
	"lelForum/models"
)

func GetCommunity() ([]*models.Community, error) {
	return postgres.GetCommunityList()
}
