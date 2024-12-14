package db

import (
	"gorm.io/gorm"
	"yujian-backend/pkg/model"
)

var postRepository PostRepository

type PostRepository struct {
	DB *gorm.DB
}

func GetPostRepository() *PostRepository {
	return &postRepository
}

// CreatePost 创建帖子
func (r *PostRepository) CreatePost(postDTO *model.PostDTO) (int64, error) {
	postDO := postDTO.TransformToDO()
	if err := r.DB.Create(postDO).Error; err != nil {
		return 0, err
	}
	return postDO.Id, nil
}

// GetPostById 根据ID获取帖子
func (r *PostRepository) GetPostById(id int64) (*model.PostDTO, error) {
	var post model.PostDO
	if err := r.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	userDTO, err := userRepository.GetUserById(post.AuthorId)
	if err != nil {
		return nil, err
	}

	comments, err := r.GetPostCommentsByPostId(post.Id)
	if err != nil {
		return nil, err
	}

	return post.TransformToDTO(userDTO, comments), nil
}

// UpdatePost 更新帖子
func (r *PostRepository) UpdatePost(postDTO *model.PostDTO) error {
	postDO := postDTO.TransformToDO()
	return r.DB.Save(postDO).Error
}

// DeletePost 删除帖子
func (r *PostRepository) DeletePost(id int64) error {
	return r.DB.Delete(&model.PostDO{}, id).Error
}

// ListPosts 获取帖子列表
func (r *PostRepository) ListPosts(offset, limit int) ([]*model.PostDTO, error) {
	var posts []model.PostDO
	if err := r.DB.Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, err
	}

	postDTOs := make([]*model.PostDTO, len(posts))
	for i, post := range posts {
		userDTO, err := userRepository.GetUserById(post.AuthorId)
		if err != nil {
			return nil, err
		}

		comments, err := r.GetPostCommentsByPostId(post.Id)
		if err != nil {
			return nil, err
		}

		postDTOs[i] = post.TransformToDTO(userDTO, comments)
	}
	return postDTOs, nil
}

// GetPostCommentsByPostId 根据帖子id获取帖子评论
func (r *PostRepository) GetPostCommentsByPostId(postId int64) ([]*model.PostCommentDTO, error) {
	var comments []model.PostCommentDO
	if err := r.DB.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		return nil, err
	}

	// 将PostCommentDO转换为PostCommentDTO
	postCommentDTOs := make([]*model.PostCommentDTO, len(comments))
	for i, comment := range comments {
		postCommentDTOs[i] = comment.TransformToDTO()
	}
	return postCommentDTOs, nil
}

// BatchGetPostCommentById 批量获取帖子评论
func (r *PostRepository) BatchGetPostCommentById(ids []int64) ([]*model.PostCommentDTO, error) {
	var comments []model.PostCommentDO
	if err := r.DB.Where("id IN (?)", ids).Find(&comments).Error; err != nil {
		return nil, err
	}

	// 将PostCommentDO转换为PostCommentDTO
	postCommentDTOs := make([]*model.PostCommentDTO, len(comments))
	for i, comment := range comments {
		postCommentDTOs[i] = comment.TransformToDTO()
	}
	return postCommentDTOs, nil
}
