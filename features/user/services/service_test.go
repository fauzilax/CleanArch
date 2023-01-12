package services

import (
	"CleanArch/features/user"
	"CleanArch/helper"
	"CleanArch/mocks"
	"errors"
	"log"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t) // membuat interface data
	inputDataUser := user.Core{
		ID:       uint(0),
		Name:     "fauzi",
		Email:    "fauzi@gmail.com",
		Address:  "bandung",
		HP:       "0812345",
		Password: "123",
	}
	returnData := user.Core{
		ID:       uint(1),
		Name:     "fauzi",
		Email:    "fauzi@gmail.com",
		Address:  "bandung",
		HP:       "0812345",
		Password: "123",
	}
	t.Run("Berhasil Register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(returnData, nil).Once() // On(method , input/parameter) . Return(Output/Return Value) | Once() digunakan untuk memastikan data yang di tes lebih safety tidak memanggil berulang di On().Return() dibawahnya
		srv := New(data)                                                  //ambil data palsu
		res, err := srv.Register(inputDataUser)                           // Jalankan Register di Services
		log.Println(res.Password)
		assert.Nil(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		assert.Equal(t, returnData.Name, res.Name)
		assert.Equal(t, returnData.Password, res.Password)
		data.AssertExpectations(t) // memastikan apakah ujicoba berjalan sesuai dengan On().Return()
	})
	t.Run("Register Gagal", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("internal server error")).Once()
		srv := New(data)
		res, err := srv.Register(user.Core{})
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", res.Name)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
	t.Run("Duplicated", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()
		srv := New(data)
		res, err := srv.Register(user.Core{})
		assert.ErrorContains(t, err, "already exist")
		assert.Equal(t, "", res.Password)
		data.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t) // mock data
	// input dan respond untuk mock data
	inputEmail := "fauzi@gmail.com"
	hashed, _ := helper.GeneratePassword("123")
	// res dari data akan mengembalikan password yang sudah di hash
	resData := user.Core{ID: uint(1), Name: "jerry", Email: "fauzi@gmail.com", HP: "08123456", Password: hashed}
	t.Run("Berhasil login", func(t *testing.T) {
		data.On("Login", inputEmail).Return(resData, nil).Once() // simulasi method login pada layer data
		srv := New(data)
		token, res, err := srv.Login(inputEmail, "123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("Tidak ditemukan", func(t *testing.T) {
		data.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(data)
		token, res, err := srv.Login(inputEmail, "123")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "jerry@alterra.id"
		hashed, _ := helper.GeneratePassword("be1422")
		resData := user.Core{ID: uint(1), Name: "jerry", Email: "jerry@alterra.id", HP: "08123456", Password: hashed}
		data.On("Login", inputEmail).Return(resData, nil).Once()

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1423")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not matched")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("Sukses lihat profile", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Name: "jerry", Email: "jerry@alterra.id", HP: "08123456"}

		repo.On("Profile", uint(1)).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		res, err := srv.Profile(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Profile", uint(4)).Return(user.Core{}, errors.New("data not found")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	data := mocks.NewUserData(t)
	inputData := user.Core{
		ID:      uint(0),
		Name:    "fauzi",
		Email:   "fauzi@gmail.com",
		Address: "bandung",
		HP:      "0812345",
	}
	resData := user.Core{
		ID:      uint(1),
		Name:    "fauzila",
		Email:   "fauzi@gmail.com",
		Address: "bandung",
		HP:      "0812345",
	}
	t.Run("Sukses Update", func(t *testing.T) {
		data.On("Update", int(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		token := tokenIDUser.(*jwt.Token)
		token.Valid = true
		res, err := srv.Update(token, inputData)
		assert.Equal(t, resData.ID, res.ID)
		assert.Nil(t, err)
		data.AssertExpectations(t)

	})
	t.Run("Update fail", func(t *testing.T) {
		data.On("Update", int(1), inputData).Return(resData, errors.New("query error")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		token := tokenIDUser.(*jwt.Token)
		token.Valid = true
		res, err := srv.Update(token, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "fail")
		assert.Equal(t, "", res.Name)
		data.AssertExpectations(t)
	})

}
func TestDelete(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("Success Delete Account", func(t *testing.T) {
		data.On("Delete", int(1)).Return(nil).Once()
		_, token := helper.GenerateToken(1)
		IDToken := token.(*jwt.Token)
		IDToken.Valid = true
		srv := New(data)
		err := srv.Delete(IDToken)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})
	t.Run("Delete Fail", func(t *testing.T) {
		data.On("Delete", int(1)).Return(errors.New("fail to delete")).Once()
		_, token := helper.GenerateToken(1)
		IDToken := token.(*jwt.Token)
		IDToken.Valid = true
		srv := New(data)
		err := srv.Delete(IDToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "fail")
		data.AssertExpectations(t)
	})
	t.Run("Token Atau ID salah", func(t *testing.T) {
		_, token := helper.GenerateToken(1)
		srv := New(data)
		err := srv.Delete(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

}
