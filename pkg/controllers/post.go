package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/models"
)

func (handler *Handler) CreatePost(ctx *gin.Context) {
	var postInput models.CreatePostInput
	if err := ctx.ShouldBind(&postInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	userId := ctx.GetUint("userId")
	post := models.Post{
		Content: postInput.Content,
		UserID: userId,
	}
	newPost, err := post.Create(handler.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to create post",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post": newPost,
		"message": "Post created successfully",
	})
}

func (handler *Handler) GetAllPosts(ctx *gin.Context) {
	posts, err := models.FindAllPosts(handler.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to get posts",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts": posts,
		"message": "Posts fetched successfully",
	})
}

func (handler *Handler) GetPost(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid post id",
		})
		return
	}

	post, err := models.FindPostById(handler.DB, uint(postId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to get post",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post": post,
		"message": "Post fetched successfully",
	})
}