package main

import (
	"context"
	comment "simpledouyin/kitex_gen/douyin/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// ActionComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ActionComment(ctx context.Context, req *comment.ActionCommentRequest) (resp *comment.ActionCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// ListComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ListComment(ctx context.Context, req *comment.ListCommentRequest) (resp *comment.ListCommentResponse, err error) {
	// TODO: Your code here...
	return
}

// CountComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CountComment(ctx context.Context, req *comment.CountCommentRequest) (resp *comment.CountCommentResponse, err error) {
	// TODO: Your code here...
	return
}
