package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudnativedesign/apigateway/graph/generated"
	"github.com/cloudnativedesign/apigateway/graph/model"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, input model.NewArticle) (*model.Article, error) {
	id := input.ID
	var article model.Article
	article.Title = input.Title
	article.Subtitle = input.Subtitle
	article.Content = input.Content

	n := len(r.Resolver.ArticleStore)
	if n == 0 {
		r.Resolver.ArticleStore = make(map[string]model.Article)
	}

	if id != nil {
		_, ok := r.Resolver.ArticleStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		r.Resolver.ArticleStore[*id] = article
	} else {
		// generate unique id
		nid := strconv.Itoa(n + 1)
		article.ID = nid
		r.Resolver.ArticleStore[nid] = article
	}

	return &article, nil
}

func (r *mutationResolver) CreateBlog(ctx context.Context, input model.NewBlog) (*model.Blog, error) {
	id := input.ID
	var blog model.Blog
	blog.Title = input.Title
	blog.Subtitle = input.Subtitle
	blog.Content = input.Content

	n := len(r.Resolver.BlogStore)
	if n == 0 {
		r.Resolver.BlogStore = make(map[string]model.Blog)
	}

	if id != nil {
		_, ok := r.Resolver.BlogStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		r.Resolver.BlogStore[*id] = blog
	} else {
		nid := strconv.Itoa(n + 1)
		blog.ID = nid
		r.Resolver.BlogStore[nid] = blog
	}

	return &blog, nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	article, ok := r.Resolver.ArticleStore[id]
	if !ok {
		return nil, fmt.Errorf("Not found")
	}
	return &article, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	articles := make([]*model.Article, 0)
	for idx := range r.Resolver.ArticleStore {
		article := r.Resolver.ArticleStore[idx]
		articles = append(articles, &article)
	}
	return articles, nil
}

func (r *queryResolver) Blog(ctx context.Context, id string) (*model.Blog, error) {
	blog, ok := r.Resolver.BlogStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &blog, nil
}

func (r *queryResolver) Blogs(ctx context.Context) ([]*model.Blog, error) {
	blogs := make([]*model.Blog, 0)
	for idx := range r.BlogStore {
		blog := r.Resolver.BlogStore[idx]
		blogs = append(blogs, &blog)
	}
	return blogs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
