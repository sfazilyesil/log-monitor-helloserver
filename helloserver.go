package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type HttpServer interface {
	Run() error
}

type helloHttpServer struct {
	adrr string
}

func NewHelloHttpServer(port string) *helloHttpServer {
	if len(port) == 0 {
		port = "80"
	}
	return &helloHttpServer{adrr: ":" + port}
}

func (h helloHttpServer) Run() (err error) {
	defer func() {
		log.Fatal("Error occured while starting the server: %+v \n", err)
	}()

	log.Println("Starting the hello http server...")

	counter := 0

	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		log.Println("GET /hello received the request.")

		counter++

		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		msg := `{"msg":"Merhaba ziyaretçi...", "count":` + strconv.Itoa(counter) + `}`
		io.WriteString(res, msg)

		log.Println("GET /hello send the response: " + msg)
	})
	err = http.ListenAndServe(h.adrr, nil)

	return
}
