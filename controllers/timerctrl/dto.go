package timerctrl

import (
	"pentag.kr/distimer/controllers/groupctrl"
	"pentag.kr/distimer/controllers/subjectctrl"
)

type timerDTO struct {
	ID             string   `json:"id" validate:"required"`
	SubjectID      string   `json:"subject_id" validate:"required"`
	Content        string   `json:"content" validate:"required"`
	StartAt        string   `json:"start_at" validate:"required"`
	SharedGroupIDs []string `json:"shared_group_ids" validate:"required"`
}

type timerWithEdgeInfoDTO struct {
	ID          string                   `json:"id" validate:"required"`
	Subject     subjectctrl.SubjectDTO   `json:"subject" validate:"required"`
	Content     string                   `json:"content" validate:"required"`
	StartAt     string                   `json:"start_at" validate:"required"`
	Affiliation groupctrl.AffiliationDTO `json:"affiliation" validate:"required"`
}

type timerMetadataDTO struct {
	SubjectID      string   `json:"subject_id" validate:"required,uuid"`
	Content        string   `json:"content" validate:"required" example:"content between 0 and 30"`
	SharedGroupIDs []string `json:"shared_group_ids" validate:"required"`
}
