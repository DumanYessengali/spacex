package repository

import (
	"garyshker"
	"github.com/jinzhu/gorm"
	"time"
)

type PostPostgres struct {
	db *gorm.DB
}

func NewPostPostgres(db *gorm.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (p *PostPostgres) GetAllPost() (*[]garyshker.VideoPost, *[]garyshker.ArticlePost, error) {
	videoPosts := []garyshker.VideoPost{}
	articlePosts := []garyshker.ArticlePost{}

	rows, err := p.db.Debug().Table("post_connections").Select("*").Rows()

	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		vPosts := &garyshker.VideoPost{}
		aPosts := &garyshker.ArticlePost{}
		post := &garyshker.PostConnection{}
		err = rows.Scan(&post.Id, &post.PostId, &post.PostType)
		if post.PostType == garyshker.VideoPosts {
			err = p.db.Debug().Where("id = ?", post.PostId).Take(&vPosts).Error
			if err != nil {
				return nil, nil, err
			}
			videoPosts = append(videoPosts, *vPosts)
		} else {
			err = p.db.Debug().Where("id = $1", post.PostId).Take(&aPosts).Error
			if err != nil {
				return nil, nil, err
			}
			articlePosts = append(articlePosts, *aPosts)
		}
	}
	return &videoPosts, &articlePosts, nil
}

func (p *PostPostgres) GetPostById(postId int) (interface{}, *garyshker.PostConnection, error) {
	videoPosts := &garyshker.VideoPost{}
	articlePosts := &garyshker.ArticlePost{}
	post := &garyshker.PostConnection{}

	err := p.db.Debug().Table("post_connections").Select("*").Where("id = $1", postId).Scan(&post).Error
	if err != nil {
		return nil, nil, err
	}

	if post.PostType == garyshker.VideoPosts {
		err = p.db.Debug().Where("id = ?", post.PostId).Take(&videoPosts).Error

		if err != nil {
			return nil, nil, err
		}

		return videoPosts, post, nil
	} else {
		err = p.db.Debug().Where("id = ?", post.PostId).Take(&articlePosts).Error

		if err != nil {
			return nil, nil, err
		}

		return articlePosts, post, nil
	}
}

func (p *PostPostgres) GetVideoPostById(id uint64) (*garyshker.VideoPost, error) {
	videoPost := &garyshker.VideoPost{}
	err := p.db.Debug().Where("id = $1", id).Take(&videoPost).Error
	if err != nil {
		return nil, err
	}
	return videoPost, nil
}

func (p *PostPostgres) GetArticlePostById(id uint64) (*garyshker.ArticlePost, error) {
	articlePost := &garyshker.ArticlePost{}
	err := p.db.Debug().Where("id = $1", id).Take(&articlePost).Error
	if err != nil {
		return nil, err
	}
	return articlePost, nil
}

func (p *PostPostgres) EnrollPost(post *garyshker.PostConnection, userId uint64) error {
	userPost := &garyshker.UserSavedPost{}
	userPost.PostConnectionId = uint64(post.Id)
	userPost.UserId = userId

	err := p.db.Debug().Create(&userPost).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostPostgres) GetAllMySavedPosts(userId uint64) (*[]garyshker.VideoPost, *[]garyshker.ArticlePost, error) {
	videoPosts := []garyshker.VideoPost{}
	articlePosts := []garyshker.ArticlePost{}

	rows, err := p.db.Debug().Raw("select * from post_connections where id in (select post_connection_id from user_saved_posts where user_id = $1)", userId).Rows()
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		vPosts := &garyshker.VideoPost{}
		aPosts := &garyshker.ArticlePost{}
		post := &garyshker.PostConnection{}
		err = rows.Scan(&post.Id, &post.PostId, &post.PostType)
		if post.PostType == garyshker.VideoPosts {
			err = p.db.Debug().Where("id = $1", post.PostId).Take(&vPosts).Error
			if err != nil {
				return nil, nil, err
			}
			videoPosts = append(videoPosts, *vPosts)
		} else {
			err = p.db.Debug().Where("id = $1", post.PostId).Take(&aPosts).Error
			if err != nil {
				return nil, nil, err
			}
			articlePosts = append(articlePosts, *aPosts)
		}
	}

	return &videoPosts, &articlePosts, nil
}

func (p *PostPostgres) CreateVideoPost(videoPost *garyshker.VideoPost) (*garyshker.VideoPost, error) {
	videoPost.Updated = time.Now()
	videoPost.Created = time.Now()
	err := p.db.Debug().Create(&videoPost).Error
	if err != nil {
		return nil, err
	}
	postConnection := &garyshker.PostConnection{}
	postConnection.PostId = videoPost.Id
	postConnection.PostType = garyshker.VideoPosts
	err = p.db.Debug().Create(&postConnection).Error
	if err != nil {
		return nil, err
	}
	return videoPost, nil
}

func (p *PostPostgres) CreateArticlePost(articlePost *garyshker.ArticlePost) (*garyshker.ArticlePost, error) {
	articlePost.Updated = time.Now()
	articlePost.Created = time.Now()
	err := p.db.Debug().Create(&articlePost).Error
	if err != nil {
		return nil, err
	}
	postConnection := &garyshker.PostConnection{}
	postConnection.PostId = articlePost.Id
	postConnection.PostType = garyshker.ArticlePosts
	err = p.db.Debug().Create(&postConnection).Error
	if err != nil {
		return nil, err
	}
	return articlePost, nil
}

func (p *PostPostgres) DeleteMySavedPost(postId, userId uint64) error {
	userSavedPost := &garyshker.UserSavedPost{}
	err := p.db.Debug().Where("post_connection_id =$1 and user_id = $2", postId, userId).Delete(&userSavedPost).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostPostgres) UpdateVideoPost(videoPost *garyshker.VideoPost) (*garyshker.VideoPost, error) {
	err := p.db.Debug().Model(&videoPost).Updates(garyshker.VideoPost{
		Title:         videoPost.Title,
		TitleType:     videoPost.TitleType,
		VideoDuration: videoPost.VideoDuration,
		Description:   videoPost.Description,
		VideoUrl:      videoPost.VideoUrl,
		Updated:       time.Now(),
	}).Error
	if err != nil {
		return nil, err
	}

	return videoPost, nil
}

func (p *PostPostgres) UpdateArticlePost(articlePost *garyshker.ArticlePost) (*garyshker.ArticlePost, error) {
	err := p.db.Debug().Model(&articlePost).Updates(garyshker.ArticlePost{
		Title:                      articlePost.Title,
		TitleType:                  articlePost.TitleType,
		Duration:                   articlePost.Duration,
		AuthorInformationParagraph: articlePost.AuthorInformationParagraph,
		ParagraphName:              articlePost.ParagraphName,
		Description:                articlePost.Description,
		AuthorName:                 articlePost.AuthorName,
		AuthorPosition:             articlePost.AuthorPosition,
		Updated:                    time.Now(),
	}).Error
	if err != nil {
		return nil, err
	}

	return articlePost, nil
}
