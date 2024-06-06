package dto

import "capstone/entities"

type ApplicationRequest struct {
	IgImageURL string `json:"ig_image_url"`
	YtImageURL string `json:"yt_image_url"`
	CustomerID uint   `json:"customer_id"`
	VacancyID  uint   `json:"vacancy_id"`
	Reason     string `json:"reason"`
	Age        int    `json:"age"`
	Job        string `json:"job"`
}

func (req *ApplicationRequest) ToEntity() entities.Application {
	return entities.Application{
		IgImageURL: req.IgImageURL,
		YtImageURL: req.YtImageURL,
		CustomerID: req.CustomerID,
		VacancyID:  req.VacancyID,
		Reason:     req.Reason,
		Age:        req.Age,
		Job:        req.Job,
	}
}

type ApplicationResponse struct {
	ID         uint   `json:"id"`
	IgImageURL string `json:"ig_image_url"`
	YtImageURL string `json:"yt_image_url"`
	CustomerID uint   `json:"customer_id"`
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
		CustomerID: application.CustomerID,
		VacancyID:  application.VacancyID,
		Reason:     application.Reason,
		Age:        application.Age,
		Job:        application.Job,
	}
}
