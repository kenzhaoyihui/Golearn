package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
	"fmt"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			//Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			//Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}


//use this website to test: http://www.websocket.org/echo.html

/**
1. fill in the location: ws://127.0.0.1:1323/ws
2. then click the connect, if connected, the console log will show that "CONNECTED"
3. then send some message

from the website console:
```
SENT: asdfasdf

RECEIVED: Hello, Client!

SENT: I am yihui

RECEIVED: Hello, Client!
```

from the server terminal:
```
⇨ http server started on [::]:1323
Rock it with HTML5 WebSocket
Rock it with HTML5 WebSocket
Rock it with HTML5 WebSocket
Rock it with HTML5 WebSocket
Rock it with HTML5 WebSocket
Rock it with HTML5 WebSocket
Rock it with HTML5 WebSocket
Rock it with HTML5 WebSocket
asdfasdf
I am yihui
```
 */
