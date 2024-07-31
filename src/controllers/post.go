package controllers

import (
	"html-aiccesible/httputil"
	m "html-aiccesible/models"

	"github.com/gin-gonic/gin"
)

func (b *Controller) CreatePost(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*m.CreatePostBody)
	user := c.MustGet("user").(*m.User)
	post, err := b.PostRepo.CreatePost(user, body)
	if err != nil {
		httputil.InternalServerError(c, err.Error())
		return
	}
	httputil.Created(c, post)
}

func (b *Controller) UpdatePost(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*m.UpdatePostBody)
	user := c.MustGet("user").(*m.User)
	post, err := b.PostRepo.UpdatePost(user, body)
	if err != nil {
		httputil.InternalServerError(c, err.Error())
		return
	}
	httputil.OK(c, post)
}

func (b *Controller) GetPost(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*m.GetOptions)
	post, err := b.PostRepo.GetPost(getOpt.Id)
	if err != nil {
		httputil.NotFound(c, err.Error())
		return
	}
	httputil.OK(c, post)
}

func (b *Controller) DeletePost(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*m.GetOptions)
	user := c.MustGet("user").(*m.User)
	err := b.PostRepo.DeletePost(user, getOpt.Id)
	if err != nil {
		httputil.NotFound(c, err.Error())
		return
	}
	httputil.OK(c, "Deleted post successfully")
}

func (b *Controller) ListPosts(c *gin.Context) {
	lo := c.MustGet("lo").(*m.ListOptions)
	posts, err := b.PostRepo.ListPosts(lo.Page, lo.Size)
	if err != nil {
		httputil.InternalServerError(c, err.Error())
		return
	}
	httputil.OK(c, posts)
}

func (b *Controller) LikePost(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*m.GetOptions)
	user := c.MustGet("user").(*m.User)
	err := b.PostRepo.LikePost(user, getOpt.Id)
	if err != nil {
		httputil.InternalServerError(c, err.Error())
		return
	}
	httputil.OK(c, "Liked post successfully")
}

func (b *Controller) GetPostLikes(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*m.GetOptions)
	likes, err := b.PostRepo.GetPostLikes(getOpt.Id)
	if err != nil {
		httputil.InternalServerError(c, err.Error())
		return
	}
	httputil.OK(c, likes)
}
