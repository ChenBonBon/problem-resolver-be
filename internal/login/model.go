package login

type UsernameWithPassword struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type EmailWithCode struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type UsernameWithToken struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
