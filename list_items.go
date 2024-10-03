package main

type Item struct {
	Armet          Item_Armor
	CrusaderHelm   Item_Armor
	Doublet        Item_Armor
	RivetedGloves  Item_Armor
	OccultistPants Item_Armor
	LightfootBoots Item_Armor
	RadiantCloak   Item_Armor
}

var item = Item{
	Armet:          createItem("itemArmet2.json"),
	CrusaderHelm:   createItem("itemCrusaderHelm.json"),
	Doublet:        createItem("itemDoublet.json"),
	RivetedGloves:  createItem("itemRivetedGloves.json"),
	OccultistPants: createItem("itemOccultistPants.json"),
	LightfootBoots: createItem("itemLightfootBoots.json"),
	RadiantCloak:   createItem("itemRadiantCloak.json"),
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
	},
}

//var Items = []Item_Armor{
//	Armet:        createItem("itemArmet2.json"),
//	CrusaderHelm: createItem("itemCrusaderHelm.json"),
//	Doublet:      createItem("itemDoublet.json"),
//}
