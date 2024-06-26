package dto

import (
	"capstone/entities"
	"fmt"
	"time"
)

type VolunteerRequest struct {
	OrganizationID       uint   `json:"organization_id" form:"organization_id"`
	Title                string `json:"title" form:"title"`
	ContentActivity      string `json:"content_activity" form:"content_activity"`
	Location             string `json:"location" form:"location"`
	LinkWA               string `json:"link_wa" form:"link_wa"`
	StarDate             string `json:"start_date" form:"start_date"`
	EndDate              string `json:"end_date" form:"end_date"`
	TargetVolunteer      int    `json:"target_volunteer" form:"target_volunteer"`
	RegisteredVolunteer  int    `json:"registered_volunteer"`
	RegistrationDeadline string `json:"registration_deadline" form:"registration_deadline"`
	ImageURL             string `json:"image_url" form:"image_url"`
}

func (r *VolunteerRequest) ToEntity(imgUrl string) (entities.Volunteer, error) {
	// loc, err := time.LoadLocation("Asia/Jakarta") // GMT+7 timezone
	// if err != nil {
	// 	return entities.Volunteer{}, fmt.Errorf("failed to load location: %v", err)
	// }

	// startDate, err := time.ParseInLocation("02/01/2006", r.StarDate, loc)
	// if err != nil {
	// 	return entities.Volunteer{}, fmt.Errorf("invalid date format: %v", err)
	// }

	// endDate, err := time.ParseInLocation("02/01/2006", r.EndDate, loc)
	// if err != nil {
	// 	return entities.Volunteer{}, fmt.Errorf("invalid date format: %v", err)
	// }

	// registrationDeadline, err := time.ParseInLocation("02/01/2006", r.RegistrationDeadline, loc)
	// if err != nil {
	// 	return entities.Volunteer{}, fmt.Errorf("invalid registration deadline format: %v", err)
	// }

	startDate, err := time.Parse("2006-01-02", r.StarDate)
	if err != nil {
		return entities.Volunteer{}, fmt.Errorf("Invalid start date format")
	}

	endDate, err := time.Parse("2006-01-02", r.EndDate)
	if err != nil {
		return entities.Volunteer{}, fmt.Errorf("Invalid end date format")
	}

	registrationDeadline, err := time.Parse("2006-01-02", r.RegistrationDeadline)
	if err != nil {
		return entities.Volunteer{}, fmt.Errorf("Invalid registration deadline format")
	}

	return entities.Volunteer{
		OrganizationID:       r.OrganizationID,
		Title:                r.Title,
		ContentActivity:      r.ContentActivity,
		Location:             r.Location,
		LinkWA:               r.LinkWA,
		StartDate:            startDate,
		EndDate:              endDate,
		TargetVolunteer:      r.TargetVolunteer,
		RegistrationDeadline: registrationDeadline,
		ImageURL:             imgUrl,
		Status:               "active",
	}, nil
}

type VolunteerResponse struct {
	ID                  uint                         `json:"id"`
	Organization        VolunteerOrg                 `json:"organization"`
	Title               string                       `json:"title"`
	ContentActivity     string                       `json:"content_activity"`
	Location            string                       `json:"location"`
	LinkWA              string                       `json:"link_wa"`
	StartDate           string                       `json:"start_date"`
	EndDate             string                       `json:"end_date"`
	TargetVolunteer     int                          `json:"target_volunteer"`
	RegisteredVolunteer int                          `json:"registered_volunteer"`
	RegisTionDeadline   string                       `json:"registration_deadline"`
	ImageURL            string                       `json:"image_url"`
	UserRegistered      UserRegisteredResponse       `json:"user_registered"`
	Testimoni           []VolunteerTestimoniResponse `json:"testimoni"`
}

type VolunteerTestimoniResponse struct {
	ID            uint                `json:"id"`
	UserTestimoni UserCommentResponse `json:"user_testimoni"`
	Body          string              `json:"body"`
	Rating        string              `json:"rating"`
	CreatedAt     string              `json:"created_at"`
}

func ToUserCommentResponse(user entities.User) UserCommentResponse {
	return UserCommentResponse{
		UserID:   user.ID,
		Avatar:   user.Avatar,
		Username: user.Username,
	}
}

func ToVolunteerTestimoniResponse(testimoni entities.TestimoniVolunteer) VolunteerTestimoniResponse {
	return VolunteerTestimoniResponse{
		ID:            testimoni.ID,
		UserTestimoni: ToUserCommentResponse(testimoni.User),
		Body:          testimoni.Testimoni,
		Rating:        testimoni.Rating,
		CreatedAt:     testimoni.CreatedAt.Format("2006-01-02"),
	}
}

func ToAllVolunteerTestimoniResponse(testimoni []entities.TestimoniVolunteer) []VolunteerTestimoniResponse {
	result := []VolunteerTestimoniResponse{}
	for _, t := range testimoni {
		result = append(result, ToVolunteerTestimoniResponse(t))
	}

	return result
}

