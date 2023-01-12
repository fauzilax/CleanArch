package data

import (
	"CleanArch/features/user"
	"errors"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// Register diambil dari kontrak UserData berdasarkan func New
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	cekDupe := CoreToData(newUser)
	err := uq.db.Where("email=?", cekDupe.Email).First(&cekDupe).Error
	if err == nil {
		log.Println("email already registered")
		return user.Core{}, errors.New("duplicated")
	}
	convert := CoreToData(newUser)
	err = uq.db.Create(&convert).Error
	if err != nil {
		log.Println("error insert register", err.Error())
		return user.Core{}, err
	}
	newUser.ID = convert.ID
	return newUser, nil
}

// Login diambil dari kontrak UserData berdasarkan func New
func (uq *userQuery) Login(email string) (user.Core, error) {
	res := User{}
	// Lakukan Query sekaligus cek apakah error
	err := uq.db.Where("email=?", email).First(&res).Error
	if err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}
	return DataToCore(res), nil
}

// Profile diambil dari kontrak UserData berdasarkan func New
func (uq *userQuery) Profile(id uint) (user.Core, error) {
	res := User{}
	err := uq.db.Where("id=?", id).First(&res).Error
	if err != nil {
		log.Println("result find data from id error", err.Error())
		return user.Core{}, errors.New("data not found")
	}
	return DataToCore(res), nil
}

// Update implements user.UserData
func (uq *userQuery) Update(id int, updateData user.Core) (user.Core, error) {
	res := CoreToData(updateData)
	qry := uq.db.Where("id = ?", id).Updates(&res)
	if qry.RowsAffected <= 0 {
		log.Println("update book query error : data not found")
		return user.Core{}, errors.New("not found")
	}
	err := qry.Error
	if err != nil {
		log.Println("update book query error :", err.Error())
		return user.Core{}, err
	}
	return DataToCore(res), nil
}

func (uq *userQuery) Delete(userID int) error {
	qry := uq.db.Unscoped().Delete(&User{}, userID)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}
	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}
	return nil
}
