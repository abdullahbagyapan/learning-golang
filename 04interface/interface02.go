package main

import "fmt"

// our structs
type ninjaStar struct {
	owner string
}
type ninjaSword struct {
	owner string
}

// generic interface
type ninjaWeapon interface {
	attack()
}

// functions for each struct
func (star ninjaStar) attack() {
	fmt.Println("Throwing ninja star")
}

func (sword ninjaSword) attack() {
	fmt.Println("Swinging ninja sword")
}

func main() {

	weapons := []ninjaWeapon{ninjaStar{owner: "Jack"}, ninjaSword{owner: "Jack"}}

	for _, weapon := range weapons {
		weapon.attack()
	}

}
