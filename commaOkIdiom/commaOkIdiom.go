package commaOkIdiom

import "fmt"

var myMap = map[int]int {
	0: 1,
	1: 2,
	3: 4,
}

func ValExists()  {
	fmt.Println(myMap)

	if val, exists := myMap[1]; exists { // [1] is the KEY
		fmt.Println("val:", val) // Value of val: is the VALUE
		fmt.Println("exists:", exists) // Value of exists: TRUE
		fmt.Println("deleting", val)
		delete(myMap, 1)
	}  else {
		fmt.Println("The value", val, "didn't exist")
	}

}

