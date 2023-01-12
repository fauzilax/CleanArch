package user

import "github.com/labstack/echo/v4"

//buat struct perjanjian apa saja yang ada dalam fitur user
type Core struct {
	ID       uint
	Name     string
	Email    string
	Address  string
	HP       string
	Password string
}

//buat kontrak fitur apa saja dalam user
type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

//buat kontrak apa saja parameter dan return Value yang dibutuhkan
type UserService interface {
	Register(newUser Core) (Core, error)
	Login(email, password string) (string, Core, error)
	Profile(tokenIDUser interface{}) (Core, error)
	Update(tokenIDUser interface{}, updateData Core) (Core, error)
	Delete(tokenIDUser interface{}) (error)
}

//buat kontrak apa saja parameter dan return value yang dibutuhkan untuk diquery di Database
type UserData interface {
	Register(newUser Core) (Core, error)
	Login(email string) (Core, error)
	Profile(id uint) (Core, error)
	Update(id int, updateData Core) (Core, error)
	Delete(userID int) (error)
}
