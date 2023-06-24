package sharedkernel

type ItemType int32

const (
	ItemTypeCappuccino ItemType = iota
	ItemTypeCoffeeBlack
	ItemTypeCoffeeWithRoom
	ItemTypeEspresso
	ItemTypeEspressoDouble
	ItemTypeLatte
	ItemTypeCakePop
	ItemTypeCroissant
	ItemTypeMuffin
	ItemTypeCroissantChocolate
)

func (e ItemType) String() string {
	return []string{
		"CAPPUCCINO",
		"COFFEE_BLACK",
		"COFFEE_WITH_ROOM",
		"ESPRESSO",
		"ESPRESSO_DOUBLE",
		"LATTE",
		"CAKEPOP",
		"CROISSANT",
		"MUFFIN",
		"CROISSANT_CHOCOLATE",
		"CAPPUCCINO",
	}[e]
}
