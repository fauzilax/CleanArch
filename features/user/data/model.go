package data

import (
	"CleanArch/features/book/data"
	"CleanArch/features/user"

	"gorm.io/gorm"
)

// buat struct yang akan menjadi table di database
type User struct {
	gorm.Model
	Name     string
	Email    string
	Address  string
	HP       string
	Password string
	Book     []data.Books
}

// karena kolom database berbeda dari struct perjanjian maka harus dikonversi dengan membuat sebuah fungsi
// Pertama dari Struct User ke Struct Core
func DataToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Address:  data.Address,
		HP:       data.HP,
		Password: data.Password,
	}
}

// lakukan kembali fungsi conversi kebalikannya Core ke User
func CoreToData(core user.Core) User {
	return User{
		Model:    gorm.Model{ID: core.ID},
		Name:     core.Name,
		Email:    core.Email,
		Address:  core.Address,
		HP:       core.HP,
		Password: core.Password,
	}
}
