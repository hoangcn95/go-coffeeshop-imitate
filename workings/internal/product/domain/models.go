package domain

/*
Fixme: should not put model's tags in domain layer
*/
type ItemTypeDto struct {
	Name  string  `json:"name"`
	Type  int     `json:"type"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ItemDto struct {
	Price float64 `json:"price"`
	Type  int     `json:"type"`
}
