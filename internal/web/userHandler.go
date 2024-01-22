package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	userModel "github.com/sri2103/htmx_go/internal/userAuth/model"
)

func (w *webHandler) LoginPage(c echo.Context) error {
	data := make(map[string]interface{})
	data["PageTitle"] = "Login"
	return c.Render(http.StatusOK, "pages/login", data)
}
func (w *webHandler) Register(c echo.Context) error {
	data := make(map[string]interface{})
	data["PageTitle"] = "Register"
	return c.Render(http.StatusOK, "pages/register", data)
}

func (w *webHandler) PostLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := w.userService.Login(c.Request().Context(), email, password)
	if err != nil {
		return err
	}
	w.session.Put(c.Request().Context(), "user", email)
	w.session.Put(c.Request().Context(), "userId", user.ID)
	return c.Redirect(http.StatusSeeOther, "/web/home")

}

func (w *webHandler) PostRegister(c echo.Context) error {
	var user userModel.User
	email := c.FormValue("email")
	pass := c.FormValue("password")
	name := c.FormValue("name")
	user.Name = name
	user.Password = pass
	user.Email = email
	_, err := w.userService.Register(c.Request().Context(), &user)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/web/login")

}

func (w *webHandler) LogOut(c echo.Context) error {
	ctx := c.Request().Context()
	_ = w.session.Destroy(ctx)
	_ = w.session.RenewToken(ctx)
	return c.Redirect(http.StatusSeeOther, "/web/login")
}
