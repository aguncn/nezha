package page

//Application 用户管理结构体
type Application struct {
	ID        uint   `json:"ID"`
	Name      string `json:"name"`
	Git       string `json:"git"`
	Jenkins   string `json:"jenkins"`
	State     uint   `json:"state"`
	CreatedAt string `json:"created_at"`
}
