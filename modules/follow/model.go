package follow

type Follow struct {
	Id          int `json:"id"`
	FollowerId  int `json:"follower_id"`
	FollowingId int `json:"following_id"`
}

type FollowWithName struct {
	Follow
	FollowingUsername string `json:"following_username"`
	FollowerUsername  string `json:"follower_username"`
}

type FollowingList struct {
	Id            int    `json:"id"`
	FollowingId   int    `json:"following_user_id"`
	FollowingName string `json:"following_username"`
}

type FollowerList struct {
	Id           int    `json:"id"`
	FollowerId   int    `json:"follower_user_id"`
	FollowerName string `json:"follower_username"`
}
