package main

import (
	"context"
	favorite "simpledouyin/kitex_gen/douyin/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteRequest) (resp *favorite.FavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// IsFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) IsFavorite(ctx context.Context, req *favorite.IsFavoriteRequest) (resp *favorite.IsFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// CountFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CountFavorite(ctx context.Context, req *favorite.CountFavoriteRequest) (resp *favorite.CountFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// CountUserFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CountUserFavorite(ctx context.Context, req *favorite.CountUserFavoriteRequest) (resp *favorite.CountUserFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// CountUserTotalFavorited implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) CountUserTotalFavorited(ctx context.Context, req *favorite.CountUserTotalFavoritedRequest) (resp *favorite.CountUserTotalFavoritedResponse, err error) {
	// TODO: Your code here...
	return
}
