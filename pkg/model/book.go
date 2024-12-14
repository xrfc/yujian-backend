package model

import (
	"yujian-backend/pkg/utils"
)

// BookInfoDTO 书信息DTO
type BookInfoDTO struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	ISBN   string  `json:"ISBN"`
	Score  float64 `json:"score"`
	Intro  string  `json:"intro"`
}

// BookInfoDO 书信息数据库对象
type BookInfoDO struct {
	Id     int64   `gorm:"column:id;primaryKey" json:"id"`
	Name   string  `gorm:"column:name" json:"name"`
	Author string  `gorm:"column:author" json:"author"`
	ISBN   string  `gorm:"column:isbn" json:"ISBN"`
	Score  float64 `gorm:"column:score" json:"score"`
	Intro  string  `gorm:"column:intro" json:"intro"`
}

// TransformToDTO 将BookInfoDO转换为BookInfoDTO
func (bookInfoDO *BookInfoDO) Transfer() *BookInfoDTO {
	return &BookInfoDTO{
		Id:     bookInfoDO.Id,
		Name:   bookInfoDO.Name,
		Author: bookInfoDO.Author,
		ISBN:   bookInfoDO.ISBN,
		Score:  bookInfoDO.Score,
		Intro:  bookInfoDO.Intro,
	}
}

// TransformToDO 将BookInfoDTO转换为BookInfoDO
func (bookInfoDTO *BookInfoDTO) TransformToDO() *BookInfoDO {
	return &BookInfoDO{
		Id:     bookInfoDTO.Id,
		Name:   bookInfoDTO.Name,
		Author: bookInfoDTO.Author,
		ISBN:   bookInfoDTO.ISBN,
		Score:  bookInfoDTO.Score,
		Intro:  bookInfoDTO.Intro,
	}
}

// BookCommentDTO 书评DTO
type BookCommentDTO struct {
	Id             int64   `json:"id"`
	BookId         int64   `json:"book_id"`
	Content        string  `json:"content"`
	Like           int64   `json:"like"`
	Dislike        int64   `json:"dislike"`
	LikeUserIds    []int64 `json:"like_user_ids"`
	DislikeUserIds []int64 `json:"dislike_user_ids"`
}

// BookCommentDO 书评数据库对象
type BookCommentDO struct {
	Id             int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	BookId         int64  `gorm:"column:book_id" json:"book_id"`
	Content        string `gorm:"column:content" json:"content"`
	Like           int64  `gorm:"column:like" json:"like"`
	Dislike        int64  `gorm:"column:dislike" json:"dislike"`
	LikeUserIds    string `gorm:"column:like_user_ids" json:"like_user_ids"`
	DislikeUserIds string `gorm:"column:dislike_user_ids" json:"dislike_user_ids"`
}

// TransformToDO 将BookCommentDTO转换为BookCommentDO
func (bookCommentDTO *BookCommentDTO) Transfer() *BookCommentDO {
	return &BookCommentDO{
		Id:             bookCommentDTO.Id,
		BookId:         bookCommentDTO.BookId,
		Content:        bookCommentDTO.Content,
		Like:           bookCommentDTO.Like,
		Dislike:        bookCommentDTO.Dislike,
		LikeUserIds:    utils.MustToJSONString(bookCommentDTO.LikeUserIds),
		DislikeUserIds: utils.MustToJSONString(bookCommentDTO.DislikeUserIds),
	}
}

// TransformToDTO 将BookCommentDO转换为BookCommentDTO
func (bookCommentDO *BookCommentDO) TransformToDTO() *BookCommentDTO {
	return &BookCommentDTO{
		Id:             bookCommentDO.Id,
		BookId:         bookCommentDO.BookId,
		Content:        bookCommentDO.Content,
		Like:           bookCommentDO.Like,
		Dislike:        bookCommentDO.Dislike,
		LikeUserIds:    utils.MustParseJSONString[[]int64](bookCommentDO.LikeUserIds),
		DislikeUserIds: utils.MustParseJSONString[[]int64](bookCommentDO.DislikeUserIds),
	}
}
