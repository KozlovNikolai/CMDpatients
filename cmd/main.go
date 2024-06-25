package main

import (
	"gin-pg-crud/internal/server"
	"gin-pg-crud/internal/store"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	store.InitDB()
	defer store.CloseDB()
	router := gin.Default()

	router.POST("/employers", server.CreateEmployer)
	router.POST("/persons", server.CreatePerson)
	router.GET("/employers/:id", server.GetEmployer)
	router.GET("/persons/:id", server.GetPerson)
	router.PUT("/employers/:id", server.UpdateEmployer)
	router.DELETE("/employers/:id", server.DeleteEmployer)
	router.DELETE("/persons/:id", server.DeletePerson)
	router.GET("/persons/list", server.GetPersonList)
	router.GET("/employers/list", server.GetEmployerList)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  14 * time.Second,
		WriteTimeout: 14 * time.Second,
		IdleTimeout:  63 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
