package repositories

import (
	"errors"
	m "html-aiccesible/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	ListComments(page, size, postID int) ([]m.Comment, error)
	GetComment(commentID int) (*m.Comment, error)
	CreateComment(user *m.User, commentBody *m.CreateCommentBody) (*m.Comment, error)
	UpdateComment(user *m.User, commentBody *m.UpdateCommentBody) (*m.Comment, error)
	DeleteComment(user *m.User, commentID int) error
}

type commentRepository struct {
	DB *gorm.DB
}

func CommentRepo(db *gorm.DB) CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (r *commentRepository) ListComments(page, size, postID int) ([]m.Comment, error) {
	var comments []m.Comment
	res := r.DB.Where("post_id = ?", postID).Limit(size).Offset((page - 1) * size).Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}

func (r *commentRepository) GetComment(commentID int) (*m.Comment, error) {
	var comment m.Comment
	res := r.DB.First(&comment, commentID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &comment, nil
}

func (r *commentRepository) CreateComment(user *m.User, commentBody *m.CreateCommentBody) (*m.Comment, error) {
	comment := &m.Comment{
		Title:   commentBody.Title,
		Content: commentBody.Content,
		PostID:  commentBody.PostID,
		UserID:  user.ID,
	}
	res := r.DB.Create(comment)
	if res.Error != nil {
		return nil, res.Error
	}
	return comment, nil
}

func (r *commentRepository) UpdateComment(user *m.User, commentBody *m.UpdateCommentBody) (*m.Comment, error) {
	comment, err := r.GetComment(int(commentBody.ID))
	if err != nil {
		return nil, err
	}
	if comment.UserID != user.ID {
		return nil, errors.New("user is not the owner of the comment")
	}
	comment.Title = commentBody.Title
	comment.Content = commentBody.Content
	res := r.DB.Save(comment)
	if res.Error != nil {
		return nil, res.Error
	}
	return comment, nil
}

func (r *commentRepository) DeleteComment(user *m.User, commentID int) error {
	comment, err := r.GetComment(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != user.ID {
		return errors.New("user is not the owner of the comment")
	}
	res := r.DB.Delete(&m.Comment{}, commentID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
