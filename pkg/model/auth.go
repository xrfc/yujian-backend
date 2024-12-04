package model

type LoginRequestDTO struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string    `json:"token"`
	User  UserDTO   `json:"user"`
	Error error     `json:"error"`
	Code  ErrorCode `json:"code"`
}

type RegisterRequestDTO struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterResponseDTO struct {
	Token string    `json:"token"`
	User  UserDTO   `json:"user"`
	Error error     `json:"error"`
	Code  ErrorCode `json:"code"`
}
