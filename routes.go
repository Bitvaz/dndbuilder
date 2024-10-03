package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func charBuilderHandler(c *gin.Context) {

	selected_Class := c.Param("classSelection")
	newStats := SelectClass(selected_Class)
	computedStats := calculateComputedValues(newStats, 0)

	//newStats := SelectClass(selected_Class)

	// character := calculateCharacter("default", "")

	c.HTML(http.StatusOK, "charbuilder.html", gin.H{
		"stats":         newStats,
		"computedstats": computedStats,
	},
	)
}

func updateStatsHandler(c *gin.Context) {

	selected_Class := c.Param("classSelection")
	selectedItem_helmet := c.Query("itemhelmet")
	selectedRarity := c.Query("rarityselect")
	selectedItemRating := c.Query("armorrating")

	classStatSelect := SelectClass(selected_Class)
	itemHelmetSelect := getFieldValue(selectedItem_helmet)
	raritySelect := StringtoInt(selectedRarity)
	itemRatingSelect := StringtoInt(selectedItemRating)

	printClass := InttoClass(selected_Class)
	//fmt.Println(selectedItem_helmet)

	helmetList := Itemclassdos(selected_Class, "Head")
	chestList := Itemclassdos(selected_Class, "Chest")
	glovesList := Itemclassdos(selected_Class, "Hands")
	pantsList := Itemclassdos(selected_Class, "Legs")
	bootsList := Itemclassdos(selected_Class, "Foot")
	cloakList := Itemclassdos(selected_Class, "Back")

	updatedStats := SetItemStats(classStatSelect, itemHelmetSelect, raritySelect)

	//fmt.Println(itemHelmetSelected.ArmorRatings[2])

	computedStats := calculateComputedValues(updatedStats, itemRatingSelect) // rating variable armor rating

	// character := calculateCharacter(selected_Class, selectedItem_helmet) - character LATER
	//c.ShouldBind(&newStats) /// USED FOR MODIFY EVERY SINGLE STATS --> TESTING TOTAL VALUES

	c.HTML(http.StatusOK, "charbuilder.html", gin.H{

		"stats":          updatedStats,
		"computedstats":  computedStats,
		"helmetlist":     helmetList,
		"chestlist":      chestList,
		"gloveslist":     glovesList,
		"pantslist":      pantsList,
		"bootslist":      bootsList,
		"cloaklist":      cloakList,
		"test":           selectedItem_helmet,
		"testdo":         selectedRarity,
		"printclass":     printClass,
		"itemratingsone": itemHelmetSelect.ArmorRatings[raritySelect],
		//"itemRatingstwo":

		//"printout":      selectedItem_chest,
	})

}

func setupRoutes(r *gin.Engine) {
	//r.GET("/charbuilder/", charBuilderHandler)

	r.GET("/charbuilder/", charBuilderHandler)
	r.GET("/charbuilder/:classSelection", updateStatsHandler)

}
