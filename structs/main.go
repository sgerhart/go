package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	//contact contactinfo//
	contactInfo


}
func main() {

	//alex := person{"alex", "anderson"}
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo.email = "aanderson@gmail.com"
	alex.contactInfo.zipCode = 19380

	tim := person {
		firstName: "Tim",
		lastName: "Tom",
		contactInfo: contactInfo {
			email: "timtom@gmail.com",
			zipCode: 19380,
	},
	}

	//fmt.Printf("%+v\n", alex)
	//fmt.Printf("%+v", tim)

	//timPointer := &tim
	//timPointer.updateName("Bobby")
	tim.updateName("bobby")
	tim.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
	//fmt.Println(p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}