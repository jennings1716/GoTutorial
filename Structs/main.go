package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func main() {
	name := "JenJen"
	fmt.Print(&name)

	alex := person{"alex", "anderson", contactInfo{"jen@gmail", 123}}
	alex.firstName = "jennings"

	var jeniffer person
	jeniffer.firstName = "Jeniffer"
	jeniffer.lastName = "James"

	alex.print()
	jPointer := &jeniffer
	fmt.Println(jPointer)
	jPointer.updateName("Jennniippperrrrrr")
	jeniffer.updateName("Works")
	jeniffer.print()

}
