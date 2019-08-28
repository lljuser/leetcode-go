package condition

import "fmt"

func doWhile() {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println("hello! ", i)
		i++
	}
}
