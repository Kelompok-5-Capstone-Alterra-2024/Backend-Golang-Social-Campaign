package dto

import (
	"capstone/entities"
)

type ArticleResponses struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ImageURL  string `json:"image_url"`
}

func ToArticleResponses(article entities.Article) ArticleResponses {

	return ArticleResponses{
		ID:        article.ID,
		CreatedAt: article.CreatedAt.Format("2006-01-02"),
		Title:     article.Title,
		Content:   article.Content,
		ImageURL:  article.ImageURL,
	}
}

func ToArticleResponsesList(articles []entities.Article) []ArticleResponses {
	var res []ArticleResponses
	for _, article := range articles {
		res = append(res, ToArticleResponses(article))
	}
	return res
}

type ArticleRequest struct {
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	ImageURL string `json:"image_url" form:"image_url"`
}

func (req *ArticleRequest) ToEntity(imageUrl string) entities.Article {
	return entities.Article{
		Title:    req.Title,
		Content:  req.Content,
		ImageURL: imageUrl,
	}
}
