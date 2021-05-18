package request

type (
	AddProjectRequest struct {
		Name        string  `json:"name" valid:"required"`
		Description string  `json:"description"`
		Logo        string  `json:"logo"`
		Category    string  `json:"category"`
		Level       int     `json:"level" valid:"required"`
		PM          int     `json:"pm"`
		TD          int     `json:"td"`
		Background  string  `json:"background"`
		Worth       string  `json:"worth"`
		Target      string  `json:"target"`
		Milestone   string  `json:"milestone"`
		Budget      float64 `json:"budget"`
		IsShow      int     `json:"is_show" valid:"required"`
	}

	DeleteProjectRequest struct {
		Id int `json:"id" valid:"required"`
	}

	UpdateProjectRequest struct {
		Id          int     `json:"id" valid:"required"`
		Status      int     `json:"status" valid:"required"`
		Name        string  `json:"name" valid:"required"`
		Description string  `json:"description"`
		Logo        string  `json:"logo"`
		Weight      int     `json:"weight" valid:"required"`
		Category    string  `json:"category"`
		Level       int     `json:"level" valid:"required"`
		PM          int     `json:"pm"`
		TD          int     `json:"td"`
		Background  string  `json:"background"`
		Worth       string  `json:"worth"`
		Target      string  `json:"target"`
		Milestone   string  `json:"milestone"`
		Budget      float64 `json:"budget"`
		IsShow      int     `json:"is_show" valid:"required"`
	}

	FindAllProjectRequest struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}

	FindProjectByIdRequest struct {
		Id int `json:"id" valid:"required"`
	}
)
