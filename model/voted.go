package model

type VotedData struct {
	UserId    int  `json:"userId,string" binding:"required"`
	Direction int8 `json:"direction,string" binding:"required,oneof=1 0 -1"`
}
