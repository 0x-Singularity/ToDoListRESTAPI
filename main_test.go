package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "encoding/json"
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

func TestGetTodos(t *testing.T) {
    router := setupRouter()

    req, err := http.NewRequest("GET", "/todos", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }

    // Decode the HTTP response
    var got []todo
    err = json.Unmarshal(w.Body.Bytes(), &got)
    if err != nil {
        t.Fatalf("Couldn't unmarshal response body: %v\n", err)
    }

    // Compare the response with the expected result
    want := todos // Assuming 'todos' is the expected slice of todos
    if !compareTodos(got, want) {
        t.Errorf("Handler returned unexpected body: got %v want %v", got, want)
    }
}

// Helper function to compare two slices of todos
func compareTodos(a, b []todo) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i].ID != b[i].ID || a[i].Item != b[i].Item || a[i].Completed != b[i].Completed {
            return false
        }
    }
    return true
}
