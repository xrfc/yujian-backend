package model

import (
	"time"
)

// PostDTO 帖子DTO
type PostDTO struct {
	Id        int64             `json:"id"`
	Author    *UserDTO          `json:"author"`
	标题     string            `json:"title"`
	ContentId string            `json:"content_id"`
	EditTime  time.Time         `json:"edit_time"`
	Comments  []*PostCommentDTO `json:"comments"`
	LikeUserIds    []int64   `json:"like_user_ids"`    // 点赞的用户ID列表
	DislikeUserIds []int64   `json:"dislike_user_ids"` // 点踩的用户ID列表
}

// TransformToDO 将PostDTO转换为PostDO
func (p *PostDTO) TransformToDO() *PostDO {
	return &PostDO{
		Id:         p.Id,
		AuthorId:   p.Author.Id,
		AuthorName: p.Author.Name,
		Title:      p.Title,
		ContentId:  p.ContentId,
		EditTime:   p.EditTime,
	}
}

// PostDO 帖子DO
type PostDO struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	AuthorId   int64     `gorm:"column:author_id" json:"author_id"`
	AuthorName string    `gorm:"column:author_name" json:"author_name"`
	标题      string    `gorm:"column:title" json:"title"`
	ContentId  string    `gorm:"column:content_id" json:"content_id"`
	EditTime   time.Time `gorm:"column:edit_time" json:"edit_time"`
	LikeUserIds    string    `gorm:"column:like_user_ids" json:"like_user_ids"`
	DislikeUserIds string    `gorm:"column:dislike_user_ids" json:"dislike_user_ids"`
}

func (p PostDO) TableName() string {
	return "post"
}

// TransformToDTO 将PostDO转换为PostDTO
func (p *PostDO) TransformToDTO(userDTO *UserDTO, comments []*PostCommentDTO) *PostDTO {
	return &PostDTO{
		Id:        p.Id,
		Author:    userDTO,
		Title:     p.Title,
		ContentId: p.ContentId,
		EditTime:  p.EditTime,
		Comments:  comments,
	}
}

// PostEsModel 帖子ES模型
type PostEsModel struct {
	Id      string  `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Score   float64 `json:"score"`
}

func (p *PostEsModel) GetID() string {
	return p.Id
}

func (p *PostEsModel) SetScore(score float64) {
	p.Score = score
}

func (p *PostEsModel) GetScore() float64 {
	return p.Score
}

func (p *PostEsModel) GetIndexName() string {
	return "post"
}

func (p *PostEsModel) GetContent() string {
	return p.Content
}

func (p *PostEsModel) GetTitle() string {
	return p.Title
}

// PostCommentDTO 帖子评论DTO
type PostCommentDTO struct {
	Id             int64     `json:"id"`
	PostId         int64     `json:"post_id"`
	Author         UserDTO   `json:"author"`
	EditTime       time.Time `json:"edit_time"`
	Content        string    `json:"content"`          // 评论的内容不会很长,直接存mysql
	Score          int       `json:"score"`            // 评论的分数
	LikeUserIds    []int64   `json:"like_user_ids"`    // 点赞的用户ID列表
	DislikeUserIds []int64   `json:"dislike_user_ids"` // 点踩的用户ID列表
}

// PostCommentDO 帖子评论DO
type PostCommentDO struct {
	Id             int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PostId         int64     `gorm:"column:post_id" json:"post_id"`
	AuthorId       int64     `gorm:"column:author_id" json:"author_id"`
	AuthorName     string    `gorm:"column:author_name" json:"author_name"`
	EditTime       time.Time `gorm:"column:edit_time" json:"edit_time"`
	Content        string    `gorm:"column:content" json:"content"` // 评论的内容不会很长,直接存mysql
	Score          int       `gorm:"column:score" json:"score"`
	LikeUserIds    string    `gorm:"column:like_user_ids" json:"like_user_ids"`
	DislikeUserIds string    `gorm:"column:dislike_user_ids" json:"dislike_user_ids"`
}

func (p PostCommentDO) TableName() string {
	return "post_comment"
}

// TransformToDTO 将PostCommentDO转换为PostCommentDTO
func (p *PostCommentDO) TransformToDTO() *PostCommentDTO {
	return &PostCommentDTO{
		Id:       p.Id,
		PostId:   p.PostId,
		Author:   UserDTO{Id: p.AuthorId, Name: p.AuthorName},
		EditTime: p.EditTime,
		Content:  p.Content,
	}
}

func (p *PostCommentDTO) TransformToDO() *PostCommentDO {
	return &PostCommentDO{
		Id:         p.Id,
		PostId:     p.PostId,
		AuthorId:   p.Author.Id,
		AuthorName: p.Author.Name,
		EditTime:   p.EditTime,
		Content:    p.Content,
	}
}

// CreatePostRequestDTO 创建帖子请求DTO
type CreatePostRequestDTO struct {
	标题   string `json:"title"`
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
}

// CreatePostResponseDTO 创建帖子响应DTO
type CreatePostResponseDTO struct {
	BaseResp
	PostId int64 `json:"post_id"`
}
