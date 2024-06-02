package dto

import (
	"capstone/entities"
	"time"
)

type TestimoniVolunteerRequest struct {
	CustomerID  uint   `json:"customer_id"`
	VolunteerID uint   `json:"volunteer_id"`
	Testimoni   string `json:"testimoni"`
	Rating      string `json:"rating"`
	Date        string `json:"date"`
}

func (r *TestimoniVolunteerRequest) ToEntity() (entities.TestimoniVolunteer, error) {
	date, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return entities.TestimoniVolunteer{}, err
	}

	return entities.TestimoniVolunteer{
		CustomerID:  r.CustomerID,
		VolunteerID: r.VolunteerID,
		Testimoni:   r.Testimoni,
		Rating:      r.Rating,
		Date:        date,
	}, nil
}
