package data

import (
	"CleanArch/features/book"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Judul       string
	TahunTerbit int
	Penulis     string
	UserID      uint
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Address  string
	HP       string
	Password string
}

func DataToCore(data Books) book.Core {
	return book.Core{
		ID:          data.ID,
		Judul:       data.Judul,
		TahunTerbit: data.TahunTerbit,
		Penulis:     data.Penulis,
	}
}

func CoreToData(core book.Core) Books {
	return Books{
		Model:       gorm.Model{ID: core.ID},
		Judul:       core.Judul,
		Penulis:     core.Penulis,
		TahunTerbit: core.TahunTerbit,
	}
}
