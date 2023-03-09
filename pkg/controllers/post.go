package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/soicchi/chatapp_backend/pkg/models"
	"github.com/soicchi/chatapp_backend/pkg/utils"
)

func (handler *Handler) CreatePost(ctx *gin.Context) {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	var postInput models.PostInput
	if err := ctx.ShouldBind(&postInput); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	userId := ctx.GetUint("userId")
	post := models.Post{
		Content: postInput.Content,
		UserID:  userId,
	}
	newPost, err := post.Create(handler.DB)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create post",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post":    newPost,
		"message": "Post created successfully",
	})
}

func (handler *Handler) GetAllPosts(ctx *gin.Context) {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	posts, err := models.FindAllPosts(handler.DB)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get posts",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts":   posts,
		"message": "Posts fetched successfully",
	})
}

func (handler *Handler) GetPost(ctx *gin.Context) {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
		})
		return
	}

	post, err := models.FindPostById(handler.DB, uint(postId))
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post":    post,
		"message": "Post fetched successfully",
	})
}

func (handler *Handler) UpdatePost(ctx *gin.Context) {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	var postInput models.PostInput
	if err := ctx.ShouldBind(&postInput); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
		})
		return
	}

	userId := ctx.GetUint("userId")
	post, err := models.FindPostById(handler.DB, uint(postId))
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post",
		})
		return
	}

	if post.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	updatedPost, err := post.Update(handler.DB, postInput)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update post",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post":    updatedPost,
		"message": "Post updated successfully",
	})
}

func (handler *Handler) DeletePost(ctx *gin.Context) {
	logger, err := utils.SetupLogger()
	if err != nil {
		panic(err)
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid post id",
		})
		return
	}

	targetPost, err := models.FindPostById(handler.DB, uint(postId))
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get post",
		})
		return
	}

	userId := ctx.GetUint("userId")
	if targetPost.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	if err := targetPost.Delete(handler.DB); err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete post",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
