package main

import "github.com/gin-gonic/gin"

func Dashboard(r *gin.RouterGroup) {
	stats := r.Group("/stats")
	{
		stats.GET("", func(c *gin.Context) {

		})
	}
}
