package service

import (
	"douyin/app/dao"
	"douyin/app/errs"
	"douyin/pkg/assert"
)

func CommentAction(userId int64, videoId int64, commentText string, actionType int, commentId int64) (*dao.Comment, error) {
	//发布评论
	if actionType == 1 {
		//检查评论是否为空
		if err := assert.NotEmpty(commentText); err != nil {
			return nil, err
		}

		comment, err := dao.SaveComment(userId, videoId, commentText)
		return comment, err
	}

	//删除评论
	if actionType == 2 {
		if commentId <= 0 {
			return nil, errs.CommentIdNotFound
		}

		err := dao.DeleteCommet(commentId)
		return nil, err
	}

	return nil, errs.ActionTypeNotFound
}

func CommentList(videoId int64) ([]dao.Comment, error) {
	if videoId <= 0 {
		return nil, errs.VideoIdNotFound
	}

	comments, err := dao.GetComments(videoId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
