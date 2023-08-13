package main

import (
	"fmt"
)

func main() {
	Quad(5, 5) // example
}

func Quad(x_size int, y_size int) {
	// these are the definitions. (defining symbols so we can replace them easier)
	horizontal_rune := '*'
	vertical_rune := '*'

	corner_rune_topleft := 'A'
	corner_rune_bottomleft := 'C'
	corner_rune_topright := 'B'
	corner_rune_bottomright := 'D'

	// create the matrix with the specified dimensions
	// the matrix := make() creates each row of the matrix, and the for loop
	// creates the individual elements within each row. (this is how we get a 2d matrix of x and y)
	matrix := make([][]rune, y_size)

	for i := 0; i < y_size; i++ {
		matrix[i] = make([]rune, x_size)
	}

	for i := 0; i < y_size; i++ {
		for j := 0; j < x_size; j++ {
			// fill all slots with a character, we'll overwrite these later with the appropriate symbol
			matrix[i][j] = ' '

			// checking sides and overwriting them with the horizontal / vertical runes
			if i == 0 || i == y_size-1 {
				matrix[i][j] = horizontal_rune
			} else if j == 0 || j == x_size-1 {
				matrix[i][j] = vertical_rune
			}

			// set the top left corner rune.
			// this if checks whether x coordinate and y coordinate is 0
			if i == 0 && j == 0 {
				matrix[i][j] = corner_rune_topleft
			}

			// NOTE :: (the reason we use -1 is because the count starts from 0, so coords are x0-y4 not x1-y5)
			// i'm also too lazy to write separate functions so they stay as x_size-1 and y_size-1. not clean code

			// set the top right corner rune
			// this if checks whether x coordinate is max and y coordinate is 0.
			// EXCLUDE IF ITS TOP LEFT CORNER. WE WANT IT!
			if i == 0 && j == x_size-1 && j != 0 {
				matrix[i][j] = corner_rune_topright
			}

			// set the bottom left corner rune
			// this if checks whether x coordinate is 0 and y coordinate is max.
			// EXCLUDE IF ITS TOP LEFT CORNER. WE WANT IT!
			if i == y_size-1 && j == 0 && i != 0 {
				matrix[i][j] = corner_rune_bottomleft
			}

			// set the bottom right corner rune
			// this if checks whether x coordinate is max and y coordinate is max.
			// EXCLUDE IF ITS BOTTOM LEFT /AND/ TOP RIGHT CORNER. WE WANT THOSE!
			if i == y_size-1 && j == x_size-1 && i != 0 && j != 0 {
				matrix[i][j] = corner_rune_bottomright
			}
		}
	}

	// this is just to print it out. it goes over the table for the row
	// and checks what the [x][y] matrix is, and prints it.
	// so if it was [5][3] matrix, it would get x 5 and y 3 of the matrix and print the symbol.
	for _, row := range matrix {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Printf("\n")
	}
}
