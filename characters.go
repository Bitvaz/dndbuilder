package main

// STRUCT STAT
type Stats struct {
	Strength        int `form:"strength"`
	Vigor           int `form:"vigor"`
	Agility         int `form:"agility"`
	Dexterity       int `form:"dexterity"`
	Will            int `form:"will"`
	Knowledge       int `form:"knowledge"`
	Resourcefulness int `form:"resourcefulness"`
}

// STRUCT COMPLETE STATS
type ComputedStats struct {
	Health                  float64
	ActionSpeed             float64
	RegularInteractionSpeed float64
	MoveSpeed               float64
	MoveSpeedCalc           float64
	PhysicalPower           float64
	PhysicalPowerBonus      float64
	HealthRecovery          float64
	ManualDexterity         float64
	EquipSpeed              float64
	MagicalPower            float64
	BuffDuration            float64
	MagicResistance         float64
	DebuffDuration          float64
	MemoryCapacity          float64
	SpellRecovery           float64
	SpellCastingSpeed       float64
	MagicalInteractionSpeed float64
	Persuasiveness          float64
	CooldownReduction       float64
}

var characterStats = Stats{
	Strength:        15,
	Vigor:           15,
	Agility:         15,
	Dexterity:       15,
	Will:            15,
	Knowledge:       15,
	Resourcefulness: 15,
}

var classFighter = Stats{
	Strength:        15,
	Vigor:           15,
	Agility:         15,
	Dexterity:       15,
	Will:            15,
	Knowledge:       15,
	Resourcefulness: 15,
}

var classBarbarian = Stats{
	Strength:        20,
	Vigor:           25,
	Agility:         13,
	Dexterity:       12,
	Will:            18,
	Knowledge:       5,
	Resourcefulness: 12,
}

var classRogue = Stats{
	Strength:        9,
	Vigor:           10,
	Agility:         21,
	Dexterity:       25,
	Will:            10,
	Knowledge:       10,
	Resourcefulness: 20,
}

var classWizard = Stats{
	Strength:        6,
	Vigor:           7,
	Agility:         15,
	Dexterity:       17,
	Will:            20,
	Knowledge:       25,
	Resourcefulness: 15,
}

var classCleric = Stats{
	Strength:        11,
	Vigor:           13,
	Agility:         12,
	Dexterity:       14,
	Will:            23,
	Knowledge:       20,
	Resourcefulness: 12,
}

var classWarlock = Stats{
	Strength:        11,
	Vigor:           14,
	Agility:         14,
	Dexterity:       15,
	Will:            22,
	Knowledge:       15,
	Resourcefulness: 14,
}

var classBard = Stats{
	Strength:        13,
	Vigor:           13,
	Agility:         13,
	Dexterity:       20,
	Will:            11,
	Knowledge:       20,
	Resourcefulness: 15,
}

var classDruid = Stats{
	Strength:        12,
	Vigor:           13,
	Agility:         12,
	Dexterity:       12,
	Will:            18,
	Knowledge:       20,
	Resourcefulness: 18,
}

var classRanger = Stats{
	Strength:        10,
	Vigor:           10,
	Agility:         20,
	Dexterity:       18,
	Will:            10,
	Knowledge:       12,
	Resourcefulness: 25,
}