package model

// UserDTO `用户`DTO结构体
type UserDTO struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserDO `用户`存储数据结构体
type UserDO struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (userDTO *UserDTO) Transfer() *UserDO {
	return &UserDO{
		Id:       userDTO.Id,
		Email:    userDTO.Email,
		Name:     userDTO.Name,
		Password: userDTO.Password,
	}
}

func (userDO *UserDO) Transfer() *UserDTO {
	return &UserDTO{
		Id:       userDO.Id,
		Email:    userDO.Email,
		Name:     userDO.Name,
		Password: userDO.Password,
	}
}


// 修改密码的请求结构体
type ChangePasswordRequest struct {
	OldPassword     string `json:"OldPassword"`     // 旧密码
	NewPassword     string `json:"NewPassword"`     // 新密码
	ConfirmPassword string `json:"ConfirmPassword"` // 确认新密码
}
