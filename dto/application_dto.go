package dto

import "capstone/entities"

type ApplicationRequest struct {
	IgImageURL string `json:"ig_image_url"`
	YtImageURL string `json:"yt_image_url"`
	CustomerID uint   `json:"customer_id"`
	VacancyID  uint   `json:"vacancy_id"`
	Reason     string `json:"reason"`
	Age        int    `json:"age"`
}

func (req *ApplicationRequest) ToEntity() entities.Application {
	return entities.Application{
		IgImageURL: req.IgImageURL,
		YtImageURL: req.YtImageURL,
		CustomerID: req.CustomerID,
		VacancyID:  req.VacancyID,
		Reason:     req.Reason,
		Age:        req.Age,
	}
}
