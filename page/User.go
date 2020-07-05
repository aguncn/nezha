package page

//User 用户登录信息结构体
type User struct {
	UserID       uint   `json:"UserId"`
	UserType     uint   `json:"UserType"`
	Introduction string `json:"introduction"`
	UserName     string `json:"UserName"`
}

//Users 用户管理结构体
type Users struct {
	ID        uint   `json:"id"`
	Name      string `json:"username"`
	Password  string `json:"password"`
	UserType  string `json:"user_type"`
	State     string `json:"state"`
	Deleted   string `json:"deteled"`
	CreatedAt string `json:"created_at"`
}
