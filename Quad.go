package main

import (
	"github.com/01-edu/z01"
)

func main() {
	QuadA(3, 5)
}

func QuadA(x_size int, y_size int) {
	// these are the definitions. add symbols here as you desire
	horizontal_rune := 'x'
	vertical_rune := 'x'
	corner_rune_topleft := 'o'
	corner_rune_bottomleft := 'o'
	corner_rune_topright := 'o'
	corner_rune_bottomright := 'o'

	// create the matrix with the specified dimensions
	// the matrix := make() creates each row of the matrix, and the for loop
	// creates the individual elements within each row. (this is how we get a 2d matrix of x and y.)
	matrix := make([][]rune, x_size)

	for i := 0; i < x_size; i++ {
		matrix[i] = make([]rune, y_size)
	}

	// this goes over the matrix. the first for loop goes over the x row. the second for loop does the y row.
	// then it assigns the position by setting getting matrix[x_index][y_index].
	for x_index := 0; x_index < x_size; x_index++ {
		for y_index := 0; y_index < y_size; y_index++ {
			matrix[x_index][y_index] = ' '

			if x_index == 0 || x_index == x_size-1 { // <-- this checks whether x index is 0 or max when for looping
				matrix[x_index][y_index] = horizontal_rune
			} else if y_index == 0 || y_index == y_size-1 { // <--- thischecks whether y index is 0 or max when for looping
				matrix[x_index][y_index] = vertical_rune
			}

			// set the top left corner rune.
			// this if checks whether x coordinate and y coordinate is 0.
			if x_index == 0 && y_index == 0 {
				matrix[x_index][y_index] = corner_rune_topleft
			}

			// set the top right corner rune
			// this if checks whether x coordinate is max and y coordinate is 0.
			// NOTE :: (the reason we use -1 is because the count starts from 0, so coords are x0-y4 not x1-y5)
			// i'm also too lazy to write separate functions so they stay as x_size-1 and y_size-1. not clean code
			if x_index == x_size-1 && y_index == 0 {
				matrix[x_index][y_index] = corner_rune_topright
			}

			// set the bottom left corner rune
			// this if checks whether x coordinate is 0 and y coordinate is max.
			if x_index == 0 && y_index == y_size-1 {
				matrix[x_index][y_index] = corner_rune_bottomleft
			}

			// set the bottom right corner rune
			// this if checks whether x coordinate is max and y coordinate is max.
			if x_index == x_size-1 && y_index == y_size-1 {
				matrix[x_index][y_index] = corner_rune_bottomright
			}
		}
	}

	// this is just to print it out. it goes over the table for the row
	// and checks what the [x][y] matrix is, and prints it.
	// so if it was [5][2] matrix, it would get x 5 and y 2 of the matrix and print the symbol.
	for _, row := range matrix {
		for _, r := range row {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}

}
