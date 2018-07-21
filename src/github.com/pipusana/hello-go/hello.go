package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) fullname() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p *Person) growup() {
	p.Age = p.Age + 1
}

func init() {
	fmt.Println("init")
}

func main() {
	name := map[string]int{
		"a": 20,
		"b": 30,
		"c": 40,
	}

	for k, v := range name {
		fmt.Println(k, v)
	}

	// me := Person{
	// 	FirstName: "Pipusana",
	// 	LastName:  "petgumpoom",
	// 	Age:       24,
	// }

	// fmt.Printf("%+v", me)

	people := []Person{
		Person{
			FirstName: "Pipusana",
			LastName:  "petgumpoom",
			Age:       24,
		},
		Person{
			FirstName: "ter",
			LastName:  "bone",
			Age:       23,
		},
	}

	for _, p := range people {
		fmt.Println(p.fullname())
		p.growup()

		fmt.Sprintf("%d", p.Age)

		// result, err := json.Marshal(p)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(string(fullname))
	}
}
