package dto

import "capstone/entities"

type ApplicationRequest struct {
	IgImageURL string `json:"ig_image_url" form:"ig_image_url"`
	YtImageURL string `json:"yt_image_url" form:"yt_image_url"`
	Reason     string `json:"reason" form:"reason"`
	Age        string `json:"age" form:"age"`
	Job        string `json:"job" form:"job"`
}

func (req *ApplicationRequest) ToEntity(igUrl, ytUrl string, userID uint, vacancy_id uint) entities.Application {
	return entities.Application{
		IgImageURL: igUrl,
		YtImageURL: ytUrl,
		UserID:     userID,
		VacancyID:  vacancy_id,
		Reason:     req.Reason,
		Age:        req.Age,
		Job:        req.Job,
	}
}

type ApplicationResponse struct {
	ID           uint   `json:"id"`
	IgImageURL   string `json:"ig_image_url"`
	YtImageURL   string `json:"yt_image_url"`
	UserID       uint   `json:"user_id"`
	UserFullname string `json:"user_fullname"`
	VacancyID    uint   `json:"vacancy_id"`
	Reason       string `json:"reason"`
	Age          string `json:"age"`
	Job          string `json:"job"`
	Status       string `json:"status"`
}

func ToApplicationResponse(application entities.Application) ApplicationResponse {
	return ApplicationResponse{
		ID:           application.ID,
		IgImageURL:   application.IgImageURL,
		YtImageURL:   application.YtImageURL,
		UserID:       application.UserID,
		UserFullname: application.User.Fullname,
		VacancyID:    application.VacancyID,
		Reason:       application.Reason,
		Age:          application.Age,
		Job:          application.Job,
		Status:       application.Volunteer.Status,
	}
}

func ToApplicationsResponse(applications []entities.Application) []ApplicationResponse {
	var applicationsResponse []ApplicationResponse
	for _, application := range applications {
		applicationsResponse = append(applicationsResponse, ToApplicationResponse(application))
	}
	return applicationsResponse
}
