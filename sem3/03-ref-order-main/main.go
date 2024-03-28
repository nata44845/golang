package main

import (
	"fmt"
)

type order struct {
	quantity, price           int
	customerName, productName string
}

func printOrder(o order) string {
	total := o.quantity * o.price
	return fmt.Sprintf("Заказ от %s: %dx %s (Итого: $%d)\n", o.customerName, o.quantity, o.productName, total)
}

func main() {
	fmt.Println(printOrder(order{12, 300, "Иван", "Книга"}))
}
