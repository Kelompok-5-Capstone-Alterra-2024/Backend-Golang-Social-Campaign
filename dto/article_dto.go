package dto

import "capstone/entities"

type ArticleRequest struct {
	AdminID  uint   `json:"admin_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}

func (r *ArticleRequest) ToEntity() entities.Article {
	return entities.Article{
		AdminID:  r.AdminID,
		Title:    r.Title,
		Content:  r.Content,
		ImageURL: r.ImageURL,
	}
}