type VolunteerOrg struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	IsVerified bool   `json:"is_verified"`
}

type UserRegisteredResponse struct {
	UserAvatarRegistered     []UserAvatarRegisteredResponse `json:"user_avatar_registered"`
	TotalRegisteredVolunteer int                            `json:"total_registered_volunteer"`
}

type UserAvatarRegisteredResponse struct {
	UserID uint   `json:"user_id"`
	Avatar string `json:"avatar"`
}

func ToVolunteerResponse(volunteer entities.Volunteer, application []entities.Application, testimoni []entities.TestimoniVolunteer) VolunteerResponse {

	uniqueUserAvatars := map[uint]string{}
	for _, app := range application {
		uniqueUserAvatars[app.UserID] = app.User.Avatar
	}

	// Get the avatar of the first four unique user who registered
	userAvatarRegisteredResponse := []UserAvatarRegisteredResponse{}
	for userID, avatar := range uniqueUserAvatars {
		if len(userAvatarRegisteredResponse) == 4 {
			break
		}
		userAvatarRegisteredResponse = append(userAvatarRegisteredResponse, UserAvatarRegisteredResponse{
			UserID: userID,
			Avatar: avatar,
		})
	}

	userRegisteredResponse := UserRegisteredResponse{
		UserAvatarRegistered:     userAvatarRegisteredResponse,
		TotalRegisteredVolunteer: len(application),
	}

	volunteerOrg := VolunteerOrg{
		ID:         volunteer.Organization.ID,
		Name:       volunteer.Organization.Name,
		Avatar:     volunteer.Organization.Avatar,
		IsVerified: volunteer.Organization.IsVerified,
	}

	testimoniResponse := []VolunteerTestimoniResponse{}
	for _, testimoni := range testimoni {
		testimoniResponse = append(testimoniResponse, VolunteerTestimoniResponse{
			UserTestimoni: UserCommentResponse{
				UserID:   testimoni.User.ID,
				Avatar:   testimoni.User.Avatar,
				Username: testimoni.User.Username,
			},
			ID:        testimoni.ID,
			Body:      testimoni.Testimoni,
			Rating:    testimoni.Rating,
			CreatedAt: testimoni.CreatedAt.Format("2006-01-02"),
		})
	}

	return VolunteerResponse{
		ID:                  volunteer.ID,
		Organization:        volunteerOrg,
		Title:               volunteer.Title,
		ContentActivity:     volunteer.ContentActivity,
		Location:            volunteer.Location,
		LinkWA:              volunteer.LinkWA,
		StartDate:           volunteer.StartDate.Format("2006-01-02"),
		EndDate:             volunteer.EndDate.Format("2006-01-02"),
		TargetVolunteer:     volunteer.TargetVolunteer,
		RegisteredVolunteer: volunteer.RegisteredVolunteer,
		RegisTionDeadline:   volunteer.RegistrationDeadline.Format("2006-01-02"),
		ImageURL:            volunteer.ImageURL,
		UserRegistered:      userRegisteredResponse,
		Testimoni:           testimoniResponse,
	}
}

type VolunteersResponses struct {
	ID                   uint   `json:"id"`
	Title                string `json:"title"`
	OrganizationName     string `json:"organization_name"`
	RegisteredVolunteer  int    `json:"registered_volunteer"`
	TargetVolunteer      int    `json:"target_volunteer"`
	RegistrationDeadline string `json:"registration_deadline"`
	ImageUrl             string `json:"image_url"`
}

func ToVolunteersResponses(volunteer entities.Volunteer) VolunteersResponses {

	return VolunteersResponses{
		ID:                   volunteer.ID,
		Title:                volunteer.Title,
		OrganizationName:     volunteer.Organization.Name,
		RegisteredVolunteer:  volunteer.RegisteredVolunteer,
		TargetVolunteer:      volunteer.TargetVolunteer,
		RegistrationDeadline: volunteer.RegistrationDeadline.Format("2006-01-02"),
		ImageUrl:             volunteer.ImageURL,
	}
}

func ToVolunteersResponsesList(volunteers []entities.Volunteer) []VolunteersResponses {
	var res []VolunteersResponses
	for _, volunteer := range volunteers {
		res = append(res, ToVolunteersResponses(volunteer))
	}
	return res
}

type ConfirmVolunteerResponse struct {
	VolunteerID uint   `json:"volunteer_id"`
	UserID      uint   `json:"user_id"`
	ImageURL    string `json:"image_url"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

func ToConfirmVolunteerResponse(volunteer entities.Volunteer, userID uint) ConfirmVolunteerResponse {
	return ConfirmVolunteerResponse{
		VolunteerID: volunteer.ID,
		UserID:      userID,
		ImageURL:    volunteer.ImageURL,
		Title:       volunteer.Title,
		Location:    volunteer.Location,
		StartDate:   volunteer.StartDate.Format("2006-01-02"),
		EndDate:     volunteer.EndDate.Format("2006-01-02"),
	}
}
