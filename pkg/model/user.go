package model

// UserDTO `用户`DTO结构体
type UserDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserDO `用户`存储数据结构体
type UserDO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (userDTO *UserDTO) Transfer() *UserDO {
	return &UserDO{
		Id:       userDTO.Id,
		Name:     userDTO.Name,
		Password: userDTO.Password,
	}
}

func (userDO *UserDO) Transfer() *UserDTO {
	return &UserDTO{
		Id:       userDO.Id,
		Name:     userDO.Name,
		Password: userDO.Password,
	}
}
