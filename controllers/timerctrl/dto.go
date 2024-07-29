package timerctrl

type timerDTO struct {
	ID        string `json:"id" validate:"required"`
	SubjectID string `json:"subject_id" validate:"required"`
	Content   string `json:"content" validate:"required"`
	StartAt   string `json:"start_at" validate:"required"`
}
