package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user
type byName []user

func (age byAge) Len() int           { return len(age) }
func (age byAge) Less(i, j int) bool { return age[i].Age < age[j].Age }
func (age byAge) Swap(i, j int)      { age[i], age[j] = age[j], age[i] }

func (name byName) Len() int           { return len(name) }
func (name byName) Less(i, j int) bool { return name[i].First < name[j].First }
func (name byName) Swap(i, j int)      { name[i], name[j] = name[j], name[i] }

func main() {

	u1 := user{
		First:   "jon",
		Age:     63,
		Last:    "uepppppa",
		Sayings: []string{"thats", "what", "she", "said"},
	}

	u2 := user{
		First:   "oba",
		Age:     43,
		Last:    "rapaz",
		Sayings: []string{"oi", "kkkk", "yeo", "jedi"},
	}
	u3 := user{
		First:   "dooku",
		Age:     83,
		Last:    "aaaaaaaaaaaaa",
		Sayings: []string{"jedi", "quite", "opposite", "yaag"},
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(byAge(users))
	fmt.Println(users)

	sort.Sort(byName(users))
	fmt.Println(users)

	for i, user := range users {
		fmt.Print(i, ".", " Name: ", user.First, "Lastname", user.Last)
		for _, saying := range user.Sayings {
			fmt.Print("\t", "Sayings: ", saying, "\n")
		}
	}

}
