package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func charBuilderHandler(c *gin.Context) {

	selected_Class := c.Param("classSelection")
	newStats := SelectClass(selected_Class)
	computedStats := CalculateComputedValues(newStats, 0, 0)

	c.HTML(http.StatusOK, "charbuilder.html", gin.H{
		"stats":         newStats,
		"computedstats": computedStats,
	},
	)
}

func updateStatsHandler(c *gin.Context) {

	classStatSelect := SelectClass(c.Param("classSelection"))
	printClass := InttoClass(c.Param("classSelection"))

	itemsSelected := []Item_Armor{
		ItemsByName(c.Query("itemhelmet")),
		ItemsByName(c.Query("itemchest")),
		ItemsByName(c.Query("itemgloves")),
		ItemsByName(c.Query("itempants")),
		ItemsByName(c.Query("itemboots")),
		ItemsByName(c.Query("itemcloak")),
	}

	raritySelected := []int{
		StringtoInt(c.Query("rarityselect_helmet")),
		StringtoInt(c.Query("rarityselect_chest")),
		StringtoInt(c.Query("rarityselect_gloves")),
		StringtoInt(c.Query("rarityselect_pants")),
		StringtoInt(c.Query("rarityselect_boots")),
		StringtoInt(c.Query("rarityselect_cloak")),
	}

	ratingSelected := []int{
		StringtoInt(c.Query("armorrating_helmet")),
		StringtoInt(c.Query("armorrating_chest")),
		StringtoInt(c.Query("armorrating_gloves")),
		StringtoInt(c.Query("armorrating_pants")),
		StringtoInt(c.Query("armorrating_boots")),
		StringtoInt(c.Query("armorrating_cloak")),
	}

	ratingListHelmet := ItemsByName(c.Query("itemhelmet")).ArmorRatings[StringtoInt(c.Query("rarityselect_helmet"))]
	ratingListChest := ItemsByName(c.Query("itemchest")).ArmorRatings[StringtoInt(c.Query("rarityselect_chest"))]
	ratingListGloves := ItemsByName(c.Query("itemgloves")).ArmorRatings[StringtoInt(c.Query("rarityselect_gloves"))]
	ratingListPants := ItemsByName(c.Query("itempants")).ArmorRatings[StringtoInt(c.Query("rarityselect_pants"))]
	ratingListBoots := ItemsByName(c.Query("itemboots")).ArmorRatings[StringtoInt(c.Query("rarityselect_boots"))]
	ratingListCloak := ItemsByName(c.Query("itemcloak")).ArmorRatings[StringtoInt(c.Query("rarityselect_cloak"))]

	totalrating := RatingCalc(ratingSelected)
	totalspeed := SpeedCalc(itemsSelected, raritySelected)

	helmetList := ItemsBySlotType(c.Param("classSelection"), "Head")
	chestList := ItemsBySlotType(c.Param("classSelection"), "Chest")
	glovesList := ItemsBySlotType(c.Param("classSelection"), "Hands")
	pantsList := ItemsBySlotType(c.Param("classSelection"), "Legs")
	bootsList := ItemsBySlotType(c.Param("classSelection"), "Foot")
	cloakList := ItemsBySlotType(c.Param("classSelection"), "Back")

	updatedStats := SetItemStats(classStatSelect, itemsSelected, raritySelected)

	computedStats := CalculateComputedValues(updatedStats, totalrating, totalspeed) // rating variable armor rating

	//c.ShouldBind(&newStats) /// USED FOR MODIFY EVERY SINGLE STATS --> TESTING TOTAL VALUES

	c.HTML(http.StatusOK, "charbuilder.html", gin.H{

		"stats":             updatedStats,
		"computedstats":     computedStats,
		"helmetlist":        helmetList,
		"chestlist":         chestList,
		"gloveslist":        glovesList,
		"pantslist":         pantsList,
		"bootslist":         bootsList,
		"cloaklist":         cloakList,
		"printclass":        printClass,
		"helmet_ratinglist": ratingListHelmet,
		"chest_ratinglist":  ratingListChest,
		"gloves_ratinglist": ratingListGloves,
		"pants_ratinglist":  ratingListPants,
		"boots_ratinglist":  ratingListBoots,
		"cloak_ratinglist":  ratingListCloak,
		"test":              totalspeed,
	})

}

func setupRoutes(r *gin.Engine) {

	r.GET("/charbuilder/", charBuilderHandler)
	r.GET("/charbuilder/:classSelection", updateStatsHandler)

}
