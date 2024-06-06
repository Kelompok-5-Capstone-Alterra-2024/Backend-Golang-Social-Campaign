package dto

import "capstone/entities"

type TestimoniVolunteerRequest struct {
	CustomerID  uint   `json:"customer_id"`
	VolunteerID uint   `json:"volunteer_id"`
	Testimoni   string `json:"testimoni"`
	Rating      string `json:"rating"`
}

func (r *TestimoniVolunteerRequest) ToEntity() entities.TestimoniVolunteer {
	return entities.TestimoniVolunteer{
		CustomerID:  r.CustomerID,
		VolunteerID: r.VolunteerID,
		Testimoni:   r.Testimoni,
		Rating:      r.Rating,
	}
}

type TestimoniVolunteerResponse struct {
	ID          uint   `json:"id"`
	CustomerID  uint   `json:"customer_id"`
	VolunteerID uint   `json:"volunteer_id"`
	Testimoni   string `json:"testimoni"`
	Rating      string `json:"rating"`
}

func ToTestimoniVolunteerResponse(tv entities.TestimoniVolunteer) TestimoniVolunteerResponse {
	return TestimoniVolunteerResponse{
		ID:          tv.ID,
		CustomerID:  tv.CustomerID,
		VolunteerID: tv.VolunteerID,
		Testimoni:   tv.Testimoni,
		Rating:      tv.Rating,
	}
}
