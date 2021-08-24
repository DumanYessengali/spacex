package handler

import (
	"fmt"
	"garyshker"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) tokenCheck(c *gin.Context) (*garyshker.Auth, error) {
	token, err := h.services.Authorization.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return nil, err
	}

	foundAuth, err := h.services.Authorization.FetchAuth(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return nil, err
	}
	return foundAuth, nil
}

func (h *Handler) getAllPost(c *gin.Context) {
	videoPosts, articlePosts, err := h.services.Posts.GetAllPost()
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"video_post":   videoPosts,
		"article_post": articlePosts,
	})
}

func (h *Handler) getPostById(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	post, _, err := h.services.Posts.GetPostById(postId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) enrollPost(c *gin.Context) {
	foundAuth, err := h.tokenCheck(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	_, postType, err := h.services.Posts.GetPostById(postId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	fmt.Println(postType.Id)
	err = h.services.Posts.EnrollPost(postType, foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Post Saved")
}

func (h *Handler) getAllMySavedPost(c *gin.Context) {
	foundAuth, err := h.tokenCheck(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	videoPost, articlePost, err := h.services.Posts.GetAllMySavedPosts(foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"video_post":   videoPost,
		"article_post": articlePost,
	})
}

func (h *Handler) deleteMySavedPost(c *gin.Context) {
	foundAuth, err := h.tokenCheck(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	_, postConnection, err := h.services.Posts.GetPostById(postId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	err = h.services.Posts.DeleteMySavedPost(postConnection.Id, foundAuth.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, "You removed post")

}

func (h *Handler) createVideoPost(c *gin.Context) {
	_, err := h.tokenCheck(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var videoPost garyshker.VideoPost

	if err := c.ShouldBindJSON(&videoPost); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "invalid json "+err.Error())
		return
	}
	videoPostTable, err := h.services.Posts.CreateVideoPost(&videoPost)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, videoPostTable)
}

func (h *Handler) createArticlePost(c *gin.Context) {
	_, err := h.tokenCheck(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var articlePost garyshker.ArticlePost

	if err := c.ShouldBindJSON(&articlePost); err != nil {
		newErrorResponse(c, http.StatusUnprocessableEntity, "invalid json")
		return
	}
	articlePostTable, err := h.services.Posts.CreateArticlePost(&articlePost)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, articlePostTable)
}

type UpdateVideoPost struct {
	Title         *string                  `json:"title"`
	TitleType     *garyshker.PostTitleType `json:"title_type"`
	VideoDuration *int                     `json:"video_duration"`
	Description   *string                  `json:"description"`
	VideoUrl      *string                  `json:"video_url"`
}

type UpdateArticlePost struct {
	Title                      *string                  `json:"title"`
	TitleType                  *garyshker.PostTitleType `json:"title_type"`
	Duration                   *int                     `json:"duration"`
	AuthorInformationParagraph *string                  `json:"author_information_paragraph"`
	ParagraphName              *string                  `json:"paragraph_name"`
	Description                *string                  `json:"description"`
	AuthorName                 *string                  `json:"author_name"`
	AuthorPosition             *string                  `json:"author_position"`
}

func (h *Handler) updatePost(c *gin.Context) {
	_, err := h.tokenCheck(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	_, postConnection, err := h.services.Posts.GetPostById(postId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if postConnection.PostType == garyshker.VideoPosts {
		var v UpdateVideoPost
		if err := c.ShouldBindJSON(&v); err != nil {
			newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		videoPost, err := h.services.Posts.GetVideoPostById(postConnection.PostId)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		updateVideoPost, err := h.services.Posts.UpdateVideoPost(v.Title, v.Description, v.VideoUrl, v.TitleType, v.VideoDuration, videoPost)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, updateVideoPost)
	} else {
		var a UpdateArticlePost
		if err := c.ShouldBindJSON(&a); err != nil {
			newErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		articlePost, err := h.services.Posts.GetArticlePostById(postConnection.PostId)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		updateArticlePost, err := h.services.Posts.UpdateArticlePost(a.Title, a.AuthorInformationParagraph, a.ParagraphName, a.Description, a.AuthorName, a.AuthorPosition, a.Duration, a.TitleType, articlePost)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		c.JSON(http.StatusOK, updateArticlePost)
	}
}
