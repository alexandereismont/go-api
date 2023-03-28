package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

var customers = []Customer{
	{Firstname: "Alex", Lastname: "SomeLastname", Age: 20},
	{Firstname: "Alex", Lastname: "Othername", Age: 78},
}

func main() {

	for _, customer := range customers {
		fmt.Println(customer)
	}

	router := gin.Default()
	router.GET("/customers", getCustomers)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "HEALTHY")
	})
	router.Run("localhost:8001")
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}
