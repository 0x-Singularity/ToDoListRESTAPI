package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "strings"
)

func setupRouter() *gin.Engine {
    // Set up the router for testing
    router := gin.Default()
    router.GET("/todos", getTodos)
    router.GET("/todos/:id", getTodo)
    router.PATCH("/todos/:id", toggleTodoStatus)
    router.POST("/todos", addTodo)
    return router
}
