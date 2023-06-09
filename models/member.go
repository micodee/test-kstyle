package models

type Member struct {
	IDMember      uint          `json:"id_member" gorm:"primary_key: auto_increment"`
	Username      string        `json:"username" gorm:"type: varchar(255)"`
	Gender        string        `json:"gender" gorm:"type: varchar(50)"`
	SkinType      string        `json:"skin_type" gorm:"type: varchar(100)"`
	SkinColor     string        `json:"skin_color" gorm:"type: varchar(100)"`
	ReviewProduct ReviewProduct `json:"review_product" gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LikeReview    LikeReview    `json:"like_review" gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
