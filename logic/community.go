package logic

import (
	"lelForum/database/postgres"
	"lelForum/models"
)

func GetCommunity() ([]*models.Community, error) {
	return postgres.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return postgres.GetCommunityDetailByID(id)
}
