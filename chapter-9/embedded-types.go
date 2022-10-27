package main

import "fmt"

/*
A struct's fields represent the has-a relation.
For example a Person has a Name.
*/
type Person struct {
	Name string
}

/*
The Android struct has a is-a relation with the Person struct: an Android is a Person.
Go supports relationships like this by using embedded type a.k.a anonymous fields
We use the type (Person) and don't give it a name:
*/
type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	p := new(Person)
	p.Talk()

	a := new(Android)

	/*
	   The is-a relationsip works this way intuitively:
	   People can talk,
	   an Android is a Person,
	   therefore an Android can talk.
	*/

	// The person struct can be accessed using the type name.
	a.Person.Talk()
	// We can also call any Person methods directly on the Android.
	a.Talk()
}
