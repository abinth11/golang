package main

import (
	"fmt"
	"sort"
)

func main() {
	var users = []string{"user1", "user2", "user3"}
	fmt.Printf("Type of the slice %T \n", users)

	//* multiple values can be added
	users = append(users, "user4", "user5")
	fmt.Println(users)

	//* how to slice a slice
	// users = append(users[2:])
	// users = append(users[0:3])
	users = append(users[:1])
	fmt.Println("Sliced users", users)

	//* creating slices with make method

	//* for integers default allocated value is zero
	myArr := make([]int, 5)
	fmt.Println(myArr)

	//* different ways to add elements
	myArr[0] = 10
	myArr[1] = 30
	myArr[2] = 15
	myArr[3] = 25
	myArr[4] = 5
	//* index out of range error
	// myArr[5] = 500
	fmt.Println(myArr)

	//* can add more elements with append method; it will allocate additional memory when new elements are added

	myArr = append(myArr, 50, 100, 150)
	fmt.Println("Array after adding new elements", myArr)

	fmt.Println("array before sorting", sort.IntsAreSorted(myArr))
	sort.Ints(myArr)
	fmt.Println("Sorted array", myArr)
	fmt.Println("array after sorting", sort.IntsAreSorted(myArr))

	//* how to remove a value from slice based on index

	var bikes = []string{"Ducati Panigale v4", "Aprilia Rsv4", "Yamaha R1m", "Kawasaki zx10r", "Honda cbr fireblade RRR", "Mv Agusta f4"}
	var indexToRemove int = 3
	//* :indexToRemove slices from 0 to 2 , indexToRemove+1: slices from 4 to 5 , ... similar to spread operator
	//* append method accepts multiple arguments
	bikes = append(bikes[:indexToRemove], bikes[indexToRemove+1:]...)
	fmt.Println(bikes)
}
