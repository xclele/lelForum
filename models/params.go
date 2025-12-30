package models

/*
Here the binding tags are used for validation when binding JSON input,
that's a feature of the Gin framework.
*/
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,min=3,max=20"`
	Password   string `json:"password" binding:"required,min=3,max=20"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password,min=3,max=20"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`               // Post ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1" ` // Vote(1) or downvote(-1)
}
