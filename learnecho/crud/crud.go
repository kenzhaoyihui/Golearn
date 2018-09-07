package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/labstack/echo/middleware"
	"fmt"
)

type user struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var (
	users = map[int]*user{}
	seq = 1
)

func createUser(c echo.Context) error {
	u := &user{
		ID:seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(200, users[id])
}


func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(200, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	//return c.NoContent(http.StatusNoContent)
	return c.JSON(200, users)
}

func getAllUser(c echo.Context) error {
	return c.JSON(200, users)

}


func main() {
	e := echo.New()

	//MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Pre(middleware.AddTrailingSlash())   default automatically to add the "/" after URL, http://localhost/admin -->
	//http://localhost/admin/

	//Routes
	e.GET("/users/:id", getUser)
	e.POST("/users/", createUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.GET("/getall", getAllUser)

	e.GET("/request", func(c echo.Context) error {
		req := c.Request()
		format := `
			<code>
				Protocol: %s<br>
				Host: %s<br>
				Remote Address: %s<br>
				Method: %s<br>
				Path: %s<br>
			</code>`
			return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr,
				req.Method, req.URL.Path))
	})

	//e.Logger.Fatal(e.Start(":1323"))
	//e.StartTLS(":1323", "learnecho/gentls/cert.pem", "learnecho/gentls/key.pem")
	e.Logger.Fatal(e.StartTLS(":1323", "learnecho/gentls/cert.pem", "learnecho/gentls/key.pem"))
}
