package main

import (
	"golearn/errhandling/filelistingserver/filelist"
	"log"
	"net/http"
	"os"
)

type addHandler func(writer http.ResponseWriter,
	request *http.Request) error

type userError interface {
	error
	Message() string
}

func errWrapper(
	handler addHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)

		if err != nil {
			log.Printf("Error handleing request: %s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(),
					http.StatusBadRequest)
				return
			}
			statusCode := http.StatusOK
			switch {
			// case os.IsNotExist(err):
			// 	http.Error(writer,
			// 		http.StatusText(http.StatusNotFound),
			// 		http.StatusNotFound)
			case os.IsNotExist(err):
				statusCode = http.StatusNotFound
			case os.IsPermission(err):
				statusCode = http.StatusForbidden
			default:
				statusCode = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(statusCode), statusCode)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelist.HandleFileList))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
