package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"hello word")
	})
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,"hello %s",name)
	})
	// param2 中还包含/
	router.GET("/user/:name/*param2", func(context *gin.Context) {
		name := context.Param("name")
		param2 := context.Param("param2")
		context.String(http.StatusOK,"hello %s+%s",name,param2)
	})

	v1 := router.Group("/v1")
	v1.GET("/read", func(context *gin.Context) {
		
	})


	router.Run(":8000")
}
