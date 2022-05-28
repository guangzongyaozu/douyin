package api

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/app/service"
	"douyin/pkg/com"
	"douyin/pkg/security"
	"douyin/pkg/validate"

	"github.com/gin-gonic/gin"
)

type CommentRequest struct {
	VideoId     int64  `form:"video_id"`
	ActionType  int    `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type CommentResponse struct {
	com.Response
	Comment dao.Comment `json:"comment"`
}

type CommentListResponse struct {
	com.Response
	CommentList []dao.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	myUserId := security.GetUserId(c)

	rq := validate.StructQuery(c, &CommentRequest{})
	if rq == nil {
		return
	}

	if myUserId <= 0 {
		com.Error(c, errs.UserNotFound)
		return
	}

	comment, err := service.CommentAction(myUserId, rq.VideoId, rq.CommentText, rq.ActionType, rq.CommentId)
	if err != nil {
		com.Error(c, err)
		return
	}

	if comment != nil {
		comment.CreateDate = comment.CreatedAt.Format("01-02")
		com.Success(c, &CommentResponse{
			Comment: *comment,
		})
	} else {
		// com.Success(c, &CommentResponse{})
		com.SuccessStatus(c)
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	rq := validate.StructQuery(c, &CommentRequest{})
	if rq == nil {
		return
	}

	comments, err := service.CommentList(rq.VideoId)
	if err != nil {
		com.Error(c, err)
		return
	}

	for _, comment := range comments {
		comment.CreateDate = comment.CreatedAt.Format("01-02")
	}

	com.Success(c, &CommentListResponse{
		CommentList: comments,
	})
}
