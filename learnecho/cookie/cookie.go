package cookie

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
	"fmt"
)

func WriteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "yizhao"
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.SetCookie(cookie)
	return c.String(200, "Write a cookie")
}

func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}

	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(200, "Read a cookie")
}

func ReadAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}

	return c.String(200, "Read all cookies")
}
