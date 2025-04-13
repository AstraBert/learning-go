package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	descriptions := []string{"quarter_of_a_dozen", "half_of_a_dozen", "dozen", "small_gross", "gross", "great_gross"}
	quantities := []int{3, 6, 12, 120, 144, 1728}
	for i := 0; i < len(quantities); i++ {
		units[descriptions[i]] = quantities[i]
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	newBill := make(map[string]int)
	return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	valUnit, exUnit := units[unit]
	if !exUnit {
		return exUnit
	} else {
		valBill, exBill := bill[item]
		if exBill {
			bill[item] = valBill + valUnit
		} else {
			bill[item] = valUnit
		}
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	valBill, exBill := bill[item]
	valUnit, exUnit := units[unit]
	if !exBill || !exUnit {
		return false
	} else {
		if valBill - valUnit < 0 {
			return false
		} else if valBill - valUnit == 0 {
			delete(bill, item)
			return true
		} else {
			bill[item] = valBill - valUnit
			return true
		}
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qt, ex := bill[item]
	return qt, ex
}
