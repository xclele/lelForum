package models

import "time"

// Memory aligned struct(var with same type should be together)
// This can reduce memory usage and improve access speed
type Post struct {
	ID          uint64    `json:"id" db:"post_id"`
	AuthorID    uint64    `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// Use a seperate struct to get authorname and community info
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	*Post                               //Integrate Post info
	*CommunityDetail `json:"community"` //Integrate Community info
}
