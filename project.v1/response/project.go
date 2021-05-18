package response

import "ict.com/project.v1/model"

type (
	FindProjectByIdResponse struct {
		Code    int32         `json:"code"`
		Message string        `json:"message"`
		Data    model.Project `json:"data"`
	}

	FindAllProjectsResponse struct {
		Code    int32           `json:"code"`
		Message string          `json:"message"`
		Data    []model.Project `json:"data"`
	}
)
