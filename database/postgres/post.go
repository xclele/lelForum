package postgres

import "lelForum/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO post
    (post_id, author_id, community_id, title, content) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
	return
}
