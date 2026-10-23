package orders

import "errors"

type Item struct {
	Name  string
	Price int
}

type Order struct {
	ID     int
	Status string
	Items  []Item
}

var StatusNames = make(map[string]string)

func init() {
	StatusNames["new"] = "новый"
	StatusNames["paid"] = "оплачен"
	StatusNames["completed"] = "выполнен"
	StatusNames["cancelled"] = "отменен"
}

func NewOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, errors.New("ID заказа должен быть положительным")
	}

	return Order{
		ID:     id,
		Status: "new",
		Items:  make([]Item, 0),
	}, nil
}

func AddItem(order *Order, item Item) error {
	if order == nil {
		return errors.New("заказ не должен быть nil")
	}
	if item.Name == "" {
		return errors.New("пустое название позиции")
	}
	if item.Price < 0 {
		return errors.New("отрицательнаяцена позиции")
	}

	order.Items = append(order.Items, item)
	return nil
}

func Total(order Order) int {
	var total int
	for _, item := range order.Items {
		total += item.Price
	}
	return total
}
