package categoryctrl

import "pentag.kr/distimer/controllers/subjectctrl"

type categoryDTO struct {
	ID       string                   `json:"id" validate:"required"`
	Name     string                   `json:"name" validate:"required"`
	Order    int8                     `json:"order" validate:"required,min=0"`
	Subjects []subjectctrl.SubjectDTO `json:"subjects" validate:"required"`
}
