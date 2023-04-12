package models

type Product struct {
	IDProduct     uint          `json:"id_product" gorm:"primary_key: auto_increment"`
	Name          string        `json:"name_product" gorm:"type: varchar(255)"`
	Price         float64       `json:"price" gorm:"type: float"`
	ReviewProduct ReviewProduct `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
