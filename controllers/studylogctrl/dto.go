package studylogctrl

import "pentag.kr/distimer/controllers/subjectctrl"

type myStudyLogDTO struct {
	ID            string   `json:"id" validate:"required"`
	SubjectID     string   `json:"subject_id" validate:"required"`
	StartAt       string   `json:"start_at" validate:"required"`
	EndAt         string   `json:"end_at" validate:"required"`
	Content       string   `json:"content" validate:"required"`
	GroupsToShare []string `json:"groups_to_share" validate:"required"`
}

type groupStudyLogDTO struct {
	ID           string                 `json:"id" validate:"required"`
	Subject      subjectctrl.SubjectDTO `json:"subject" validate:"required"`
	CategoryID   string                 `json:"category_id" validate:"required"`
	CategoryName string                 `json:"category_name" validate:"required"`
	StartAt      string                 `json:"start_at" validate:"required"`
	EndAt        string                 `json:"end_at" validate:"required"`
	Content      string                 `json:"content" validate:"required"`
}
