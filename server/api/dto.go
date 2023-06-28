package api

type LoginUserDTO struct {
	Username string `validate:"required,excludesall=0x20" json:"username"`
	Password string `validate:"required" json:"password"`
}

type UpdateCncDTO struct {
	Status string `validate:"required,oneof=working stopped broken" json:"status"`
}
