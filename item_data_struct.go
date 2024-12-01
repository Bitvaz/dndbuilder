package main

type BaseAttribute struct {
	Strength        map[int]int `json:"strength"`
	Vigor           map[int]int `json:"vigor"`
	Agility         map[int]int `json:"agility"`
	Dexterity       map[int]int `json:"dexterity"`
	Will            map[int]int `json:"will"`
	Knowledge       map[int]int `json:"knowledge"`
	Resourcefulness map[int]int `json:"resourcefulness"`
}

type Item_Armor struct {
	Name                string          `json:"name"`
	ArmorRatings        map[int][]int   `json:"armorRatings"`
	MaxHealthAdd        map[int]int     `json:"maxHealthAdd"`
	MagicResistance     map[int][]int   `json:"magicResistance"`
	ProjectileReduction float64         `json:"projectileReduction"` // %
	HeadshotReduction   float64         `json:"headshotReduction"`   // %
	BaseAttribute       BaseAttribute   `json:"BaseAttribute"`       // BaseStats 1 to 7
	MoveSpeed           map[int]int     `json:"moveSpeed"`           // %
	MoveSpeedBonus      map[int]float64 `json:"moveSpeedbonus"`
	ArmorPenetration    float64         `json:"armorpPenetration"`
	MagicPenetration    float64         `json:"magicPenetration"`
	MagicalPower        int             `json:"magicalPower"`
	RegularInteraction  int             `json:"regularInteraction"`
	MagicalHealing      int             `json:"magicalHealing"`
	Luck                map[int]int     `json:"luck"`
	TruePhysicalDamage  map[int]int     `json:"truePhysicaldamage"`
	TrueMagicalDamage   map[int]int     `json:"trueMagicaldamage"`
	ArmorType           string          `json:"armorType"`    // Armor type (e.g., "Plate")
	Rarities            []int           `json:"rarities"`     // List of rarities
	SlotType            string          `json:"slotType"`     // Slot type (e.g., "Head")
	InvWidth            int             `json:"invWidth"`     // Inventory width
	InvHeight           int             `json:"invHeight"`    // Inventory height
	FlavorText          string          `json:"flavorText"`   // Description text
	MaxAmmoCount        int             `json:"maxAmmoCount"` // Max ammo count
	MaxCount            int             `json:"maxCount"`     // Max count
	AP                  map[int]int     `json:"ap"`           // Action points for levels 3 to 7
	XP                  map[int]int     `json:"xp"`           // XP for levels 1 to 7
	GearScore           map[int]int     `json:"gearScore"`    // Gear score for levels 0 to 7
	Classes             []string        `json:"classes"`      // List of classes (e.g., "Fighter")
	SellPrices          map[int]int     `json:"sellPrices"`   // Sell prices for levels 1 to 7
	NumEnchants         map[int]int     `json:"numEnchants"`  // Number of enchants for levels 3 to 7

}

type Item_Accessory struct {
	Name            string
	BaseAttribute   map[int]int   `json:"baseAttribute"`  // BaseStats 1 to 7
	BaseAttribute2  map[int]int   `json:"baseAttribute2"` // BaseStats 1 to 7
	MagicResistance map[int][]int `json:"magicResistance"`
	Luck            map[int]int   `json:"luck"`
	Type            string        `json:"type"`         // Accessory type
	Rarities        []int         `json:"rarities"`     // List of rarities
	SlotType        string        `json:"slotType"`     // Slot type (e.g., "Necklace")
	InvWidth        int           `json:"invWidth"`     // Inventory width
	InvHeight       int           `json:"invHeight"`    // Inventory height
	FlavorText      string        `json:"flavorText"`   // Description text
	MaxAmmoCount    int           `json:"maxAmmoCount"` // Max ammo count
	MaxCount        int           `json:"maxCount"`     // Max count
	AP              map[int]int   `json:"ap"`           // Action points for levels 3 to 7
	XP              map[int]int   `json:"xp"`           // XP for levels 1 to 7
	GearScore       map[int]int   `json:"gearScore"`    // Gear score for levels 0 to 7
	SellPrices      map[int]int   `json:"sellPrices"`   // Sell prices for levels 1 to 7
	NumEnchants     map[int]int   `json:"numEnchants"`  // Number of enchants for levels 3 to 7

}

