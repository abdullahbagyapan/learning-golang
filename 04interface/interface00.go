package main

type I interface {
	run()
}

type A struct {
	message string
}

func main() {

	var i I = A{"Hello World"}
	i.run()
	// output will be "Hello World"
}

// type A implement I
// (value , type)
func (t A) run() {
	println(t.message)
}
