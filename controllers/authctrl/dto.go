package authctrl

type accessInfoDTO struct {
	UserID       string `json:"user_id" validate:"required,uuid"`
	Name         string `json:"name" validate:"required"`
	SessionID    string `json:"session_id" validate:"required,uuid"`
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required,uuid"`
}

type refreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required,uuid"`
}
