package services

import (
	"CleanArch/features/book"
	"CleanArch/helper"
	"errors"
	"log"
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
	// error token tidak dibutuhkan karena token pasti memiliki user id karena untuk add buku dan sebagainya
	// user perlu login terlebih dahulu sehingga pasti ada id
	// if userID <= 0 {
	// 	return book.Core{}, errors.New("user tidak ditemukan")
	// }
	res, err := bs.data.Add(userID, newBook)
	if err != nil {
		return book.Core{}, errors.New("something wrong happens,server error")
	}
	return res, nil
}

// Update implements book.BookService
func (bs *bookSrv) Update(token interface{}, bookID int, updatedData book.Core) (book.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := bs.data.Update(userID, bookID, updatedData)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "not found") {
			return book.Core{}, errors.New("book not found")
		}
		return book.Core{}, errors.New("internal server error")

	}
	return res, nil
}

// BookList implements book.BookService
func (bs *bookSrv) BookList() ([]book.Core, error) {
	res, err := bs.data.BookList()
	if err != nil {
		log.Println("no result or server error")
		return []book.Core{}, errors.New("no result or server error")
	}
	// fmt.Println(res)
	return res, nil
}

// MyBook implements book.BookService
func (bs *bookSrv) MyBook(token interface{}) ([]book.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := bs.data.MyBook(userID)
	if err != nil {
		return []book.Core{}, errors.New("data not found")
	}
	return res, nil
}

// Delete implements book.BookService
func (bs *bookSrv) Delete(token interface{}, bookID int) error {
	userID := helper.ExtractToken(token)
	err := bs.data.Delete(userID, bookID)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "book") {
			msg = "user dont have any book"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}
	return nil
}
