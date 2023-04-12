package dto

type MemberRequest struct {
	Username  string `json:"username" form:"username" validate:"required"`
	Gender    string `json:"gender" form:"gender" validate:"required"`
	SkinType  string `json:"skin_type" form:"skin_type" validate:"required"`
	SkinColor string `json:"skin_color" form:"skin_color" validate:"required"`
}

type MemberRequestUpdate struct {
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	SkinType  string `json:"skin_type" form:"skin_type"`
	SkinColor string `json:"skin_color" form:"skin_color"`
}

type MemberResponse struct {
	IDMember  uint   `json:"id_member"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
	SkinType  string `json:"skin_type"`
	SkinColor string `json:"skin_color"`
}
