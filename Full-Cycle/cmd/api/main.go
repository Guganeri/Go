package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Post("/order", Hello)
	//r.Get("/order", Hello)
	//http.HandleFunc("/", Hello)
	//http.ListenAndServe(":8888", r)

	e := echo.New()
	e.GET("/test", nil)
	e.Logger.Fatal(e.Start(":9999"))

}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
