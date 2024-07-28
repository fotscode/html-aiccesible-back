package repositories

import (
	"errors"
	m "html-aiccesible/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(user *m.User, postBody *m.CreatePostBody) (*m.Post, error)
	UpdatePost(user *m.User, postBody *m.UpdatePostBody) (*m.Post, error)
	GetPost(postID int) (*m.Post, error)
	ListPosts(page, size int) ([]m.Post, error)
	DeletePost(user *m.User, postID int) error
	LikePost(user *m.User, postID int) error
	GetPostLikes(postID int) (int, error)
}

type postRepository struct {
	DB *gorm.DB
}

func PostRepo(db *gorm.DB) PostRepository {
	return &postRepository{
		DB: db,
	}
}

func (r *postRepository) CreatePost(user *m.User, postBody *m.CreatePostBody) (*m.Post, error) {
	post := &m.Post{
		Title:       postBody.Title,
		Description: postBody.Description,
		Before:      postBody.Before,
		After:       postBody.After,
		UserID:      user.ID,
	}
	res := r.DB.Create(post)
	if res.Error != nil {
		return nil, res.Error
	}
	return post, nil
}

func (r *postRepository) UpdatePost(user *m.User, postBody *m.UpdatePostBody) (*m.Post, error) {
	post, err := r.GetPost(int(postBody.ID))
	if err != nil {
		return nil, err
	}
	post.Title = postBody.Title
	post.Description = postBody.Description
	post.Before = postBody.Before
	post.After = postBody.After
	res := r.DB.Save(post)
	if res.Error != nil {
		return nil, res.Error
	}
	return post, nil
}

func (r *postRepository) GetPost(postID int) (*m.Post, error) {
	var post m.Post
	res := r.DB.Preload("Likes").Preload("Comments").First(&post, postID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &post, nil
}

func (r *postRepository) ListPosts(page, size int) ([]m.Post, error) {
	var posts []m.Post
	res := r.DB.Preload("Likes").Preload("Comments").Limit(size).Offset((page - 1) * size).Find(&posts)
	if res.Error != nil {
		return nil, res.Error
	}
	return posts, nil
}

func (r *postRepository) DeletePost(user *m.User, postID int) error {
	post, err := r.GetPost(postID)
	if err != nil {
		return err
	}
	if post.UserID != user.ID {
		return errors.New("user is not the owner of the post")
	}
	res := r.DB.Delete(&m.Post{}, postID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *postRepository) LikePost(user *m.User, postID int) error {
	post, err := r.GetPost(postID)
	if err != nil {
		return err
	}
	// if the user already liked the post, remove the like
	for _, like := range post.Likes {
		if like.ID == user.ID {
			err = r.DB.Model(post).Association("Likes").Delete(user)
			if err != nil {
				return err
			}
			return nil
		}
	}

	// if user did not like the post, add the like
	err = r.DB.Model(post).Association("Likes").Append(user)
	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) GetPostLikes(postID int) (int, error) {
	post, err := r.GetPost(postID)
	if err != nil {
		return 0, err
	}
	return len(post.Likes), nil
}
