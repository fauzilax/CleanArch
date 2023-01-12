package handler

import "CleanArch/features/user"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	HP       string `json:"hp" form:"hp"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
	HP      string `json:"hp" form:"hp"`
}

func RequestToCore(dataUser interface{}) *user.Core {
	res := user.Core{}
	switch dataUser.(type) {
	case LoginRequest:
		cnv := dataUser.(LoginRequest)
		res.Email = cnv.Email
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := dataUser.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.HP = cnv.HP
		res.Address = cnv.Address
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := dataUser.(UpdateRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.HP = cnv.HP
		res.Address = cnv.Address
	default:
		return nil
	}

	return &res
}
