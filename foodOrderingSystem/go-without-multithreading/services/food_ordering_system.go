package services

import (
	"app/foodOrderingSystem/go-without-multithreading/entities"
	"app/foodOrderingSystem/go-without-multithreading/enums"
	"app/foodOrderingSystem/go-without-multithreading/strategy"
	"fmt"
)

type FoodOrderingSystem struct {
	Restaurants map[string]*entities.Restaurant
	Orders      map[string]*entities.Order
}

var instance *FoodOrderingSystem

func GetInstance() *FoodOrderingSystem {
	if instance == nil {
		instance = &FoodOrderingSystem{
			Restaurants: make(map[string]*entities.Restaurant),
			Orders:      make(map[string]*entities.Order),
		}
	}
	return instance
}

func (s *FoodOrderingSystem) OnboardRestaurant(name string, maxOrders int, rating float64, menu map[string]float64) {
	r := entities.NewRestaurant(name, maxOrders, rating)
	for item, price := range menu {
		r.UpdateMenu(item, price)
	}
	s.Restaurants[name] = r
}

func (s *FoodOrderingSystem) UpdateMenu(name string, action string, itemName string, price float64) error {
	restaurant, exists := s.Restaurants[name]
	if !exists {
		return fmt.Errorf("restaurant structurally unmapped natively")
	}

	if action == "add" || action == "update" {
		restaurant.UpdateMenu(itemName, price)
	}
	return nil
}

func (s *FoodOrderingSystem) UpdateCapacity(name string, capacity int) error {
	restaurant, exists := s.Restaurants[name]
	if !exists {
		return fmt.Errorf("restaurant mathematically unmapped natively")
	}
	restaurant.UpdateCapacity(capacity)
	return nil
}

func (s *FoodOrderingSystem) PlaceOrder(orderId, user string, items map[string]int, strat strategy.RestaurantSelectionStrategy) (*entities.Order, error) {
	var eligibleRestaurants []*entities.Restaurant

	for _, r := range s.Restaurants {
		if r.CanFulfill(items) {
			eligibleRestaurants = append(eligibleRestaurants, r)
		}
	}

	if len(eligibleRestaurants) == 0 {
		return nil, fmt.Errorf("Output: Order can't be fulfilled")
	}

	selected := strat.SelectRestaurant(eligibleRestaurants, items)
	if selected == nil {
		return nil, fmt.Errorf("strategy failed mathematical match cleanly")
	}

	selected.IncrementCurrentOrders()

	order := &entities.Order{
		OrderID:        orderId,
		CustomerName:   user,
		RestaurantName: selected.Name,
		Items:          items,
		Status:         enums.Accepted,
	}
	s.Orders[orderId] = order

	fmt.Printf("Output: Order assigned to %s\n", selected.Name)
	return order, nil
}

func (s *FoodOrderingSystem) UpdateOrderStatus(orderId string, status enums.OrderStatus) error {
	order, exists := s.Orders[orderId]
	if !exists {
		return fmt.Errorf("order mathematically unmapped natively")
	}

	if order.Status == enums.Completed {
		return fmt.Errorf("order inherently tracked as securely completed")
	}

	if status == enums.Completed {
		order.Status = status
		s.Restaurants[order.RestaurantName].MarkOrderCompleted()
	}

	return nil
}
