package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "truck" || kind == "car" {
		return true
	} else {
		return false
	}
}

// LexiconFirst returns the string that comes first in lexicographic order.
func LexiconFirst(s1, s2 string) string {
	if s1 > s2 {
		return s2
	} else {
		return s1
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	 return LexiconFirst(option1, option2) + " is clearly the better choice."
}

// PriceDrop returns the price drop according to the age of the vehicle.
func PriceDrop(age float64) float64 {
	if age < 3 {
		return 0.8
	} else if age >= 3 && age < 10 {
		return 0.7
	} else {
		return 0.5
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	return originalPrice * PriceDrop(age)
}

