package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"yujian-backend/pkg/db"
	"yujian-backend/pkg/log"
	"yujian-backend/pkg/model"
)

var postBizInstance *PostBiz

// PostBiz 帖子业务逻辑
type PostBiz struct {
	postRepo *db.PostRepository
}

// CreatePost 发布帖子
func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户ID
		userID := c.MustGet("userID").(int64)

		// 解析请求参数
		var req model.CreatePostRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
			return
		}

		// 参数校验
		if req.Title == "" || req.Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "标题和内容不能为空"})
			return
		}

		// 构建帖子DTO
		postDTO := &model.PostDTO{
			Title:    req.Title,
			Content:  req.Content,
			Category: req.Category,
			Author:   &model.UserDTO{Id: userID},
			EditTime: time.Now(),
		}

		// 保存帖子
		if err := db.GetPostRepository().CreatePost(postDTO); err != nil {
			log.GetLogger().Errorf("创建帖子失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "发布失败"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"message": "帖子发布成功",
			"post_id": postDTO.Id,
		})
	}
}

// ListPosts 获取帖子列表
func ListPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析查询参数
		category := c.Query("category")
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
		offset := (page - 1) * size

		// 查询帖子列表
		posts, err := db.GetPostRepository().ListPosts(category, offset, size)
		if err != nil {
			log.GetLogger().Errorf("获取帖子列表失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取失败"})
			return
		}

		// 返回帖子列表
		c.JSON(http.StatusOK, gin.H{
			"data": posts,
		})
	}
}

// GetPostDetail 获取帖子详情
func GetPostDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析帖子ID
		postID := c.Param("postId")
		id, err := strconv.ParseInt(postID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "帖子ID无效"})
			return
		}

		// 增加阅读数
		db.GetPostRepository().IncrementViewCount(id)

		// 查询帖子详情
		post, err := db.GetPostRepository().GetPostById(id)
		if err != nil {
			log.GetLogger().Errorf("获取帖子详情失败: %v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "帖子不存在"})
			return
		}

		// 返回帖子详情
		c.JSON(http.StatusOK, gin.H{
			"data": post,
		})
	}
}

// CreateComment 发布评论
func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户ID
		userID := c.MustGet("userID").(int64)

		// 解析帖子ID
		postID := c.Param("postId")
		postId, err := strconv.ParseInt(postID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "帖子ID无效"})
			return
		}

		// 解析请求参数
		var req struct {
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&req); err != nil || req.Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "评论内容不能为空"})
			return
		}

		// 构建评论DTO
		comment := &model.PostCommentDTO{
			PostId:   postId,
			Author:   model.UserDTO{Id: userID},
			Content:  req.Content,
			EditTime: time.Now(),
		}

		// 保存评论
		if err := db.GetCommentRepository().Create(comment); err != nil {
			log.GetLogger().Errorf("发布评论失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "评论失败"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"message":    "评论成功",
			"comment_id": comment.Id,
		})
	}
}

// HandleLike 点赞/踩
func HandleLike() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析帖子ID
		postID := c.Param("postId")
		postId, err := strconv.ParseInt(postID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "帖子ID无效"})
			return
		}

		// 解析操作类型（like/dislike）
		action := c.Query("action")
		isLike := action == "like"

		// 执行点赞/踩操作
		if err := db.GetPostRepository().HandleLike(postId, isLike); err != nil {
			log.GetLogger().Errorf("点赞/踩操作失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "操作失败"})
			return
		}

		// 查询更新后的点赞数和踩数
		post, err := db.GetPostRepository().GetPostById(postId)
		if err != nil {
			log.GetLogger().Errorf("获取帖子信息失败: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取数据失败"})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"message":       "操作成功",
			"like_count":    post.LikeCount,
			"dislike_count": post.DislikeCount,
		})
	}
}
