package model

type LoginRequestDTO struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	BaseResp
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}

type RegisterRequestDTO struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterResponseDTO struct {
	BaseResp
	Token string  `json:"token"`
	User  UserDTO `json:"user"`
}
