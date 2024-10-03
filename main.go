package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	/*
		var inputclass string
		var inputitem string
		//var inputitem2 string

		fmt.Scanln(&inputclass)
		fmt.Scanln(&inputitem)

		class := SelectClass(inputclass)

		item := getFieldValue(inputitem)

		fmt.Println(class)
		fmt.Println(item.BaseAttribute)

		allitemfighter := Itemclass("Fighter", "Chest")


		classitem := SetItemStats(class, item, 1)

		fmt.Println(allitemfighter[0].ArmorRatings)

		fmt.Println(classitem)

		computedStats := calculateComputedValues(classitem)
		fmt.Println(computedStats)
	*/

	//test := getFieldValue("Armet")
	//fmt.Println(test)
	// default gin
	r := gin.Default()

	// Directorio Javascript & CSS
	r.Static("/static", "./static")

	// Directorio Templates html
	r.LoadHTMLGlob("templates/*")

	r.NoRoute(func(c *gin.Context) {
		c.File("")
	})

	// Function call routes
	setupRoutes(r)

	// Corre en puerto 8080
	r.Run()

}
