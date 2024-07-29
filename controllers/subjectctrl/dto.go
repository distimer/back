package subjectctrl

type SubjectDTO struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Color string `json:"color" validate:"required,hexcolor"`
}
