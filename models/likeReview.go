package models

type LikeReview struct {
	ReviewProductID int `json:"review_id" gorm:"column:review_id"`
	MemberID        int `json:"member_id"`
}
