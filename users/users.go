package users

import "fmt"

type User struct {
	Id      int
	Name    string
	Age     int
	frineds []User
}

func (u User) SayHello() {
	fmt.Println("Hola me llamo", u.Name)
}

func (u *User) AddFriend(friend User) {
	u.frineds = append(u.frineds, friend)
}

func (u User) ListFriends() {

	for i, f := range u.frineds {
		fmt.Printf("%d. %s\n", i+1, f.Name)
	}
}

func (u User) findFriend(Id int) int {
	for i, f := range u.frineds {
		if f.Id == Id {
			return i
		}
	}
	return -1
}

func (u *User) RemoveFriend(Id int) {
	index := u.findFriend(Id)

	if index < 0 {
		return
	}

	u.frineds = append(u.frineds[:index], u.frineds[index+1:]...)
}
