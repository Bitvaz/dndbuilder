package main

//CALCULA PHYSICAL POWER
func CalculatePhysicalPower(strength int) float64 {
	return float64(strength-15) * 1.0
}

//CALCULA PHYSICAL POWER BONUS
func CalculatePhysicalPowerBonus(x int) float64 {
	var percentage float64

	switch {
	case x <= 0:
		percentage = -80.0
	case x <= 5:
		percentage = -80.0 + 10.0*float64(x)
	case x <= 7:
		percentage = -30.0 + 5.0*float64(x-5)
	case x <= 11:
		percentage = -20.0 + 3.0*float64(x-7)
	case x <= 15:
		percentage = -8.0 + 2.0*float64(x-11)
	case x <= 50:
		percentage = 0.0 + 1.0*float64(x-15)
	case x <= 100:
		percentage = 35.0 + 0.5*float64(x-50)
	default:
		percentage = 60.0
	}

	return percentage
}

//CALCULA MOVE SPEED
func CalculateMoveSpeed(agility int) float64 {
	var moveSpeed float64

	switch {
	case agility <= 0:
		moveSpeed = 300 - 10 // Base value of 300 minus 10 for 0 Agility
	case agility <= 10:
		moveSpeed = 300 + (-10 + 0.5*float64(agility)) // Add 300 to the calculation
	case agility <= 15:
		moveSpeed = 300 + (-5 + 1*float64(agility-10)) // Add 300 to make output positive
	case agility <= 75:
		moveSpeed = 300 + 0.75*float64(agility-15)
	case agility <= 100:
		moveSpeed = 300 + (45 + 0.5*float64(agility-75))
	default:
		moveSpeed = 330
	}

	return moveSpeed
}

/*
// CALCULA ACTION SPEED
func CalculateActionSpeed(agility, dexterity int) float64 {
	rating := float64(agility)*0.25 + float64(dexterity)*0.75
	var actionSpeed float64
	switch {
	case rating <= 0:
		actionSpeed = -38
	case rating <= 10:
		actionSpeed = -38 + 3*rating
	case rating <= 13:
		actionSpeed = -8 + 2*(rating-10)
	case rating <= 25:
		actionSpeed = -2 + 1*(rating-13)
	case rating <= 41:
		actionSpeed = 10 + 1.5*(rating-25)
	case rating <= 50:
		actionSpeed = 34 + 1*(rating-41)
	case rating <= 100:
		actionSpeed = 43 + 0.5*(rating-50)
	default:
		actionSpeed = 68
	}
	return actionSpeed
}
*/
// CALCULA MAX HEALTH
func CalculateHealth(strength, vigor int) float64 {
	// Calculate Max Health Rating
	maxHealthRating := float64(strength)*0.25 + float64(vigor)*0.75

	// Determine Max Health based on the rating
	var maxHealth float64

	switch {
	case maxHealthRating <= 0:
		maxHealth = 75
	case maxHealthRating <= 10:
		maxHealth = 75 + 3*maxHealthRating
	case maxHealthRating <= 50:
		maxHealth = 105 + 2*(maxHealthRating-10)
	case maxHealthRating <= 75:
		maxHealth = 185 + 1*(maxHealthRating-50)
	case maxHealthRating <= 100:
		maxHealth = 210 + 0.5*(maxHealthRating-75)
	default:
		maxHealth = 210 + 0.5*(100-75) // Cap at 100 rating
	}

	return maxHealth
}

// CALCULA REGULAR INTERACTION SPEED
func CalculateRegularInteraction(agility, resourcefulness int) float64 {
	rating := float64(agility)*0.4 + float64(resourcefulness)*0.6
	var interactionSpeed float64
	switch {
	case rating <= 0:
		interactionSpeed = -26
	case rating <= 7:
		interactionSpeed = -26 + 2*rating
	case rating <= 15:
		interactionSpeed = -12 + 1.5*(rating-7)
	case rating <= 20:
		interactionSpeed = 0 + 7*(rating-15)
	case rating <= 25:
		interactionSpeed = 35 + 6*(rating-20)
	case rating <= 30:
		interactionSpeed = 65 + 5*(rating-25)
	case rating <= 35:
		interactionSpeed = 90 + 4*(rating-30)
	case rating <= 40:
		interactionSpeed = 110 + 3*(rating-35)
	case rating <= 45:
		interactionSpeed = 125 + 2*(rating-40)
	case rating <= 100:
		interactionSpeed = 135 + 1*(rating-45)
	default:
		interactionSpeed = 190
	}

	return interactionSpeed
}

