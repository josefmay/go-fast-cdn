package main

import (
	"log"
	"os"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kevinanielsen/go-fast-cdn/src/database"
	ini "github.com/kevinanielsen/go-fast-cdn/src/initializers"
	"github.com/kevinanielsen/go-fast-cdn/src/router"
	"github.com/kevinanielsen/go-fast-cdn/src/util"
)

func init() {
	util.LoadExPath()
	gin.SetMode("release")
	ini.LoadEnvVariables(true)
	ini.CreateFolders()
	database.ConnectToDB()
}

func main() {
	log.Printf("[info] strating http server, listening on port %v", os.Getenv("PORT"))
	r := router.Router()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf(port)
	s := &http.Server{
		Addr:	port,
		Handler:	r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("[error] server err: %v", err)
	}
}
