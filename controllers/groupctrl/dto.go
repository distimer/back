package groupctrl

type groupDTO struct {
	ID             string `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description" validate:"required"`
	NicknamePolicy string `json:"nickname_policy" validate:"required"`
	RevealPolicy   int8   `json:"reveal_policy" validate:"required"`
	InvitePolicy   int8   `json:"invite_policy" validate:"required"`
	CreateAt       string `json:"create_at" validate:"required"`
}

type affiliationDTO struct {
	GroupID  string `json:"group_id" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Nickname string `json:"nickname" validate:"required" example:"nickname between 1 and 20"`
	Role     int8   `json:"role" validate:"required,min=0,max=2"`
	JoinedAt string `json:"joined_at" validate:"required"`
}
