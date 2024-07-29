package timerctrl

type timerDTO struct {
	ID        string `json:"id"`
	SubjectID string `json:"subject_id"`
	Content   string `json:"content"`
	StartAt   string `json:"start_at"`
}
