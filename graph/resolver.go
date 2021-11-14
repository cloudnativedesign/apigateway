package graph

import "github.com/cloudnativedesign/apigateway/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArticleStore map[string]model.Article
	BlogStore    map[string]model.Blog
}
