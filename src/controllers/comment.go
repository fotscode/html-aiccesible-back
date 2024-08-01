package controllers

import (
	"html-aiccesible/httputil"
	m "html-aiccesible/models"

	"github.com/gin-gonic/gin"
)

func (b *Controller) CreateComment(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*m.CreateCommentBody)
	user := c.MustGet("user").(*m.User)
	comment, err := b.CommentRepo.CreateComment(user, body)
	if err != nil {
		httputil.InternalServerError(c, err)
		return
	}
	httputil.Created(c, comment)
}

func (b *Controller) UpdateComment(c *gin.Context) {
	body := c.MustGet(gin.BindKey).(*m.UpdateCommentBody)
	user := c.MustGet("user").(*m.User)
	comment, err := b.CommentRepo.UpdateComment(user, body)
	if err != nil {
		httputil.InternalServerError(c, err)
		return
	}
	httputil.OK(c, comment)
}

func (b *Controller) ListComments(c *gin.Context) {
	lo := c.MustGet("lo").(*m.ListOptions)
	getOpt := c.MustGet("getOpt").(*m.GetOptions)
	_, err := b.PostRepo.GetPost(getOpt.Id)
	if err != nil {
		httputil.NotFound(c, err)
		return
	}
	comments, err := b.CommentRepo.ListComments(lo.Page, lo.Size, getOpt.Id)
	if err != nil {
		httputil.InternalServerError(c, err)
		return
	}
	httputil.OK(c, comments)
}

func (b *Controller) GetComment(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*m.GetOptions)
	comment, err := b.CommentRepo.GetComment(getOpt.Id)
	if err != nil {
		httputil.NotFound(c, err)
		return
	}
	httputil.OK(c, comment)
}

func (b *Controller) DeleteComment(c *gin.Context) {
	getOpt := c.MustGet("getOpt").(*m.GetOptions)
	user := c.MustGet("user").(*m.User)
	err := b.CommentRepo.DeleteComment(user, getOpt.Id)
	if err != nil {
		httputil.InternalServerError(c, err)
		return
	}
	httputil.OK(c, "Deleted comment successfully")
}
