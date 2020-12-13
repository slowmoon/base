package auth2


type UserInfo struct {
	NickName   string
	AvatarUrl  string
	Sex        int    //1 for man 2 for woman
	Identify   string    // 唯一授权id
	Phone      string
	Email      string
}
