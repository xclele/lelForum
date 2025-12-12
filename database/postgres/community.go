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
