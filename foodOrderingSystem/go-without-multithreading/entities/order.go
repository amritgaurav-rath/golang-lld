package entities
import "app/foodOrderingSystem/go-without-multithreading/enums"

// Order isolates purely state representations decoupling logic structurally natively
type Order struct {
	OrderID        string
	CustomerName   string
	RestaurantName string
	Items          map[string]int
	Status         enums.OrderStatus
}
