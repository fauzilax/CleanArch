package services

import (
	"CleanArch/features/user"
	"CleanArch/helper"
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userServiceCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userServiceCase{
		qry: ud,
	}
}

// Register diambil berdasarkan kontrak UserService di func New
func (usc *userServiceCase) Register(newUser user.Core) (user.Core, error) {
	hashInpPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashInpPassword)
	res, err := usc.qry.Register(newUser)
	// log.Println(res,err)
	if err != nil {
		if strings.Contains(err.Error(), "duplicated") {
			return user.Core{}, errors.New("data already exist")
		}
		return user.Core{}, errors.New("internal server error")
	}
	log.Println("OK")
	return res, nil
}

// Login diambil berdasarkan kontrak UserService di func New
func (usc *userServiceCase) Login(email string, password string) (string, user.Core, error) {
	res, err := usc.qry.Login(email)
	if err != nil {
		return "", user.Core{}, errors.New("data not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password not matched")
	}
	useToken, _ := helper.GenerateToken(int(res.ID))

	return useToken, res, nil
}

// Profile diambil berdasarkan kontrak UserService di func New
func (usc *userServiceCase) Profile(tokenIDUser interface{}) (user.Core, error) {
	id := helper.ExtractToken(tokenIDUser)
	if id <= 0 {
		return user.Core{}, errors.New("data not found")
	}
	res, err := usc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "something wrong with server please try again later"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

// Update implements user.UserService
func (usc *userServiceCase) Update(tokenIDUser interface{}, updateData user.Core) (user.Core, error) {
	id := helper.ExtractToken(tokenIDUser)
	res, err := usc.qry.Update(id, updateData)
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("query error, update fail")
	}
	return res, nil
}

// Delete implements user.UserService
func (usc *userServiceCase) Delete(tokenIDUser interface{}) error {
	userID := helper.ExtractToken(tokenIDUser)
	if userID <= 0 {
		return errors.New("data not found")
	}
	err := usc.qry.Delete(userID)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}
