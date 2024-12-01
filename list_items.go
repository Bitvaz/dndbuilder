package main

type Item struct {
	Armet          Item_Armor
	CrusaderHelm   Item_Armor
	Doublet        Item_Armor
	RivetedGloves  Item_Armor
	OccultistPants Item_Armor
	LightfootBoots Item_Armor
	RadiantCloak   Item_Armor
	AdventureCloak Item_Armor
}

var item = Item{

	Armet:          CreateItem("./data/itemArmet2.json"),
	CrusaderHelm:   CreateItem("./data/itemCrusaderHelm.json"),
	Doublet:        CreateItem("./data/itemDoublet.json"),
	RivetedGloves:  CreateItem("./data/itemRivetedGloves.json"),
	OccultistPants: CreateItem("./data/itemOccultistPants.json"),
	LightfootBoots: CreateItem("./data/itemLightfootBoots.json"),
	RadiantCloak:   CreateItem("./data/itemRadiantCloak.json"),
	AdventureCloak: CreateItem("./data/itemAdventureCloak.json"),
}

type List_Items struct {
	ItemsArmor []Item_Armor
}

var Items = List_Items{
	ItemsArmor: []Item_Armor{
		item.Armet,
		item.CrusaderHelm,
		item.Doublet,
		item.RivetedGloves,
		item.OccultistPants,
		item.LightfootBoots,
		item.RadiantCloak,
		item.AdventureCloak,
	},
}
