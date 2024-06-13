package dto

import "capstone/entities"

type TestimoniVolunteerRequest struct {
	Testimoni string `json:"testimoni"`
	Rating    string `json:"rating"`
}

func (r *TestimoniVolunteerRequest) ToEntity(volunteer_id uint, user_id uint) entities.TestimoniVolunteer {
	return entities.TestimoniVolunteer{
		UserID:    user_id,
		VacancyID: volunteer_id,
		Testimoni: r.Testimoni,
		Rating:    r.Rating,
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
		CustomerID:  tv.UserID,
		VolunteerID: tv.VacancyID,
		Testimoni:   tv.Testimoni,
		Rating:      tv.Rating,
	}
}
