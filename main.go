package main

import (
	"errors" //Handles errors
	"github.com/gin-gonic/gin"
	"net/http" //Handles HTTP requests' status
)

// Classes are Struct in Go
type toDo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var toDos = []toDo{
	{Id: "1", Item: "Clean Room", Completed: false},
	{Id: "2", Item: "Clean Wardrobe", Completed: false},
	{Id: "3", Item: "Finish Project", Completed: false},
}

// UPDATE

func toggleToDoStatus(context *gin.Context) {
	id := context.Param("id")
	toDo, err := getToDoById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "ToDo not found"})
		return
	}

	toDo.Completed = !toDo.Completed

	context.IndentedJSON(http.StatusOK, toDo)
}

// GETS
func getToDos(context *gin.Context) { //gin.Context data about the upcoming HTTP request,
	// like JSON data and request header
	context.IndentedJSON(http.StatusOK, toDos)
}

func getToDo(context *gin.Context) {
	id := context.Param("id")
	toDo, err := getToDoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "ToDo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, toDo)
}
func getToDoById(id string) (*toDo, error) {
	for i, t := range toDos {
		if t.Id == id {
			return &toDos[i], nil
		}
	}
	return nil, errors.New("ToDo not found")
}

// CREATE
func addToDo(context *gin.Context) {
	var newToDo toDo
	if err := context.BindJSON(&newToDo); err != nil { //Ampersand actua como 'out' en C#, permite tirar el valor que da la funcion y asignarlo a una variable
		return
	}
	toDos = append(toDos, newToDo)
	context.IndentedJSON(http.StatusCreated, newToDo)
}

// SIMILAR TO PROGRAM.CS

func main() {
	// Creates Server
	router := gin.Default()

	//
	router.GET("/toDos", getToDos)
	router.GET("/toDos/:id", getToDo)
	router.PATCH("toDos/:id", toggleToDoStatus)
	router.POST("/toDos", addToDo)
	//Runs server
	router.Run("localhost:9090")
}
