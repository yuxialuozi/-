package model

type VideoList struct {
	NextTime  int     `json:"next_time"`  //本次返回的视频中，发布最早的时间
	VideoList []Video `json:"video_list"` //视频列表
}
