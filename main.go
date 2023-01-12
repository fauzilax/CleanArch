package main

import (
	"CleanArch/config"
	bkData "CleanArch/features/book/data"
	bkHdl "CleanArch/features/book/handler"
	bkSrv "CleanArch/features/book/services"
	usrData "CleanArch/features/user/data"
	usrHdl "CleanArch/features/user/handler"
	usrSrv "CleanArch/features/user/services"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	// // panggil fungsi Migrate untuk buat table baru di database
	config.Migrate(db)

	userData := usrData.New(db)
	userSrv := usrSrv.New(userData)
	userHdl := usrHdl.New(userSrv)

	bookData := bkData.New(db)
	bookSrv := bkSrv.New(bookData)
	bookHdl := bkHdl.New(bookSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.PATCH("/users", userHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))

	e.POST("/addbook", bookHdl.Add(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/books", bookHdl.MyBook(), middleware.JWT([]byte(config.JWTKey)))
	e.PATCH("/books/:id", bookHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/booklist", bookHdl.BookList())
	e.DELETE("/books/:id", bookHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}

}
