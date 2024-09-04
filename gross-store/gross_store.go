package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]

	if !unitExists {
		return false
	}

	_, itemExists := bill[item]

	if itemExists {
		bill[item] += unitValue
	} else {
		bill[item] = unitValue
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, itemInBill := bill[item]
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	if !itemInBill {
		return false
	} else {
		if itemValue-unitValue < 0 {
			return false
		} else if itemValue-unitValue == 0 {
			delete(bill, item)
		} else {
			bill[item] -= unitValue
		}
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemValue, itemExists := bill[item]

	if !itemExists {
		return 0, false
	}

	return itemValue, true
}
