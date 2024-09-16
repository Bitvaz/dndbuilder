package main

// EJECUTA CALCULO STATS
func calculateComputedValues() ComputedStats {
	return ComputedStats{
		Health:                  runcurvehybrid(characterStats.Strength, characterStats.Vigor, "curvemaxhealth.json"),
		MemoryCapacity:          runcurve(characterStats.Knowledge, "curvememorycapacity.json"),
		HealthRecovery:          runcurve(characterStats.Vigor, "curvehealthrecovery.json") * 100,                                                   //calc display adjusment
		ActionSpeed:             runcurvehybrid(characterStats.Agility, characterStats.Dexterity, "curvesactionspeed.json") * 100,                   //calc display adjusment
		RegularInteractionSpeed: runcurvehybridx(characterStats.Agility, characterStats.Resourcefulness, "curveregularinteractionspeed.json") * 100, //calc display adjusment
		MoveSpeed:               runcurve(characterStats.Agility, "curvemovespeed.json") + 300,                                                      //calc display adjusment
		MoveSpeedCalc:           runcurve(characterStats.Agility, "curvemovespeed.json")/3 + 100,                                                    //calc display adjusment
		PhysicalPower:           runcurve(characterStats.Strength, "curvephysicalpower.json") - 15,                                                  //calc display adjusment
		ManualDexterity:         runcurve(characterStats.Dexterity, "curvemanualdexterity.json") * 100,                                              //calc display adjusment
		EquipSpeed:              runcurve(characterStats.Dexterity, "curveequipspeed.json") * 100,                                                   //calc display adjusment
		BuffDuration:            runcurve(characterStats.Will, "curvebuffduration.json") * 100,                                                      //calc display adjusment
		DebuffDuration:          runcurve(characterStats.Will, "curvedebuffduration.json") * 100,                                                    //calc display adjusment
		MagicResistance:         runcurve(characterStats.Will, "curvemagicresistance.json"),
		SpellRecovery:           runcurve(characterStats.Will, "curvespellrecovery.json") * 100,
		SpellCastingSpeed:       runcurve(characterStats.Will, "curvespellcasting.json"),
		MagicalInteractionSpeed: runcurve(characterStats.Will, "curvemagicalinteractionspeed.json"),
		Persuasiveness:          runcurve(characterStats.Resourcefulness, "curvepersuasiveness.json"),
		CooldownReduction:       runcurve(characterStats.Resourcefulness, "curvecooldownreduction.json") * 100,
	}
}

// SELECT CLASS LOGIC
func SelectClass(classselected string) Stats {
	switch classselected {
	case "1":
		return classFighter
	case "2":
		return classBarbarian
	case "3":
		return classRogue
	case "4":
		return classWizard
	case "5":
		return classCleric
	case "6":
		return classWarlock
	case "7":
		return classBard
	case "8":
		return classDruid
	case "9":
		return classRanger
	default:
		return characterStats
	}
}
