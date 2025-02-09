package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yujian-backend/pkg/db"
	"yujian-backend/pkg/model"
)


func SearchBooks(c *gin.Context) {
	var req model.BookSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request parameters",
		})
		return
	}
	//默认值
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	bookRepository := db.GetBookRepository()
	books, err := bookRepository.SearchBooks(req.Keyword, req.Category, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to search books",
		})
		return
	}
	c.JSON(http.StatusOK, model.SearchResponse{
		Books: books,
	})
}

// GetBookDetail 图书详情获取接口
func GetBookDetail(c *gin.Context) {
	//获取id
	bookId, err := strconv.ParseInt(c.Param("bookId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid book ID",
		})
		return
	}

	bookRepository := db.GetBookRepository()
	// 查询详情
	bookDTO, err := bookRepository.GetBookById(bookId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "book not found",
		})
		return
	}
	// 返回
	c.JSON(http.StatusOK, gin.H{
		"data": bookDTO,
	})
}
