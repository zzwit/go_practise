package main

import "fmt"

func main()  {
	seasons := [] string{"Spring", "Summer", "Autumn", "Winter"}
	for index,val := range seasons {
		fmt.Printf("Season %d is: %s\n", index, val)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	for ix := range seasons {
		fmt.Printf("%d \n", ix)
	}

	for row := range screen {
		for column := range screen[row] {
			screen[row][column] = 1
		}
	}

}
