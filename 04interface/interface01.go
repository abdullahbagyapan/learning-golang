package main

import "fmt"

type gun struct {
	owner string
}

type weapon interface {
	reload()
}

func (gun gun) reload() {
	fmt.Println("Reloading")
}

func main() {

	weapon := weapon(gun{owner: "Jack"})

	weapon.reload()

}
