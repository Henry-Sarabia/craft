package craft

// Item represents an item that would commonly be found in a medieval RPG game.
type Item struct {
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Weight      float64 `json:"weight"`
}
