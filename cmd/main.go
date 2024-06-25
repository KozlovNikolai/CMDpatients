package main

import (
	"net/http"

	"github.com/KozlovNikolai/CMDpatients/internal/config"
	"github.com/KozlovNikolai/CMDpatients/internal/server"
	"github.com/KozlovNikolai/CMDpatients/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustLoad()

	store.InitDB(cfg.StoragePath)
	defer store.CloseDB()
	router := gin.Default()

	router.POST("/patients", server.Create)
	router.GET("/patients/:id", server.Get)
	router.DELETE("/patients/:id", server.Delete)
	router.GET("/patients/list", server.GetList)

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimout,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
