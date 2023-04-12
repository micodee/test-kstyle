package models

type ReviewProduct struct {
	IDReview   uint       `json:"id_review" gorm:"primaryKey"`
	MemberID   int        `json:"member_id"`
	ProductID  int        `json:"product_id"`
	DescReview string     `json:"desc_review" gorm:"type: text"`
	LikeReview LikeReview `json:"-" gorm:"foreignKey:ReviewProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}