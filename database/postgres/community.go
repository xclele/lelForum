package postgres

import (
	"database/sql"
	"errors"
	"lelForum/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `SELECT community_id,community_name FROM community`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Error("No community in db!")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (commDetail *models.CommunityDetail, err error) {
	commDetail = new(models.CommunityDetail)
	sqlStr := `SELECT community_id,community_name,introduction,create_time FROM community WHERE community_id=$1`
	if err = db.Get(commDetail, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("Invalid community ID")
		}
	}
	return
}
