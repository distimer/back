package studylogctrl

type myStudyLogDTO struct {
	ID            string   `json:"id"`
	SubjectID     string   `json:"subject_id"`
	StartAt       string   `json:"start_at"`
	EndAt         string   `json:"end_at"`
	Content       string   `json:"content"`
	GroupsToShare []string `json:"groups_to_share"`
}
