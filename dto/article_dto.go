package dto

import (
	"capstone/entities"
	"time"
)

type ArticleResponse struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	AdminID   uint   `json:"admin_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageURL  string `json:"image_url"`
}

func ToArticleResponse(article entities.Article) ArticleResponse {
	loc, _ := time.LoadLocation("Asia/Jakarta") // GMT+7 timezone

	createdAt := article.CreatedAt.In(loc).Format("02/01/2006 15:04:05 MST") // Example format: 02/01/2006 15:04:05 MST
	return ArticleResponse{
		ID:        article.ID,
		CreatedAt: createdAt,
		Title:     article.Title,
		Content:   article.Content,
		ImageURL:  article.ImageURL,
	}
}

type ArticleRequest struct {
	AdminID  uint   `json:"admin_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}

func ToArticleResponseList(articles []entities.Article) []ArticleResponse {
	var response []ArticleResponse
	for _, article := range articles {
		response = append(response, ToArticleResponse(article))
	}
	return response
}

func (req *ArticleRequest) ToEntity() entities.Article {
	return entities.Article{
		Title:    req.Title,
		Content:  req.Content,
		ImageURL: req.ImageURL,
	}
}
