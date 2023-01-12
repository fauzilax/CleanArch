package services

import (
	"CleanArch/features/book"
	"CleanArch/helper"
	"errors"
	"strings"
)

type bookSrv struct {
	data book.BookData
}

func New(d book.BookData) book.BookService {
	return &bookSrv{
		data: d,
	}
}

// Add implements book.BookService
func (bs *bookSrv) Add(token interface{}, newBook book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return book.Core{}, errors.New("user tidak ditemukan")
	}
	res, err := bs.data.Add(userID, newBook)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "buku tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return book.Core{}, errors.New(msg)
	}
	return res, nil
}

// Update implements book.BookService
func (bs *bookSrv) Update(token interface{}, bookID int, updatedData book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)

	res, err := bs.data.Update(userID, bookID, updatedData)
	if err != nil {
		return book.Core{}, errors.New("buku tidak ditemukan,terjadi kesalahan pada server")
	}
	return res, nil
}

// BookList implements book.BookService
func (bs *bookSrv) BookList() ([]book.Core, error) {
	res, err := bs.data.BookList()
	if err != nil {
		return []book.Core{}, errors.New("buku tidak ditemukan")
	}
	// fmt.Println(res)
	return res, nil
}

// MyBook implements book.BookService
func (bs *bookSrv) MyBook(token interface{}) ([]book.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return []book.Core{}, errors.New("user tidak ditemukan")
	}
	res, err := bs.data.MyBook(userID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "buku tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return []book.Core{}, errors.New(msg)
	}
	return res, nil
}

// Delete implements book.BookService
func (bs *bookSrv) Delete(token interface{}, bookID int) error {
	userID := helper.ExtractToken(token)
	err := bs.data.Delete(userID, bookID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "buku tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return errors.New(msg)
	}
	return nil
}
