package service

import (
	"errors"
	"garyshker"
	"garyshker/pkg/repository"
)

type PostService struct {
	repos repository.Posts
}

func NewPostService(repos repository.Posts) *PostService {
	return &PostService{repos: repos}
}

func (p *PostService) GetAllPost() (*[]garyshker.VideoPost, *[]garyshker.ArticlePost, error) {
	return p.repos.GetAllPost()
}

func (p *PostService) GetPostById(postId int) (interface{}, *garyshker.PostConnection, error) {
	return p.repos.GetPostById(postId)
}

func (p *PostService) GetVideoPostById(id uint64) (*garyshker.VideoPost, error) {
	return p.repos.GetVideoPostById(id)
}

func (p *PostService) GetArticlePostById(id uint64) (*garyshker.ArticlePost, error) {
	return p.repos.GetArticlePostById(id)
}

func (p *PostService) EnrollPost(post *garyshker.PostConnection, userId uint64) error {
	return p.repos.EnrollPost(post, userId)
}

func (p *PostService) GetAllMySavedPosts(userId uint64) (*[]garyshker.VideoPost, *[]garyshker.ArticlePost, error) {
	return p.repos.GetAllMySavedPosts(userId)
}

func (p *PostService) CreateVideoPost(videoPost *garyshker.VideoPost) (*garyshker.VideoPost, error) {
	return p.repos.CreateVideoPost(videoPost)
}

func (p *PostService) CreateArticlePost(articlePost *garyshker.ArticlePost) (*garyshker.ArticlePost, error) {
	return p.repos.CreateArticlePost(articlePost)
}

func (p *PostService) DeleteMySavedPost(postId, userId uint64) error {
	return p.repos.DeleteMySavedPost(postId, userId)
}

func (p *PostService) UpdateVideoPost(title, description, videoUrl *string, titleType *garyshker.PostTitleType, videoDuration *int, videoPost *garyshker.VideoPost) (*garyshker.VideoPost, error) {
	didUpdate := false

	if title != nil {
		videoPost.Title = *title
		didUpdate = true
	}

	if description != nil {
		videoPost.Description = *description
		didUpdate = true
	}

	if videoUrl != nil {
		videoPost.VideoUrl = *videoUrl
		didUpdate = true
	}

	if titleType != nil {
		videoPost.TitleType = *titleType
		didUpdate = true
	}

	if videoDuration != nil {
		videoPost.VideoDuration = *videoDuration
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	return p.repos.UpdateVideoPost(videoPost)
}

func (p *PostService) UpdateArticlePost(title, authorInformationParagraph, paragraphName, description, authorName, authorPosition *string, duration *int, titleType *garyshker.PostTitleType, articlePost *garyshker.ArticlePost) (*garyshker.ArticlePost, error) {
	didUpdate := false

	if title != nil {
		articlePost.Title = *title
		didUpdate = true
	}

	if description != nil {
		articlePost.Description = *description
		didUpdate = true
	}

	if titleType != nil {
		articlePost.TitleType = *titleType
		didUpdate = true
	}

	if duration != nil {
		articlePost.Duration = *duration
		didUpdate = true
	}

	if authorInformationParagraph != nil {
		articlePost.AuthorInformationParagraph = *authorInformationParagraph
		didUpdate = true
	}

	if paragraphName != nil {
		articlePost.ParagraphName = *paragraphName
		didUpdate = true
	}

	if authorName != nil {
		articlePost.AuthorName = *authorName
		didUpdate = true
	}

	if authorPosition != nil {
		articlePost.AuthorPosition = *authorPosition
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	return p.repos.UpdateArticlePost(articlePost)
}
