package main

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/labstack/echo/middleware"
	"os"
	"io"
	"golearn/learnecho/context"
	"golearn/learnecho/cookie"
	"gopkg.in/go-playground/validator.v9"
	//"encoding/xml"
	"golang.org/x/crypto/acme/autocert"
)

type (
	User struct {
		Name string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
		Email string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	//return cv.Validate(i)
	return cv.validator.Struct(i) //github.com/go-playground/validator
}


func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello, yzhao!")
}

func getUser(c echo.Context) error {

	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}


//curl -F "name=Joe Smith" -F "avatar=@/home/yzhao_sherry/Pictures/1.png" http://localhost:1323/copy
func copyImage(c echo.Context) error {

	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")

	if err != nil {
		return err
	}

	//Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}

	//copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}


func jsonXml(c echo.Context) error {
	u := new(User)
	//u := User {
	//	Name:"yzhao",
	//	Email:"2993614148@qq.com",
	//}
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
	//return c.XML(http.StatusCreated, u)
}

func jsonPretty(c echo.Context) error {
	u := &User{
		Name: "Zhaoyihui",
		Email: "2993614148@qq.com",
	}

	//return c.JSONPretty(http.StatusOK, u, "")
	//return c.JSON(200, u)

	return c.XMLPretty(http.StatusOK, u, "  ")

	//c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)
	//c.Response().WriteHeader(http.StatusOK)
	//return xml.NewEncoder(c.Response()).Encode(u)
}

func redirect(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "http://kenzhaoyihui.top")
}

func groupAuth(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	return c.String(200, "username: " + username + "\npassword:" + password)
}

func main() {
	e := echo.New()
	g := e.Group("/admin")
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	//Validator
	e.Validator = &CustomValidator{validator:validator.New()}

	e.HideBanner = true

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.CustomContext{c}
			return h(cc)
		}
	})

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "yzhao" && password == "redhat" {
			return true, nil
		}
		return false, nil
	}))

	//Routes
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.POST("/save", save)
	e.POST("/copy", copyImage)
	e.POST("/users", jsonXml)
	e.GET("/foo", func(c echo.Context) error {
		if cc, ok := c.(*context.CustomContext); ok {
			cc.Foo("yzhao")
			return cc.String(http.StatusOK, "OK")
		}
		return nil
	})

	e.POST("/usersV", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		if err = c.Validate(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	e.GET("/userJsonP", jsonPretty)


	e.GET("/createCookie", cookie.WriteCookie)
	e.GET("/readCookie", cookie.ReadCookie)
	e.GET("/readAllCookie", cookie.ReadAllCookies)

	e.GET("/redirect", redirect)

	e.GET("/admin/1", groupAuth)

	e.GET("/autotls", func(c echo.Context) error {
		return c.HTML(200, `
<h1>Welcome to echo!</h1>
<h3>TLS certificates automatically installed from let's Encrypt:)</h3>'`)
	})

	e.Logger.Fatal(e.Start(":1323"))
	//e.Logger.Fatal(e.StartAutoTLS(":443"))
}
