package model

import (
	"time"
)

// PostDTO 帖子DTO（数据传输对象）
type PostDTO struct {
	Id           int64             `json:"id"`            // 帖子ID
	Author       *UserDTO          `json:"author"`        // 发布者信息
	Title        string            `json:"title"`         // 帖子标题
	Content      string            `json:"content"`       // 帖子内容
	Category     string            `json:"category"`      // 帖子分类
	LikeCount    int               `json:"like_count"`    // 点赞数
	DislikeCount int               `json:"dislike_count"` // 踩数
	ViewCount    int               `json:"view_count"`    // 阅读数
	EditTime     time.Time         `json:"edit_time"`     // 编辑时间
	Comments     []*PostCommentDTO `json:"comments"`      // 评论列表
}

// PostDO 帖子DO（数据库对象）
type PostDO struct {
	Id           int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 帖子ID
	AuthorId     int64     `gorm:"column:author_id" json:"author_id"`            // 发布者ID
	Title        string    `gorm:"column:title" json:"title"`                    // 帖子标题
	Content      string    `gorm:"column:content" json:"content"`                // 帖子内容
	Category     string    `gorm:"column:category" json:"category"`              // 帖子分类
	LikeCount    int       `gorm:"column:like_count" json:"like_count"`          // 点赞数
	DislikeCount int       `gorm:"column:dislike_count" json:"dislike_count"`    // 踩数
	ViewCount    int       `gorm:"column:view_count" json:"view_count"`          // 阅读数
	EditTime     time.Time `gorm:"column:edit_time" json:"edit_time"`            // 编辑时间
}

// TableName 指定PostDO对应的数据库表名
func (PostDO) TableName() string {
	return "post"
}

// TransformToDO 将PostDTO转换为PostDO
func (p *PostDTO) TransformToDO() *PostDO {
	return &PostDO{
		Id:           p.Id,
		AuthorId:     p.Author.Id,
		Title:        p.Title,
		Content:      p.Content,
		Category:     p.Category,
		LikeCount:    p.LikeCount,
		DislikeCount: p.DislikeCount,
		ViewCount:    p.ViewCount,
		EditTime:     p.EditTime,
	}
}

// TransformToDTO 将PostDO转换为PostDTO
func (p *PostDO) TransformToDTO(author *UserDTO, comments []*PostCommentDTO) *PostDTO {
	return &PostDTO{
		Id:           p.Id,
		Author:       author,
		Title:        p.Title,
		Content:      p.Content,
		Category:     p.Category,
		LikeCount:    p.LikeCount,
		DislikeCount: p.DislikeCount,
		ViewCount:    p.ViewCount,
		EditTime:     p.EditTime,
		Comments:     comments,
	}
}

// PostCommentDTO 帖子评论DTO
type PostCommentDTO struct {
	Id        int64     `json:"id"`         // 评论ID
	PostId    int64     `json:"post_id"`    // 所属帖子ID
	Author    UserDTO   `json:"author"`     // 评论者信息
	Content   string    `json:"content"`    // 评论内容
	LikeCount int       `json:"like_count"` // 点赞数
	EditTime  time.Time `json:"edit_time"`  // 评论时间
}

// PostCommentDO 帖子评论DO
type PostCommentDO struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"` // 评论ID
	PostId    int64     `gorm:"column:post_id" json:"post_id"`                // 所属帖子ID
	AuthorId  int64     `gorm:"column:author_id" json:"author_id"`            // 评论者ID
	Content   string    `gorm:"column:content" json:"content"`                // 评论内容
	LikeCount int       `gorm:"column:like_count" json:"like_count"`          // 点赞数
	EditTime  time.Time `gorm:"column:edit_time" json:"edit_time"`            // 评论时间
}

// TableName 指定PostCommentDO对应的数据库表名
func (PostCommentDO) TableName() string {
	return "post_comment"
}

// TransformToDO 将PostCommentDTO转换为PostCommentDO
func (p *PostCommentDTO) TransformToDO() *PostCommentDO {
	return &PostCommentDO{
		Id:        p.Id,
		PostId:    p.PostId,
		AuthorId:  p.Author.Id,
		Content:   p.Content,
		LikeCount: p.LikeCount,
		EditTime:  p.EditTime,
	}
}

// TransformToDTO 将PostCommentDO转换为PostCommentDTO
func (p *PostCommentDO) TransformToDTO(author *UserDTO) *PostCommentDTO {
	return &PostCommentDTO{
		Id:        p.Id,
		PostId:    p.PostId,
		Author:    *author,
		Content:   p.Content,
		LikeCount: p.LikeCount,
		EditTime:  p.EditTime,
	}
}

// CreatePostRequestDTO 创建帖子请求DTO
type CreatePostRequestDTO struct {
	Title    string `json:"title"`    // 帖子标题
	Content  string `json:"content"`  // 帖子内容
	Category string `json:"category"` // 帖子分类
	UserId   int64  `json:"user_id"`  // 发布者ID
}

// CreatePostResponseDTO 创建帖子响应DTO
type CreatePostResponseDTO struct {
	PostId int64 `json:"post_id"` // 帖子ID
}

// CreateCommentRequestDTO 创建评论请求DTO
type CreateCommentRequestDTO struct {
	PostId  int64  `json:"post_id"` // 所属帖子ID
	Content string `json:"content"` // 评论内容
	UserId  int64  `json:"user_id"` // 评论者ID
}

// CreateCommentResponseDTO 创建评论响应DTO
type CreateCommentResponseDTO struct {
	CommentId int64 `json:"comment_id"` // 评论ID
}

// LikeRequestDTO 点赞请求DTO
type LikeRequestDTO struct {
	PostId int64 `json:"post_id"` // 帖子ID
	UserId int64 `json:"user_id"` // 用户ID
	IsLike bool  `json:"is_like"` // 是否点赞（true为点赞，false为踩）
}

// LikeResponseDTO 点赞响应DTO
type LikeResponseDTO struct {
	LikeCount    int `json:"like_count"`    // 更新后的点赞数
	DislikeCount int `json:"dislike_count"` // 更新后的踩数
}
