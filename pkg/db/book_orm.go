package db

import (
	"gorm.io/gorm"
	"yujian-backend/pkg/model"
)

var (
	postRepository    PostRepository
	commentRepository CommentRepository
)

// PostRepository 帖子数据库操作仓库
type PostRepository struct {
	DB *gorm.DB
}

// CommentRepository 评论数据库操作仓库
type CommentRepository struct {
	DB *gorm.DB
}

// GetPostRepository 获取帖子仓库单例
func GetPostRepository() *PostRepository {
	return &postRepository
}

// GetCommentRepository 获取评论仓库单例
func GetCommentRepository() *CommentRepository {
	return &commentRepository
}

// ------------------------- 帖子相关操作 -------------------------

// CreatePost 创建帖子
func (r *PostRepository) CreatePost(post *model.PostDTO) error {
	postDO := post.TransformToDO()
	postDO.ViewCount = 0 // 初始化阅读数
	if err := r.DB.Create(postDO).Error; err != nil {
		return err
	}
	post.Id = postDO.Id // 回填生成的ID
	return nil
}

// GetPostById 根据ID获取帖子详情
func (r *PostRepository) GetPostById(id int64) (*model.PostDTO, error) {
	var postDO model.PostDO
	if err := r.DB.First(&postDO, id).Error; err != nil {
		return nil, err
	}

	// 获取作者信息
	userDTO, err := GetUserRepository().GetUserById(postDO.AuthorId)
	if err != nil {
		return nil, err
	}

	// 获取关联评论
	comments, err := r.GetPostCommentsByPostId(postDO.Id)
	if err != nil {
		return nil, err
	}

	// 转换为DTO
	return postDO.TransformToDTO(userDTO, comments), nil
}

// ListPosts 分页获取帖子列表（支持分类过滤）
func (r *PostRepository) ListPosts(category string, offset, limit int) ([]*model.PostDTO, error) {
	var postsDO []model.PostDO
	query := r.DB
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if err := query.Offset(offset).Limit(limit).Find(&postsDO).Error; err != nil {
		return nil, err
	}

	// 转换为DTO
	postsDTO := make([]*model.PostDTO, len(postsDO))
	for i, postDO := range postsDO {
		userDTO, err := GetUserRepository().GetUserById(postDO.AuthorId)
		if err != nil {
			return nil, err
		}

		comments, err := r.GetPostCommentsByPostId(postDO.Id)
		if err != nil {
			return nil, err
		}

		postsDTO[i] = postDO.TransformToDTO(userDTO, comments)
	}
	return postsDTO, nil
}

// IncrementViewCount 增加帖子阅读数（原子操作）
func (r *PostRepository) IncrementViewCount(postId int64) error {
	return r.DB.Model(&model.PostDO{}).
		Where("id = ?", postId).
		Update("view_count", gorm.Expr("view_count + 1")).Error
}

// HandleLike 处理帖子点赞/踩（原子操作）
func (r *PostRepository) HandleLike(postId int64, isLike bool) error {
	field := "like_count"
	if !isLike {
		field = "dislike_count"
	}
	return r.DB.Model(&model.PostDO{}).
		Where("id = ?", postId).
		Update(field, gorm.Expr(field+" + 1")).Error
}

// DeletePost 删除帖子
func (r *PostRepository) DeletePost(postId int64) error {
	return r.DB.Delete(&model.PostDO{}, postId).Error
}

// ------------------------- 评论相关操作 -------------------------

// Create 创建评论
func (r *CommentRepository) Create(comment *model.PostCommentDTO) error {
	commentDO := comment.TransformToDO()
	if err := r.DB.Create(commentDO).Error; err != nil {
		return err
	}
	comment.Id = commentDO.Id // 回填生成的ID
	return nil
}

// GetPostCommentsByPostId 获取帖子下的所有评论
func (r *PostRepository) GetPostCommentsByPostId(postId int64) ([]*model.PostCommentDTO, error) {
	var commentsDO []model.PostCommentDO
	if err := r.DB.Where("post_id = ?", postId).Find(&commentsDO).Error; err != nil {
		return nil, err
	}

	// 转换为DTO
	commentsDTO := make([]*model.PostCommentDTO, len(commentsDO))
	for i, commentDO := range commentsDO {
		userDTO, err := GetUserRepository().GetUserById(commentDO.AuthorId)
		if err != nil {
			return nil, err
		}
		commentsDTO[i] = commentDO.TransformToDTO(userDTO)
	}
	return commentsDTO, nil
}

// BatchGetPostCommentById 批量获取评论（用于优化N+1查询）
func (r *PostRepository) BatchGetPostCommentById(ids []int64) ([]*model.PostCommentDTO, error) {
	var commentsDO []model.PostCommentDO
	if err := r.DB.Where("id IN (?)", ids).Find(&commentsDO).Error; err != nil {
		return nil, err
	}

	// 转换为DTO
	commentsDTO := make([]*model.PostCommentDTO, len(commentsDO))
	for i, commentDO := range commentsDO {
		userDTO, err := GetUserRepository().GetUserById(commentDO.AuthorId)
		if err != nil {
			return nil, err
		}
		commentsDTO[i] = commentDO.TransformToDTO(userDTO)
	}
	return commentsDTO, nil
}

// IncrementCommentLikeCount 增加评论点赞数（原子操作）
func (r *CommentRepository) IncrementCommentLikeCount(commentId int64) error {
	return r.DB.Model(&model.PostCommentDO{}).
		Where("id = ?", commentId).
		Update("like_count", gorm.Expr("like_count + 1")).Error
}