// CALCULA HEALTH RECOVERY BONUS
func CalculateHealthRecovery(vigor int) float64 {
	var recovery float64

	switch {
	case vigor <= 0:
		recovery = -55
	case vigor <= 5:
		recovery = -55 + 5*float64(vigor)
	case vigor <= 15:
		recovery = -30 + 3*float64(vigor-5)
	case vigor <= 25:
		recovery = 0 + 7*float64(vigor-15)
	case vigor <= 35:
		recovery = 70 + 5*float64(vigor-25)
	case vigor <= 84:
		recovery = 120 + 2*float64(vigor-35)
	case vigor <= 85:
		recovery = 218 + 1*float64(vigor-84)
	case vigor <= 86:
		recovery = 219 + 3*float64(vigor-85)
	case vigor <= 100:
		recovery = 222 + 2*float64(vigor-86)
	default:
		recovery = 250
	}

	return recovery
}

// CALCULA MANUAL DEXTERITY
func CalculateManualDexterity(dexterity int) float64 {
	var manualDexterity float64

	switch {
	case dexterity <= 0:
		manualDexterity = -15
	case dexterity <= 15:
		manualDexterity = -15 + float64(dexterity)
	case dexterity <= 23:
		manualDexterity = 0 + 3*float64(dexterity-15)
	case dexterity <= 31:
		manualDexterity = 24 + 2*float64(dexterity-23)
	case dexterity <= 37:
		manualDexterity = 40 + 1*float64(dexterity-31)
	case dexterity <= 45:
		manualDexterity = 46 + 0.5*float64(dexterity-37)
	case dexterity <= 95:
		manualDexterity = 50 + 0.1*float64(dexterity-45)
	default:
		manualDexterity = 55
	}

	return manualDexterity
}

// CALCULA EQUIP SPEED
func CalculateEquipSpeed(dexterity int) float64 {
	var equipSpeed float64

	switch {
	case dexterity == 0:
		equipSpeed = -95
	case dexterity == 1:
		equipSpeed = -95
	case dexterity == 2:
		equipSpeed = -95 + 4
	case dexterity <= 15:
		equipSpeed = -91 + 7*float64(dexterity-2)
	case dexterity <= 35:
		equipSpeed = 0 + 5*float64(dexterity-15)
	case dexterity <= 70:
		equipSpeed = 100 + 2*float64(dexterity-35)
	case dexterity <= 100:
		equipSpeed = 170 + 1*float64(dexterity-70)
	default:
		equipSpeed = 200
	}

	return equipSpeed
}

// CALCULA MAGICAL POWER
func CalculateMagicalpower(will int) float64 {
	return float64(will-15) * 1.0
}

//CALCULA BUFF DURATION
func CalculateBuffDuration(will int) float64 {
	var buffDuration float64

	switch {
	case will == 0:
		buffDuration = -80
	case will <= 5:
		buffDuration = -80 + 10*float64(will)
	case will <= 7:
		buffDuration = -30 + 5*float64(will-5)
	case will <= 11:
		buffDuration = -20 + 3*float64(will-7)
	case will <= 15:
		buffDuration = -8 + 2*float64(will-11)
	case will <= 50:
		buffDuration = 0 + 1*float64(will-15)
	case will <= 100:
		buffDuration = 35 + 0.5*float64(will-50)
	default:
		buffDuration = 60
	}

	return buffDuration
}

//CALCULA MAGIC RESISTANCE
func CalculateMagicResistance(will int) float64 {
	var magicResistance float64

	switch {
	case will == 0:
		magicResistance = -20
	case will <= 5:
		magicResistance = -20 + 4*float64(will)
	case will <= 15:
		magicResistance = 0 + 3*float64(will-5)
	case will <= 33:
		magicResistance = 30 + 4*float64(will-15)
	case will <= 48:
		magicResistance = 102 + 3*float64(will-33)
	case will <= 58:
		magicResistance = 147 + 2*float64(will-48)
	case will <= 100:
		magicResistance = 167 + 1*float64(will-58)
	default:
		magicResistance = 209
	}

	return magicResistance
}

