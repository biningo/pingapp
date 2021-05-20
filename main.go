package main

/**
*@Author lyer
*@Date 2020/12/5 22:40
*@Describe
**/
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping1", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 200,
			"data": context.Request.URL.Path,
		})
	})
	r.GET("/ping2", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 200,
			"data": context.Request.URL.Path,
		})
	})
	r.GET("/ping3", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 200,
			"data": context.Request.URL.Path,
		})
	})
	r.Run(":7777")
}
