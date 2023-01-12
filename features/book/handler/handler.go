package handler

import (
	"CleanArch/features/book"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookHandle struct {
	srv book.BookService
}

func New(bs book.BookService) book.BookHandler {
	return &bookHandle{
		srv: bs,
	}
}

// Add implements book.BookHandler
func (bh *bookHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := AddBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		cnv := ConvToCore(input)

		res, err := bh.srv.Add(c.Get("user"), *cnv) // c.Get("user") user <- table "user" yang diambil adalah ID nya
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "creating book fail",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    AddResponse(res),
			"message": "book has successfull added",
		})
	}
}

// Update implements book.BookHandler
func (bh *bookHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamBookID := c.Param("id")
		bookID, _ := strconv.Atoi(ParamBookID)
		input := UpdateBookRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		cnv := ConvToCore(input)
		res, err := bh.srv.Update(c.Get("user"), bookID, *cnv) // c.Get("user") user <- table "user" yang diambil adalah ID nya
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "update book fail",
			})
		}
		if res.Judul == "" {
			res.Judul = "tidak ada perubahan"
		}
		if res.Penulis == "" {
			res.Penulis = "tidak ada perubahan"
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    UpdateBookResponse(res),
			"message": "update successfull",
		})
	}
}

// BookList implements book.BookHandler
func (bh *bookHandle) BookList() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bh.srv.BookList() // c.Get("user") user <- table "user" yang diambil adalah ID nya
		if err != nil {
			log.Println("no book found ", err.Error())
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "no result",
			})
		}
		// log.Println(res)
		result := []BookList{}
		for i := 0; i < len(res); i++ {
			result = append(result, BookListResponse(res[i]))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "show all book list succesfull",
		})
	}
}

// MyBook implements book.BookHandler
func (bh *bookHandle) MyBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("user") //mengambil id dari token

		res, err := bh.srv.MyBook(userID)
		if err != nil {
			log.Println("no book found ", err.Error())
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "no result",
			})
		}
		result := []AddBookReponse{}
		for i := 0; i < len(res); i++ {
			result = append(result, MyBookResponse(res[i]))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "show all book list succesfull",
		})
	}
}

// Delete implements book.BookHandler
func (bh *bookHandle) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ParamBookID := c.Param("id")
		bookID, _ := strconv.Atoi(ParamBookID)
		err := bh.srv.Delete(c.Get("user"), bookID)
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "delete fail",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Delete successfull",
		})
	}
}
