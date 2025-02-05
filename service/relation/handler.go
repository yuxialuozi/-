package main

import (
	"context"
	relation "simpledouyin/kitex_gen/douyin/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// Unfollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Unfollow(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.FollowListRequest) (resp *relation.FollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// CountFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollowList(ctx context.Context, req *relation.CountFollowListRequest) (resp *relation.CountFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.FollowerListRequest) (resp *relation.FollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// CountFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollowerList(ctx context.Context, req *relation.CountFollowerListRequest) (resp *relation.CountFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.FriendListRequest) (resp *relation.FriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// IsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollow(ctx context.Context, req *relation.IsFollowRequest) (resp *relation.IsFollowResponse, err error) {
	// TODO: Your code here...
	return
}
