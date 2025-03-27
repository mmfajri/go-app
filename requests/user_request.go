package requests

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Role     string `json:"role" gorm:"-"`
	Password string `json:"password"`
	BaseRequest
}
