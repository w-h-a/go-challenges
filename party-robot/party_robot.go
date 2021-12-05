package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	str := "Welcome to my party, " + name + "!" + "\n"
	str += fmt.Sprintf("You have been assigned to table %03d.", table) + " "
	str += "Your table is " + direction + fmt.Sprintf(", exactly %.1f ", distance) + "meters from here.\n"
	str += "You will be sitting next to " + neighbor + "."
	return str
}