//CALCULA DEBUFF DURATION
func CalculateDebuffDuration(will int) float64 {
	var debuffDuration float64
	switch {
	case will <= 0:
		debuffDuration = 400
	case will <= 1:
		debuffDuration = 400
	case will <= 2:
		debuffDuration = 150
	case will <= 3:
		debuffDuration = 100
	case will <= 4:
		debuffDuration = 66.7
	case will <= 5:
		debuffDuration = 42.9
	case will <= 6:
		debuffDuration = 33.3
	case will <= 7:
		debuffDuration = 25
	case will <= 8:
		debuffDuration = 20.5
	case will <= 9:
		debuffDuration = 16.3
	case will <= 10:
		debuffDuration = 12.4
	case will <= 11:
		debuffDuration = 8.7
	case will <= 12:
		debuffDuration = 6.4
	case will <= 13:
		debuffDuration = 4.2
	case will <= 14:
		debuffDuration = 2
	case will <= 15:
		debuffDuration = 0
	case will <= 16:
		debuffDuration = -1
	case will <= 17:
		debuffDuration = -2
	case will <= 18:
		debuffDuration = -2.9
	case will <= 19:
		debuffDuration = -3.8
	case will <= 20:
		debuffDuration = -4.8
	case will <= 21:
		debuffDuration = -5.7
	case will <= 22:
		debuffDuration = -6.5
	case will <= 23:
		debuffDuration = -7.4
	case will <= 24:
		debuffDuration = -8.3
	case will <= 25:
		debuffDuration = -9.1
	case will <= 26:
		debuffDuration = -9.9
	case will <= 27:
		debuffDuration = -10.7
	case will <= 28:
		debuffDuration = -11.5
	case will <= 29:
		debuffDuration = -12.3
	case will <= 30:
		debuffDuration = -13
	case will <= 31:
		debuffDuration = -13.8
	case will <= 32:
		debuffDuration = -14.5
	case will <= 33:
		debuffDuration = -15.3
	case will <= 34:
		debuffDuration = -16
	case will <= 35:
		debuffDuration = -16.7
	case will <= 36:
		debuffDuration = -17.4
	case will <= 37:
		debuffDuration = -18
	case will <= 38:
		debuffDuration = -18.7
	case will <= 39:
		debuffDuration = -19.4
	case will <= 40:
		debuffDuration = -20
	case will <= 41:
		debuffDuration = -20.6
	case will <= 42:
		debuffDuration = -21.3
	case will <= 43:
		debuffDuration = -21.9
	case will <= 44:
		debuffDuration = -22.5
	case will <= 45:
		debuffDuration = -23.1
	case will <= 46:
		debuffDuration = -23.7
	case will <= 47:
		debuffDuration = -24.2
	case will <= 48:
		debuffDuration = -24.8
	case will <= 49:
		debuffDuration = -25.4
	case will <= 50:
		debuffDuration = -25.9
	case will <= 51:
		debuffDuration = -26.2
	case will <= 52:
		debuffDuration = -26.5
	case will <= 53:
		debuffDuration = -26.7
	case will <= 55:
		debuffDuration = -27.3
	case will <= 56:
		debuffDuration = -27.5
	case will <= 58:
		debuffDuration = -28.1
	case will <= 59:
		debuffDuration = -28.3
	case will <= 60:
		debuffDuration = -28.6
	case will <= 61:
		debuffDuration = -28.8
	case will <= 62:
		debuffDuration = -29.1
	case will <= 63:
		debuffDuration = -29.3
	case will <= 64:
		debuffDuration = -29.6
	case will <= 65:
		debuffDuration = -29.8
	case will <= 66:
		debuffDuration = -30.1
	case will <= 67:
		debuffDuration = -30.3
	case will <= 68:
		debuffDuration = -30.6
	case will <= 70:
		debuffDuration = -31
	case will <= 71:
		debuffDuration = -31.3
	case will <= 73:
		debuffDuration = -31.7
	case will <= 74:
		debuffDuration = -32
	case will <= 76:
		debuffDuration = -32.4
	case will <= 77:
		debuffDuration = -32.7
	case will <= 80:
		debuffDuration = -33.3
	case will <= 81:
		debuffDuration = -33.6
	case will <= 86:
		debuffDuration = -34.6
	case will <= 87:
		debuffDuration = -34.9
	default:
		debuffDuration = -37.5
	}
	return debuffDuration
}
