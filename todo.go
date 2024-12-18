package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
type todo struct{
    Id int
    Title string
    DueDate string
    Completed bool
}
var todos[] todo
func init() {
    err := godotenv.Load()
    if err != nil {
    log.Fatal("Error loading .env file")
  }
}
func main(){
    
    r := gin.Default()
	r.GET("/getTodo", func(c *gin.Context) {
       
		c.JSON(200, gin.H{
			"message":todos,
		})
	})
    r.POST("/addTodo", func(c *gin.Context) {
        var body struct{
            Id int
            Title string
            DueDate string
            Completed bool
        }
        c.Bind(&body);
        userTodo :=todo{Id:body.Id,Title:body.Title,DueDate:body.DueDate,Completed:body.Completed}
        todos=append(todos,userTodo)
		c.JSON(200, gin.H{
			"message":"Todo added successfully",
		})
	})
    r.DELETE("/deleteTodo/:id", func(c *gin.Context) {
        id:=c.Param("id");
        fmt.Println(id);
        parsedId, err := strconv.Atoi(id) 
        if err !=nil{
           c.JSON(400,gin.H{
            "message":"input error",
           })
           return
        }
        for index, todo := range todos {
			if todo.Id == parsedId {
				todos = append(todos[:index], todos[index+1:]...)
				c.JSON(200, gin.H{"message": "Todo deleted successfully"})
				return
			}
		}
		c.JSON(200, gin.H{
			"message": "Todo deleted successfully",
		})
	})
    r.PUT("/updateTodo/:id", func(c *gin.Context) {
        var body struct{
            Title string
            DueDate string
            Completed bool
        }
        c.Bind(&body);
        id:=c.Param("id");
        parsedId, err := strconv.Atoi(id) 
        if err !=nil{
           c.JSON(400,gin.H{
            "message":"input error",
           })
           return
        }
        for index, todo := range todos {
			if todo.Id == parsedId {
				todos[index].Title=body.Title
                todos[index].DueDate=body.DueDate
                todos[index].Completed=body.Completed
				c.JSON(200, gin.H{"message": "Todo updated successfully"})
				return
			}
		}
		c.JSON(400, gin.H{
			"message": "Todo not found",
		})
	})
	r.Run()

}