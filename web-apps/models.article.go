// models.article.go

package main

import (
	"errors"

	"github.com/lithammer/shortuuid"
)

type article struct {
	ID      string `json:"id" form:id`
	Title   string `json:"title" form:title`
	Content string `json:"content" form:content`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: shortuuid.New(), Title: "For Shay", Content: "I love you!!!"},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id string) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func addArticle(title string, content string) {
	articleList = append(articleList, article{
		ID:      shortuuid.New(),
		Title:   title,
		Content: content,
	})
}

func removeArticle(id string) {
	newArticleList := []article{}
	for _, a := range articleList {
		if a.ID != id {
			newArticleList = append(newArticleList, a)
		}
	}
	articleList = newArticleList
}
