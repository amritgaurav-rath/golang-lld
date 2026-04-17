package entities

// Restaurant strictly evaluates capability dynamically via memory maps naturally bypassing external Database calls natively.
type Restaurant struct {
	Name          string
	Rating        float64
	MaxOrders     int
	CurrentOrders int
	Menu          map[string]float64
}

func NewRestaurant(name string, maxOrders int, rating float64) *Restaurant {
	return &Restaurant{
		Name:          name,
		Rating:        rating,
		MaxOrders:     maxOrders,
		CurrentOrders: 0,
		Menu:          make(map[string]float64),
	}
}

func (r *Restaurant) UpdateMenu(itemName string, price float64) {
	r.Menu[itemName] = price
}

func (r *Restaurant) UpdateCapacity(capacity int) {
	r.MaxOrders = capacity
}

func (r *Restaurant) CanFulfill(items map[string]int) bool {
	// Safely drops request mappings failing to fulfill native array capacity
	if r.CurrentOrders >= r.MaxOrders {
		return false
	}
	// Blocks matching if specific request matrix isn't mathematically present in menu
	for itemName := range items {
		if _, exists := r.Menu[itemName]; !exists {
			return false
		}
	}
	return true
}

func (r *Restaurant) IncrementCurrentOrders() { r.CurrentOrders++ }

func (r *Restaurant) MarkOrderCompleted() {
	if r.CurrentOrders > 0 {
		r.CurrentOrders--
	}
}
