package data

import (
	"CleanArch/features/book"
	"errors"
	"log"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.BookData {
	return &bookData{
		db: db,
	}
}

// Add implements book.BookData
func (bd *bookData) Add(userID int, newBook book.Core) (book.Core, error) {
	cnv := CoreToData(newBook)
	cnv.UserID = uint(userID)
	err := bd.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return book.Core{}, errors.New("querry error,fail to add item")
	}
	newBook.ID = cnv.ID
	return newBook, nil
}

// BookList implements book.BookData
func (bd *bookData) BookList() ([]book.Core, error) {
	res := []Books{}
	err := bd.db.Find(&res).Error
	if err != nil {
		log.Println("no data found")
		return []book.Core{}, errors.New("data not found")
	}
	result := []book.Core{}
	for i := 0; i < len(res); i++ {
		temp := res[i]
		result = append(result, DataToCore(temp))
		qry := User{}
		err := bd.db.Where("id=?", res[i].UserID).First(&qry).Error
		if err != nil {
			log.Println("no data found")
			return []book.Core{}, errors.New("data not found")
		}
		result[i].Pemilik = qry.Name
	}
	// log.Println(result)
	return result, nil
}

// Update implements book.BookData
func (bd *bookData) Update(tokenUserID int, bookID int, updatedData book.Core) (book.Core, error) {
	// cek apakah yang akan diedit milik user itu
	check := []Books{}
	err := bd.db.Where("id=? AND user_id=?", bookID, tokenUserID).Find(&check).Error
	if err != nil {
		log.Println("query error", err.Error())
		return book.Core{}, errors.New("book not found,update fail")
	}
	if len(check) == 0 {
		return book.Core{}, errors.New("book not found")
	}

	cnv := CoreToData(updatedData)
	qry := bd.db.Where("id = ?", bookID).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("no book was update")
		return book.Core{}, errors.New("update fail, server error")
	}
	if err := qry.Error; err != nil {
		log.Println("update book query error :", err.Error())
		return book.Core{}, errors.New("query error, problem with server")
	}
	return DataToCore(cnv), nil
}

// MyBook implements book.BookData
func (bd *bookData) MyBook(userID int) ([]book.Core, error) {
	res := []Books{}
	err := bd.db.Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		log.Println("no result")
		return []book.Core{}, errors.New("data not found")
	}
	result := []book.Core{}
	for i := 0; i < len(res); i++ {
		temp := res[i]
		result = append(result, DataToCore(temp))
	}
	// log.Println(result)
	return result, nil
}

// Delete implements book.BookData
func (bd *bookData) Delete(userID int, bookID int) error {
	//cek kemungkinan apabila user salah menghapus buku milik user lain
	check := []Books{}
	err := bd.db.Where("id=? AND user_id=?", bookID, userID).Find(&check).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("book not found,fail deleting")
	}
	if len(check) == 0 {
		return errors.New("no book has delete")
	}
	qry := bd.db.Unscoped().Delete(&Books{}, bookID) //Hard Delete
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no book has delete")
	}
	err = qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("server error")
	}
	return nil
}
