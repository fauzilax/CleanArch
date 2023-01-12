package services

import (
	"CleanArch/features/book"
	"CleanArch/helper"
	"CleanArch/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	data := mocks.NewBookData(t)
	inputData := book.Core{
		ID:          uint(0),
		Judul:       "Avatar",
		TahunTerbit: 2003,
		Penulis:     "Aang",
	}
	resData := book.Core{
		ID:          uint(1),
		Judul:       "Avatar",
		TahunTerbit: 2003,
		Penulis:     "Aang",
		Pemilik:     "1",
	}
	t.Run("Success Add", func(t *testing.T) {
		data.On("Add", int(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.Add(id, inputData)
		assert.Nil(t, err)
		assert.Equal(t, inputData.Judul, res.Judul)
		assert.Equal(t, res.ID, resData.ID)
		data.AssertExpectations(t)
	})
	t.Run("Fail to add item", func(t *testing.T) {
		data.On("Add", int(1), inputData).Return(book.Core{}, errors.New("internal server error"))
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.Add(id, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	data := mocks.NewBookData(t)
	inputData := book.Core{
		ID:          uint(0),
		Judul:       "Avatar",
		TahunTerbit: 2003,
		Penulis:     "Aang",
	}
	resData := book.Core{
		ID:          uint(1),
		Judul:       "AvatarUy",
		TahunTerbit: 2003,
		Penulis:     "AangUy",
		Pemilik:     "1",
	}
	t.Run("Success Updating", func(t *testing.T) {
		data.On("Update", int(1), int(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		UserId := tokenIDUser.(*jwt.Token)
		UserId.Valid = true
		res, err := srv.Update(UserId, 1, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})
	t.Run("Update Fail", func(t *testing.T) {
		data.On("Update", int(1), int(1), inputData).Return(book.Core{}, errors.New("server error")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		UserId := tokenIDUser.(*jwt.Token)
		UserId.Valid = true
		res, err := srv.Update(UserId, 1, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
	t.Run("Book Not Found", func(t *testing.T) {
		data.On("Update", int(1), int(1), inputData).Return(book.Core{}, errors.New("book not found")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		UserId := tokenIDUser.(*jwt.Token)
		UserId.Valid = true
		res, err := srv.Update(UserId, 1, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestBookList(t *testing.T) {
	data := mocks.NewBookData(t)
	resData := []book.Core{}
	t.Run("Success Show Book List", func(t *testing.T) {
		data.On("BookList").Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.BookList()
		assert.Nil(t, err)
		assert.Equal(t, []book.Core{}, res)
		data.AssertExpectations(t)
	})
	t.Run("Fail to show book", func(t *testing.T) {
		data.On("BookList").Return([]book.Core{}, errors.New("no result")).Once()
		srv := New(data)
		res, err := srv.BookList()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "no result")
		assert.Equal(t, []book.Core{}, res)
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewBookData(t)
	t.Run("Success Deleting Book", func(t *testing.T) {
		data.On("Delete", int(1), int(1)).Return(nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		err := srv.Delete(id, 1)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})
	t.Run("Deleting Book Fail", func(t *testing.T) {
		data.On("Delete", int(1), int(1)).Return(errors.New("cannot delete book")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		err := srv.Delete(id, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "book")
		data.AssertExpectations(t)
	})
	t.Run("Wrong Book", func(t *testing.T) {
		data.On("Delete", int(1), int(1)).Return(errors.New("internal server error")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		err := srv.Delete(id, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}

func TestMyBook(t *testing.T) {
	data := mocks.NewBookData(t)
	t.Run("Success Show My Book", func(t *testing.T) {
		data.On("MyBook", int(1)).Return([]book.Core{}, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.MyBook(id)
		assert.Nil(t, err)
		assert.Equal(t, []book.Core{}, res)
		data.AssertExpectations(t)
	})
	t.Run("Error Show My Book", func(t *testing.T) {
		resData := []book.Core{}
		data.On("MyBook", int(1)).Return(resData, errors.New("book not found")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.MyBook(id)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, resData, res)
		data.AssertExpectations(t)
	})

}
