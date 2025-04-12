package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// SelectTable returns the table number as a three-letters code
func SelectTable(tableNum int) string {
	if tableNum < 10 {
		return fmt.Sprintf("00%d", tableNum)
	} else if tableNum >= 10 && tableNum < 100 {
		return fmt.Sprintf("0%d", tableNum)
	} else {
		return fmt.Sprintf("%d", tableNum)
	}
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return Welcome(name) + "\n" + fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %.1f meters from here.", SelectTable(table), direction, distance) + "\n" + fmt.Sprintf("You will be sitting next to %s.", neighbor) 
}
