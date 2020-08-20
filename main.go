package main

import (
	"github.com/DowneyL/now/internal/models"
	"github.com/DowneyL/now/internal/routers"
	"github.com/DowneyL/now/pkg/configs"
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
