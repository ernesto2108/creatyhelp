package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type MyServer struct {
	server *http.Server
}
func NewServer(r *gin.Engine) *MyServer{

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &MyServer{s}
}

func (s MyServer) Run() {
	log.Fatal(s.server.ListenAndServe())
}