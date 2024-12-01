package main

import (
	"encoding/json"
	"os"
	"strconv"
)

// EJECUTA CALCULO ALL STATS CURVES
func CalculateComputedValues(stats_compute Stats, rating_armor int, rating_speed int) Computed_Stats {

	return Computed_Stats{
		Health:                  runcurvehybrid(stats_compute.Strength, stats_compute.Vigor, "./calc/curvemaxhealth.json"),
		MemoryCapacity:          runcurve(stats_compute.Knowledge, "./calc/curvememorycapacity.json"),
		HealthRecovery:          runcurve(stats_compute.Vigor, "./calc/curvehealthrecovery.json") * 100,                                                  //calc display adjusment
		ActionSpeed:             runcurvehybrid(stats_compute.Agility, stats_compute.Dexterity, "./calc/curvesactionspeed.json") * 100,                   //calc display adjusment
		RegularInteractionSpeed: runcurvehybridx(stats_compute.Agility, stats_compute.Resourcefulness, "./calc/curveregularinteractionspeed.json") * 100, //calc display adjusment
		MoveSpeed:               runcurve(stats_compute.Agility, "./calc/curvemovespeed.json") + 300 + float64(rating_speed),                             //calc display adjusment
		MoveSpeedCalc:           runcurve(stats_compute.Agility, "./calc/curvemovespeed.json")/3 + 100,                                                   //calc display adjusment
		PhysicalPower:           runcurve(stats_compute.Strength, "./calc/curvephysicalpower.json") - 15,                                                 //calc display adjusment
		ManualDexterity:         runcurve(stats_compute.Dexterity, "./calc/curvemanualdexterity.json") * 100,                                             //calc display adjusment
		EquipSpeed:              runcurve(stats_compute.Dexterity, "./calc/curveequipspeed.json") * 100,                                                  //calc display adjusment
		BuffDuration:            runcurve(stats_compute.Will, "./calc/curvebuffduration.json") * 100,                                                     //calc display adjusment
		DebuffDuration:          runcurve(stats_compute.Will, "./calc/curvedebuffduration.json") * 100,                                                   //calc display adjusment
		MagicResistance:         runcurve(stats_compute.Will, "./calc/curvemagicresistance.json"),
		SpellRecovery:           runcurve(stats_compute.Will, "./calc/curvespellrecovery.json") * 100,
		SpellCastingSpeed:       runcurve(stats_compute.Will, "./calc/curvespellcastingspeed.json"),
		MagicalInteractionSpeed: runcurve(stats_compute.Will, "./calc/curvemagicalinteractionspeed.json"),
		Persuasiveness:          runcurve(stats_compute.Resourcefulness, "./calc/curvepersuasiveness.json"),
		CooldownReduction:       runcurve(stats_compute.Resourcefulness, "./calc/curvecooldownreduction.json") * 100,
		PhysicalDamageReduction: runcurve(rating_armor, "./calc/curvephysicaldamagereduction.json") * 100, // 0 defense value without equipement
	}
}

func ItemsBySlotType(class string, slot string) []Item_Armor {
	switch class {
	case "1":
		class = "Fighter"
	case "2":
		class = "Barbarian"
	case "3":
		class = "Rogue"
	case "4":
		class = "Wizard"
	case "5":
		class = "Cleric"
	case "6":
		class = "Warlock"
	case "7":
		class = "Bard"
	case "8":
		class = "Druid"
	case "9":
		class = "Ranger"
	default:

	}
	var result []Item_Armor
	for i := 0; i < len(Items.ItemsArmor); i++ {
		for _, c := range Items.ItemsArmor[i].Classes {
			if c == class && slot == Items.ItemsArmor[i].SlotType {
				result = append(result, Items.ItemsArmor[i])
				break
			}
		}

	}
	return result
}

func ItemsByName(name string) Item_Armor {
	var result Item_Armor
	for i := 0; i < len(Items.ItemsArmor); i++ {
		if name == Items.ItemsArmor[i].Name {
			result = Items.ItemsArmor[i]
			return result
		}
	}
	return result
}

// SELECT CLASS AND CONDITIONAL IF NO CLASS SELECTED
func SelectClass(classselected string) Stats {
	allClases_array := [10]Stats{characterStats, classFighter, classBarbarian, classRogue, classWizard, classCleric, classWarlock, classBard, classRanger, classDruid}
	int_converted, _ := strconv.Atoi(classselected)
	for i := 0; i <= len(allClases_array); i++ {
		if int_converted == i {
			return allClases_array[i]
		}
	}
	return characterStats
}

// LOAD ITEM FROM JSON
func readJSON(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}

// CREA UN ITEM
func CreateItem(itemfile string) Item_Armor {
	var item Item_Armor // item to create
	readJSON(itemfile, &item)
	return item
}

func SetItemStats(baseStats Stats, item []Item_Armor, rarity []int) Stats {
	// Calculate total by summing base stats and item stats
	for i := 0; i < len(item) && i < len(rarity); i++ {
		baseStats = Stats{
			Strength:        baseStats.Strength + item[i].BaseAttribute.Strength[rarity[i]],
			Vigor:           baseStats.Vigor + item[i].BaseAttribute.Vigor[rarity[i]],
			Agility:         baseStats.Agility + item[i].BaseAttribute.Agility[rarity[i]],
			Dexterity:       baseStats.Dexterity + item[i].BaseAttribute.Dexterity[rarity[i]],
			Will:            baseStats.Will + item[i].BaseAttribute.Will[rarity[i]],
			Knowledge:       baseStats.Knowledge + item[i].BaseAttribute.Knowledge[rarity[i]],
			Resourcefulness: baseStats.Resourcefulness + item[i].BaseAttribute.Resourcefulness[rarity[i]],
		}
	}
	// Return the total stats struct
	return baseStats
}

func InttoClass(convert string) string {
	switch convert {
	case "1":
		convert = "Fighter"
	case "2":
		convert = "Barbarian"
	case "3":
		convert = "Rogue"
	case "4":
		convert = "Wizard"
	case "5":
		convert = "Cleric"
	case "6":
		convert = "Warlock"
	case "7":
		convert = "Bard"
	case "8":
		convert = "Druid"
	case "9":
		convert = "Ranger"
	default:
	}
	return convert
}

func StringtoInt(convert string) int {
	intkey, _ := strconv.Atoi(convert)
	return intkey
}

func RatingCalc(ratings []int) int {
	result := 0
	for i := 0; i < len(ratings); i++ {
		result = result + ratings[i]
	}
	return result
}

func SpeedCalc(items []Item_Armor, rarity []int) int {
	speedrating := 0
	for i := 0; i < len(items); i++ {
		speedrating += items[i].MoveSpeed[1]
		if items[i].SlotType == "Foot" {
			speedrating += items[i].MoveSpeed[rarity[i]]
		}
	}
	return speedrating
}

/*
func FixedCalc(items []Item_Armor )float64 {
	result := 0
	for i := 0; i < len(items); i++ {
		result += int(items[i].ProjectileReduction)
	}

}*/
