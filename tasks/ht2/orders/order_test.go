package orders

import "testing"

func TestNewOrder(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Fatalf("NewOrder вернул неожиданную ошибку: %v", err)
	}
	if got := Total(order); got != 0 {
		t.Fatalf("ненулевая сумма пустого заказа = %d", got)
	}

	if _, err := NewOrder(0); err == nil {
		t.Fatal("NewOrder должнен вернуть ошибку для неположительного ID")
	}
}

func TestAddItemAndTotal(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	items := []Item{
		{Name: "Кофе", Price: 40000},
		{Name: "Десерт", Price: 35000},
	}
	for _, item := range items {
		if err := AddItem(&order, item); err != nil {
			t.Fatalf("AddItem вернул неожиданную ошибку: %v", err)
		}
	}

	if got, want := Total(order), 75000; got != want {
		t.Fatalf("сумма заказа = %d, ожидалось %d", got, want)
	}
}

func TestAddItemInvalidData(t *testing.T) {
	tests := []struct {
		name  string
		order *Order
		item  Item
	}{
		{name: "nil order", order: nil, item: Item{Name: "Кофе", Price: 100}},
		{name: "empty name", order: &Order{}, item: Item{Price: 100}},
		{name: "negative price", order: &Order{}, item: Item{Name: "Кофе", Price: -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddItem(tt.order, tt.item); err == nil {
				t.Fatal("AddItem должен вернуть ошибку")
			}
		})
	}
}

func TestOrderCopy(t *testing.T) {
	original, err := NewOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	copyOfOrder := original
	copyOfOrder.Status = "paid"

	if original.Status == copyOfOrder.Status {
		t.Fatal("изменение копии Order изменило оригинальный Order")
	}
}

func TestStatusNames(t *testing.T) {
	StatusNames["refunded"] = "Возвращён"
	if got := StatusNames["refunded"]; got != "Возвращён" {
		t.Fatalf("название статуса = %q, ожидалось %q", got, "Возвращён")
	}
}
