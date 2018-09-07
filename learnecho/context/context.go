package context

import "github.com/labstack/echo"

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo(foo string) {
	println(foo)
}


