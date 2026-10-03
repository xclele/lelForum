package postgres

import (
	"fmt"
	"lelForum/models"

	"github.com/lib/pq"
)

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

// GetPostList retrieves a list of posts using only post time from the database with pagination
func GetPostList(page, pageSize int64) (posts []*models.Post, err error) {
	sqlStr := `SELECT 
    post_id, author_id, community_id, title, content, create_time
	FROM post
	ORDER BY create_time
	DESC
	LIMIT $1 OFFSET $2`
	posts = make([]*models.Post, 0, pageSize)
	err = db.Select(&posts, sqlStr, pageSize, pageSize*(page-1))
	return
}

// GetPostListByIDs retrieves a list of posts by their IDs from the database
func GetPostListByIDs(ids []string) (posts []*models.Post, err error) {
	// Convert string IDs to int64
	intIds := make([]int64, 0, len(ids))
	for _, idStr := range ids {
		var id int64
		if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
			intIds = append(intIds, id)
		}
	}

	sqlStr := `
		SELECT p.post_id, p.author_id, p.community_id, p.title, p.content, p.create_time
		FROM unnest($1::bigint[]) WITH ORDINALITY AS t(id, ord)
		JOIN post p ON p.post_id = t.id
		ORDER BY t.ord`

	posts = make([]*models.Post, 0, len(ids))
	err = db.Select(&posts, sqlStr, pq.Array(intIds))
	return
}
