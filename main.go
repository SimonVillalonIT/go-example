package main

import (
    "net/http"
    "github.com/SimonVillalonIT/example-go/controllers"
    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type person struct {
    ID     string  `json:"id"`
    Name string  `json:"title"`
    Age string  `json:"artist"`
}

// albums slice to seed record album data.
var persons = []person{
    {ID: "1", Name: "Simon", Age: "18" },
    {ID: "2", Name: "Enzo", Age: "18" },
    {ID: "3", Name: "Agustin", Age: "18" },
}

func main() {
    router := gin.Default()
    router.GET("/persons", getPersons)
    router.GET("/persons/:id", getPersonByID)
    router.POST("/persons", postPersons)

    router.Run("localhost:8080")
}


