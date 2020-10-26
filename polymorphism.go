// Created by Noel Willems
// Sampling polymorphism in Golang.
package main
import ("fmt")

// Animal interface
type Animal interface {
	// Prints color
	GetColor()
	// Speaks sound
	Speak()
}

// Uses Animal interface functions
type Cat struct {
	sound string
	color string
}
// Animal interface functions:
func (c Cat)Speak() { fmt.Print(c.sound) }
func (c Cat) GetColor() { fmt.Println(c.color) }

// Uses Animal interface functions
type Horse struct {
	sound string
	color string
}
// Animal interface functions:
func (h Horse) Speak() { fmt.Print(h.sound) }
func (h Horse) GetColor() { fmt.Println(h.color) }


// Uses Animal interface functions
type Cow struct {
	sound string
	color string
}
// Animal interface functions
func (c Cow) Speak() { fmt.Print(c.sound) }
func (c Cow) GetColor() { fmt.Println(c.color) }

func main() {
	// Slice of Animal interface type
	var zoo []Animal
	cat := Cat{ sound: "Cat says miau and is ", color: "tortoiseshell" }
	zoo = append(zoo, cat)
	horse:= Horse{ sound: "Horse says whinny and is ", color: "chestnut" }
	zoo = append(zoo, horse)
	cow:= Cow{ sound: "Cow says moooo and is ", color: "spotted" }
	zoo = append(zoo, cow)

	// Call .Speak() and .GetColor() on each Animal in zoo
	for i := range zoo {
		zoo[i].Speak()
		zoo[i].GetColor()
	}
}