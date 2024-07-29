package groupctrl

type groupDTO struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	NicknamePolicy string `json:"nickname_policy"`
	RevealPolicy   int8   `json:"reveal_policy"`
	InvitePolicy   int8   `json:"invite_policy"`
	CreateAt       string `json:"create_at"`
}

type affiliationDTO struct {
	GroupID  string `json:"group_id"`
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Role     int8   `json:"role"`
	JoinedAt string `json:"joined_at"`
}
