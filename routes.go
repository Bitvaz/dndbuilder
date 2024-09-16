package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func charBuilderHandler(c *gin.Context) {

	computedStats := calculateComputedValues()

	c.HTML(http.StatusOK, "charbuilder.html", gin.H{
		"stats":         characterStats,
		"computedstats": computedStats,
	},
	)
}

func updateStatsHandler(c *gin.Context) {
	var newStats Stats
	selectedClass := c.PostForm("classSelection")
	newStats = SelectClass(selectedClass)
	if err := c.ShouldBind(&newStats); err == nil {

		characterStats = newStats
		computedStats := calculateComputedValues()

		c.HTML(http.StatusOK, "charbuilder.html", gin.H{
			"stats":         characterStats,
			"computedstats": computedStats,
		})

	} else {
		// Handle error if form binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func setupRoutes(r *gin.Engine) {
	r.GET("/charbuilder", charBuilderHandler)
	r.POST("/update", updateStatsHandler)
}
