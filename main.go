package main

import(
  "net/http"
  "github.com/gin-gonic/gin"
  "errors"
)

type todo struct {
  ID string `json:"id"`
  Item string `json:"item"`
  Completed bool `json:"completed"`
}

var todos = []todo{
  {ID: "1", Item: "Clean Room", Completed: false},
  {ID: "2", Item: "Read Book", Completed: false},
  {ID: "3", Item: "Record Video", Completed: false},
}

//converts the todo data structure into JSON
func getTodos(context *gin.Context) {
  context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
  var newTodo todo
  
  // This is going to take whatever JSON is inside request body and bind to new todo type. If a todo is passed and does not have the json ID type it will throw an error.
  if err := context.BindJSON(&newTodo); err != nil {
    return
  }

  todos = append(todos, newTodo)

  context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(id string)(*todo, error) {
  for i, t := range todos{
    if t.ID == id{
      return &todos[i], nil 
    }
  }
  return nil, errors.New("todo not found")
}

// Extracts the path parameter from the URL
func getTodo(context *gin.Context) {
  id := context.Param("id")
  todo, err := getTodoByID(id)
  
  if err != nil {
    context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
    return
  }

  context.IndentedJSON(http.StatusOK, todo)
}

func main () {
  router := gin.Default()
  router.GET("/todos", getTodos)
  router.GET("/todos/:id", getTodo)
  router.POST("/todos", addTodo)
  router.Run("localhost:9090")
}

