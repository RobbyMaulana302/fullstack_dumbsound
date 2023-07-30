package authdto

type AuthRegister struct {
	Email    string `json:"email" form:"email" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Password string `json:"password"  form:"password" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	ListAs   bool   `json:"status"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	Token    string `json:"token"`
}
