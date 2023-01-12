package book

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint
	Judul       string `validate:"required,alpha"`
	TahunTerbit int    `validate:"required"`
	Penulis     string `validate:"required,alpha"`
	Pemilik     string
}

type BookHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	BookList() echo.HandlerFunc
	Delete() echo.HandlerFunc
	MyBook() echo.HandlerFunc
}

type BookService interface {
	Add(token interface{}, newBook Core) (Core, error)
	Update(token interface{}, bookID int, updatedData Core) (Core, error)
	BookList() ([]Core, error)
	Delete(token interface{}, bookID int) error
	MyBook(token interface{}) ([]Core, error)
}

type BookData interface {
	Add(userID int, newBook Core) (Core, error)
	Update(token int, bookID int, updatedData Core) (Core, error)
	BookList() ([]Core, error)
	Delete(userID, bookID int) error
	MyBook(userID int) ([]Core, error)
}
