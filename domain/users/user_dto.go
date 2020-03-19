package users

type User struct {
	Id         int64  `json:"id"`
	CompanyId  int64  `json:"company_id"`
	UserId     string `json:"user_id"`
	AppKey     string `json:"app_key,omitempty"`
	Role       string `json:"role"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeleteddAt string `json:"-"`
}

type Users []User

type ResolveResponse struct {
	CompanuName  string `json:"company_name"`
	CompanyId    int64  `json:"company_id"`
	Domain       string `json:"domain"`
	UserId       string `json:"user_id,omitempty"`
	Role         string `json:"role"`
	Subordinates Users  `json:"subordinates"`
}
