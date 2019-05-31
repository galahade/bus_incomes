package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	c "github.com/galahade/bus_incomes/handler"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

var env string
var port int

func main() {
	getParams()
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("./log/console.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()
	router.Use(ginglog.Logger(3*time.Second), gin.Logger(), gin.Recovery())

	//router config for rout request
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "welcome to my site.",
		})
	})
	router.GET("/incomes/:year/:month", c.GetMonthLinesIncomes)
	router.DELETE("/incomes/:year/:month/:line", c.DeleteLineMonthIncome)

	router.POST("/incomes", c.AddLineMonthIncome)
	router.GET("/monthIncomesCompare/:year/:month", c.GetSortedTwoMonthComparedIncomesByDate)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}

func getParams() {
	//flag.StringVar(&env, "env", "", "application enviroment")
	flag.IntVar(&port, "p", 8000, "application port number")
	flag.Parse()
}
