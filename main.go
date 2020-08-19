package main

import (
	"github.com/DowneyL/now/models"
	"github.com/DowneyL/now/packages/configs"
	"github.com/DowneyL/now/routers"
	"log"
	"net/http"
)

func main() {
	config := configs.New()
	models.SetUp()

	r := routers.InitRouter()
	server := &http.Server{
		Addr:           config.GetHttpPort(),
		Handler:        r,
		ReadTimeout:    config.GetHttpReadTimeout(),
		WriteTimeout:   config.GetHttpWriteTimeout(),
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
}
