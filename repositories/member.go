package repositories

import (
	"project/models"

	"gorm.io/gorm"
)

type MemberRepository interface {
	FindMember() ([]models.Member, error)
	GetMember(ID int) (models.Member, error)
	CreateMember(member models.Member) (models.Member, error)
	UpdateMember(member models.Member, ID int) (models.Member, error)
	DeleteMember(member models.Member, ID int) (models.Member, error)
}

func RepositoryMember(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMember() ([]models.Member, error) {
	var members []models.Member
	err := r.db.Preload("ReviewProduct").Preload("LikeReview").Find(&members).Error
	return members, err
}

func (r *repository) GetMember(ID int) (models.Member, error) {
	var member models.Member
	err := r.db.Preload("ReviewProduct").Preload("LikeReview").First(&member, ID).Error
	return member, err
}

func (r *repository) CreateMember(member models.Member) (models.Member, error) {
	err := r.db.Create(&member).Error
	return member, err
}

func (r *repository) UpdateMember(member models.Member, ID int) (models.Member, error) {
	err := r.db.Save(&member).Error
	return member, err
}

func (r *repository) DeleteMember(member models.Member, ID int) (models.Member, error) {
	err := r.db.Delete(&member).Error
	return member, err
}
