package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	n := len(r.Resolver.PostStore)
	if n == 0 {
		r.Resolver.PostStore = make(map[string]model.Post)
	}

	id := input.ID
	var tagString string
	var post model.Post
	post.Title = &input.Title

	// Appends all tags into concatenable strings
	for _, tag := range input.Tags {
		tagString = tagString + " #" + *tag
	}

	// Define the post object

	post.Content = input.Content + " " + tagString

	if id != nil {
		_, ok := r.PostStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		r.Resolver.PostStore[*id] = post
	} else {
		nid := strconv.Itoa(n + 1)
		post.ID = nid
		r.Resolver.PostStore[nid] = post
	}

	return &post, nil
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input *model.NewAccount) (*model.SocialMediaAccount, error) {
	n := len(r.Resolver.AccountStore)
	if n == 0 {
		r.Resolver.AccountStore = make(map[string]model.SocialMediaAccount)
	}

	id := input.ID
	provider := model.SocialMediaProvider(strings.ToUpper(input.Provider))
	if !model.SocialMediaProvider.IsValid(provider) {
		return nil, fmt.Errorf("social media provider not supported. Please check correct typing")
	}
	var account model.SocialMediaAccount
	account.Provider = provider
	account.ProviderID = input.ProviderID
	account.ProviderToken = input.ProviderToken
	account.ProviderUsername = input.ProviderUsername

	if id != nil {
		_, ok := r.Resolver.AccountStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		r.Resolver.AccountStore[*id] = account

	} else {
		nid := strconv.Itoa(n + 1)
		account.ID = nid
		r.Resolver.AccountStore[nid] = account
	}

	return &account, nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	article, ok := r.Resolver.ArticleStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
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

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	post, ok := r.Resolver.PostStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &post, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	for idx := range r.Resolver.PostStore {
		post := r.Resolver.PostStore[idx]
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *queryResolver) Account(ctx context.Context, id string) (*model.SocialMediaAccount, error) {
	account, ok := r.Resolver.AccountStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &account, nil
}

func (r *queryResolver) Accounts(ctx context.Context) ([]*model.SocialMediaAccount, error) {
	accounts := make([]*model.SocialMediaAccount, 0)
	for idx := range r.Resolver.AccountStore {
		account := r.Resolver.AccountStore[idx]
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (r *queryResolver) AccountsByProvider(ctx context.Context, provider *string) ([]*model.SocialMediaAccount, error) {
	accounts := make([]*model.SocialMediaAccount, 0)
	setProvider := model.SocialMediaProvider(strings.ToUpper(*provider))

	if !setProvider.IsValid() {
		return nil, fmt.Errorf("provider not supported. Please check typging")
	}
	for idx := range r.Resolver.AccountStore {
		account := r.Resolver.AccountStore[idx]
		if account.Provider == setProvider {
			accounts = append(accounts, &account)
		}
	}
	return accounts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
