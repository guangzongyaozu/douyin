package errs

import "douyin/pkg/com"

const (
	CodeActionTypeNotFound = BaseComment + iota
	CodeCommentIdNotFound
	CodeVideoIdNotFound
)

var (
	ActionTypeNotFound = com.NewAPIError(CodeActionTypeNotFound, "ActionType not found")
	CommentIdNotFound  = com.NewAPIError(CodeCommentIdNotFound, "commentId not found")
	VideoIdNotFound    = com.NewAPIError(CodeVideoIdNotFound, "videoId not found")
)
