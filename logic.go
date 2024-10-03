package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"
)

// EJECUTA CALCULO ALL STATS CURVES
func calculateComputedValues(stats_compute Stats, rating int) Computed_Stats {

	return Computed_Stats{
		Health:                  runcurvehybrid(stats_compute.Strength, stats_compute.Vigor, "curvemaxhealth.json"),
		MemoryCapacity:          runcurve(stats_compute.Knowledge, "curvememorycapacity.json"),
		HealthRecovery:          runcurve(stats_compute.Vigor, "curvehealthrecovery.json") * 100,                                                  //calc display adjusment
		ActionSpeed:             runcurvehybrid(stats_compute.Agility, stats_compute.Dexterity, "curvesactionspeed.json") * 100,                   //calc display adjusment
		RegularInteractionSpeed: runcurvehybridx(stats_compute.Agility, stats_compute.Resourcefulness, "curveregularinteractionspeed.json") * 100, //calc display adjusment
		MoveSpeed:               runcurve(stats_compute.Agility, "curvemovespeed.json") + 300,                                                     //calc display adjusment
		MoveSpeedCalc:           runcurve(stats_compute.Agility, "curvemovespeed.json")/3 + 100,                                                   //calc display adjusment
		PhysicalPower:           runcurve(stats_compute.Strength, "curvephysicalpower.json") - 15,                                                 //calc display adjusment
		ManualDexterity:         runcurve(stats_compute.Dexterity, "curvemanualdexterity.json") * 100,                                             //calc display adjusment
		EquipSpeed:              runcurve(stats_compute.Dexterity, "curveequipspeed.json") * 100,                                                  //calc display adjusment
		BuffDuration:            runcurve(stats_compute.Will, "curvebuffduration.json") * 100,                                                     //calc display adjusment
		DebuffDuration:          runcurve(stats_compute.Will, "curvedebuffduration.json") * 100,                                                   //calc display adjusment
		MagicResistance:         runcurve(stats_compute.Will, "curvemagicresistance.json"),
		SpellRecovery:           runcurve(stats_compute.Will, "curvespellrecovery.json") * 100,
		SpellCastingSpeed:       runcurve(stats_compute.Will, "curvespellcastingspeed.json"),
		MagicalInteractionSpeed: runcurve(stats_compute.Will, "curvemagicalinteractionspeed.json"),
		Persuasiveness:          runcurve(stats_compute.Resourcefulness, "curvepersuasiveness.json"),
		CooldownReduction:       runcurve(stats_compute.Resourcefulness, "curvecooldownreduction.json") * 100,
		PhysicalDamageReduction: runcurve(rating, "curvephysicaldamagereduction.json"), // 0 defense value without equipement
	}
}

/*
func SelectItem_chest(itemselected string) Item_Armor {
	int_converted, _ := strconv.Atoi(itemselected)
	allItem_chest := []Item_Armor{Items.Doublet}
	return allItem_chest[int_converted]
}*/
/*
func Itemclass(class string, slot string) []Item_Armor {
	var result []Item_Armor
	for i := 0; i < len(Items.ItemsArmor[slot]); i++ {
		for _, c := range Items.ItemsArmor[slot][i].Classes {
			if c == class {
				result = append(result, Items.ItemsArmor[slot][i])
				break
			}
		}

	}
	return result
}
*/

func Itemclassdos(class string, a string) []Item_Armor {
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
			if c == class && a == Items.ItemsArmor[i].SlotType {
				result = append(result, Items.ItemsArmor[i])
				break
			}
		}
	}
	return result
}

// FUNCTION RETURN ITEM BY NAME STRING
func getFieldValue(fieldName string) Item_Armor {
	v := reflect.ValueOf(item)
	field := v.FieldByName(fieldName)
	if field.IsValid() && field.Kind() == reflect.Struct { // Check if the field exists and is a string
		return field.Interface().(Item_Armor) // Return the field value
	}
	return Item_Armor{}
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

/*
// FUNCION SUMA ITEM  & ATTRIBUTES
func sumStats(base Stats, add Item_Armor, key int) Stats {
	base.Strength -= add.BaseAttribute.Strength[key]
	base.Vigor -= add.BaseAttribute.Vigor[key]
	base.Agility -= add.BaseAttribute.Agility[key]
	base.Dexterity -= add.BaseAttribute.Dexterity[key]
	base.Will -= add.BaseAttribute.Will[key]
	base.Knowledge -= add.BaseAttribute.Knowledge[key]
	base.Resourcefulness -= add.BaseAttribute.Resourcefulness[key]
	return base
}*/

/*
func subStats(base Stats, add Item_Armor, key int) Stats {
	base.Strength -= add.BaseAttribute.Strength[key]
	base.Vigor -= add.BaseAttribute.Vigor[key]
	base.Agility -= add.BaseAttribute.Agility[key]
	base.Dexterity -= add.BaseAttribute.Dexterity[key]
	base.Will -= add.BaseAttribute.Will[key]
	base.Knowledge -= add.BaseAttribute.Knowledge[key]
	base.Resourcefulness -= add.BaseAttribute.Resourcefulness[key]
	return base
}*/

// CREA UN ITEM
func createItem(itemfile string) Item_Armor {
	var item Item_Armor // item to create
	readJSON(itemfile, &item)
	return item
}

func SetItemStats(baseStats Stats, item Item_Armor, key int) Stats {
	// Calculate total by summing base stats and item stats
	totalStats := Stats{
		Strength:        baseStats.Strength + item.BaseAttribute.Strength[key],
		Vigor:           baseStats.Vigor + item.BaseAttribute.Vigor[key],
		Agility:         baseStats.Agility + item.BaseAttribute.Agility[key],
		Dexterity:       baseStats.Dexterity + item.BaseAttribute.Dexterity[key],
		Will:            baseStats.Will + item.BaseAttribute.Will[key],
		Knowledge:       baseStats.Knowledge + item.BaseAttribute.Knowledge[key],
		Resourcefulness: baseStats.Resourcefulness + item.BaseAttribute.Resourcefulness[key],
	}
	// Return the total stats struct
	return totalStats
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

/*
func calculateCharacter(class string, itemHelmet string) Character {


	// Calculate new stats based on item


	// Calculate computed stats based on updated stats


	return Character{
		Stats:         updatedStats,
		Computed:      computedStats,
		SelectedClass: class,
		SelectedItem:  itemHelmet,
	}
}
*/
