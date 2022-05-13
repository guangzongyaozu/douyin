package api

import (
	"douyin/app/errs"
	"douyin/pkg/com"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type VideoListResponse struct {
	com.Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; !exist {
		com.Error(c, errs.UserNotFound)
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		com.Error(c, err)
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		com.Error(c, err)
		return
	}

	com.Success(c, &com.Response{
		StatusMsg: finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	com.Success(c, &VideoListResponse{
		VideoList: DemoVideos,
	})
}
