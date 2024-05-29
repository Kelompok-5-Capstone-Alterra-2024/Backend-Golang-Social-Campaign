package dto

import "capstone/entities"

type ApplicationRequest struct {
	VolunteerID uint   `json:"volunteer_id"`
	UserID      uint   `json:"user_id"`
	Status      string `json:"status"`
	IgImage     string `json:"ig_image"`
}

type ApplicationResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func (req *ApplicationRequest) ToEntity() entities.Application {
	return entities.Application{
		VolunteerID: req.VolunteerID,
		UserID:      req.UserID,
		Status:      req.Status,
		IgImage:     req.IgImage,
	}
}
