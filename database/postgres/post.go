package postgres

import "lelForum/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO post
    (post_id, author_id, community_id, title, content) VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
	return
}

func GetPostByID(pid uint64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := `SELECT post_id, author_id, community_id, title, content, create_time
	FROM post WHERE post_id = $1`
	err = db.Get(data, sqlStr, pid)
	return
}

func GetPostList(page, pageSize int64) (posts []*models.Post, err error) {
	sqlStr := `SELECT 
    post_id, author_id, community_id, title, content, create_time
	FROM post
	LIMIT $1 OFFSET $2`
	posts = make([]*models.Post, 0, pageSize)
	err = db.Select(&posts, sqlStr, pageSize, pageSize*(page-1))
	return
}
