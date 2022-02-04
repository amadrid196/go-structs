package main

import (
	"fmt"

	"github.com/amadrid196/go-structs.git/users"
)

func main() {
	martha := users.User{Id: 1, Name: "Maria", Age: 20}
	pedro := users.User{Id: 2, Name: "Pedro", Age: 22}
	carlos := users.User{Id: 3, Name: "Carlos", Age: 26}
	martha.SayHello()
	martha.AddFriend(pedro)
	martha.AddFriend(carlos)

	martha.ListFriends()

	martha.RemoveFriend(10)

	fmt.Println("================================")

	martha.ListFriends()
}