type Item_Weapon struct {
	Name                 string
	Classes              []string    `json:"classes"` // List of classes (e.g., "Fighter", "Barbarian")
	PhysicalWeaponDamage map[int]int `json:"baseAttribute"`
	MoveSpeed            int         `json:"moveSpeed"`
	DamageType           string      `json:"damageType"`
	ImpactZones          []int       `json:"impactZones"`
	ImpactPower          int         `json:"impactPower"`
	WeaponType           string      `json:"type"`        // Weapon type (e.g., "Axe")
	Rarities             []int       `json:"rarities"`    // List of rarities
	HandType             string      `json:"handType"`    // Hand type (e.g., "Two Handed")
	SlotType             string      `json:"slotType"`    // Slot type (e.g., "Main Hand")
	GearScore            map[int]int `json:"gearScore"`   // Gear score for levels 0 to 7
	InvWidth             int         `json:"invWidth"`    // Inventory width
	InvHeight            int         `json:"invHeight"`   // Inventory height
	FlavorText           string      `json:"flavorText"`  // Description text
	AP                   map[int]int `json:"ap"`          // Action points for levels 3 to 7
	XP                   map[int]int `json:"xp"`          // XP for levels 1 to 7
	SellPrices           map[int]int `json:"sellPrices"`  // Sell prices for levels 1 to 7
	NumEnchants          map[int]int `json:"numEnchants"` // Number of enchants for levels 3 to 7

}
type EnchantmentValue struct {
	Slottypes []string
	Attribute []float64
}

type Enchantment struct {
	AllAttributes            EnchantmentValue
	Strength                 EnchantmentValue
	Vigor                    EnchantmentValue
	Agility                  EnchantmentValue
	Dexterity                EnchantmentValue
	Will                     []int
	Knowledge                []int
	Resourcefulness          []int
	ActionSpeed              []float64
	SpellCastingSpeed        []float64
	BuffDurationBonus        []float64
	DebuffDurationBonus      []float64
	RegularInteractionSpeed  []float64
	MagicalInteractionSpeed  []float64
	MemoryCapacityBonus      []float64
	MemoryCapacityAdd        []float64
	MaxHealthBonus           []float64
	MaxHealthAdd             []float64
	MovementSpeedBonus       []float64
	MovementSpeedAdd         []float64
	PhysicalPower            []float64
	PhysicalPowerBonus       []float64
	MagicalPower             []float64
	MagicPowerBonus          []float64
	ArmorRating              []float64
	ArmorPenetration         []float64
	AdditionalPhysicalDamage []float64
	TruePhysicalDamage       []float64
	PhysicalDamageBonus      []float64
	PhysicalWeaponDamage     []float64
	MagicPenetration         []float64
	AdditionalMagicalDamage  []float64
	TrueMagicalDamage        []float64
	PhysicalDamageReduction  []float64
	MagicResistance          []float64
	MagicalDamageReduction   []float64
	Luck                     []int
}

/*
var allEnchanment = Enchantment{
	AllAttributes: EnchantmentValue{
		Slottypes: []string{"head", "chest", "legs", "hands", "foot"},
		Attribute: []float64{}},
	Strength: EnchantmentValue{
		Slottypes: []string{"head", "chest", "legs", "hands", "foot"},
		Attribute: []float64{1, 2}},
	Vigor: EnchantmentValue{
		Slottypes: []string{"head", "chest", "legs", "hands", "foot"},
		Attribute: []float64{1, 2}},
	Agility: EnchantmentValue{
		Slottypes: []string{"head", "chest", "legs", "hands", "foot"},
		Attribute: []float64{1, 2}},
	Dexterity: EnchantmentValue{
		Slottypes: []string{"head", "chest", "legs", "hands", "foot"},
		Attribute: []float64{1, 2}},
}
*/

type Listechanment struct {
}
