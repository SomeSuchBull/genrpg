package pirateborg

type Item struct {
	Price string
	Name  string
}

type Weapon struct {
	Item
	Damage string
	Extra  string
}

type Clothing struct {
	Item
	Description string
	Armor       string
}

type Hat struct {
	Item
	Armor string
	Extra string
}
