package dto

import "capstone/entities"

type ApplicationRequest struct {
	IgImageURL string `json:"ig_image_url" form:"ig_image_url"`
	YtImageURL string `json:"yt_image_url" form:"yt_image_url"`
	Reason     string `json:"reason" form:"reason"`
	Age        int    `json:"age" form:"age"`
	Job        string `json:"job" form:"job"`
}

func (req *ApplicationRequest) ToEntity(igUrl, ytUrl string, userID uint, vacancy_id uint) entities.Application {
	return entities.Application{
		IgImageURL:         igUrl,
		YtImageURL:         ytUrl,
		UserID:             userID,
		VolunteerVacancyID: vacancy_id,
		Reason:             req.Reason,
		Age:                req.Age,
		Job:                req.Job,
	}
}

type ApplicationResponse struct {
	ID         uint   `json:"id"`
	IgImageURL string `json:"ig_image_url"`
	YtImageURL string `json:"yt_image_url"`
	UserID     uint   `json:"user_id"`
	VacancyID  uint   `json:"vacancy_id"`
	Reason     string `json:"reason"`
	Age        int    `json:"age"`
	Job        string `json:"job"`
}

func ToApplicationResponse(application entities.Application) ApplicationResponse {
	return ApplicationResponse{
		ID:         application.ID,
		IgImageURL: application.IgImageURL,
		YtImageURL: application.YtImageURL,
		UserID:     application.UserID,
		VacancyID:  application.VolunteerVacancyID,
		Reason:     application.Reason,
		Age:        application.Age,
		Job:        application.Job,
	}
}
