package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    amount, ok := units[unit]
    if !ok {
        return ok
    }
    bill[item] += amount
    return ok
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    amount1, ok1 := bill[item]
    amount2, ok2 := units[unit]
    if !(ok1 && ok2) {
        return false
    }
    if amount1 - amount2 < 0 {
        return false
    } else if amount1 - amount2 == 0 {
        delete(bill, item)
    } else {
        bill[item] -= amount2
    }    
    return ok1 && ok2
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amount, ok := bill[item]
    return amount, ok
}
