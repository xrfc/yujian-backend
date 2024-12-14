package post

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"yujian-backend/pkg/db"
	"yujian-backend/pkg/log"
	"yujian-backend/pkg/model"
	"yujian-backend/pkg/utils"
)

var postBizInstance *PostBiz

// PostBiz 帖子业务逻辑
type PostBiz struct {
	postRepo *db.PostRepository
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取参数
		var req model.CreatePostRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// todo 这里外层的逻辑需要修改
		resp, err := postBizInstance.CreatePost(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)

		return
	}
}

// CreatePost 创建帖子
func (b *PostBiz) CreatePost(req *model.CreatePostRequestDTO) (*model.CreatePostResponseDTO, error) {
	resp := &model.CreatePostResponseDTO{
		BaseResp: model.BaseResp{
			Code: model.Success,
		},
	}

	// 参数校验
	if req.Title == "" || req.Content == "" {
		resp.Code = model.UserNotExists
		resp.ErrMsg = "标题或内容不能为空"
		return resp, errors.New("标题或内容不能为空")
	}

	// 生成内容ID
	contentId := b.generateContentId(req.Title, req.UserId)

	// 构建帖子DO
	postDTO := &model.PostDTO{
		Title:     req.Title,
		ContentId: contentId,
		Author: &model.UserDTO{
			Id: req.UserId,
		},
		EditTime: time.Now(),
		Comments: []*model.PostCommentDTO{},
	}

	// 保存帖子
	if id, err := b.postRepo.CreatePost(postDTO); err != nil {
		log.GetLogger().Error("创建帖子失败: %v", err)
		resp.Code = model.UserNotExists
		resp.ErrMsg = "创建帖子失败"
		resp.Error = err
		return resp, err
	} else {
		resp.PostId = id
	}

	return resp, nil
}

func (b *PostBiz) generateContentId(title string, uid int64) string {
	return title + strconv.FormatInt(uid, 10) + utils.GenerateUUID()
}
