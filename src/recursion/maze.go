package recursion

import "fmt"

func ConstructMaze(size int) [][]string {
	rows := [][]string{}
	for i := 0; i < size; i++ {
		cols := make([]string, size)
		for j := 0; j < size; j++ {
			if i == j {
				cols[j] = "X"
			}
		}
		rows = append(rows, cols)
	}
	return rows
}

func GetMaze() [][]string {
	maze := [][]string{
		[]string{"X", "X", "X", "X", "X", "X", "X"},
		[]string{" ", "X", "X", " ", " ", " ", "X"},
		[]string{" ", " ", "X", " ", "X", " ", "X"},
		[]string{"X", " ", " ", " ", "X", " ", " "},
		[]string{"X", "X", "X", " ", "X", "X", "X"},
		[]string{"X", "X", "X", " ", "X", "X", "X"},
		[]string{"X", "X", "X", "X", "X", "X", "X"},
	}
	return maze
}

func PrintMaze(maze [][]string) {
	for _, r := range maze {
		for _, c := range r {
			fmt.Printf("%s ", c)
		}
		fmt.Printf("\n")
	}
	fmt.Println("==============")
}

func IsValidCell(size, row, column int) bool {
	if row >= size || column >= size {
		return false
	}
	if row < 0 || column < 0 {
		return false
	}
	return true
}

func IsBorder(size, row, column int) bool {
	if row == size-1 || column == size-1 || row == 0 || column == 0 {
		return true
	}
	return false
}

func EscapeMaze(maze [][]string, posx, posy int) bool {
	if !IsValidCell(len(maze[0]), posx, posy) {
		return false
	}
	if maze[posx][posy] == "X" {
		return false
	}
	if maze[posx][posy] == "V" || maze[posx][posy] == "D" {
		return false
	}
	if maze[posx][posy] == " " && IsBorder(len(maze[0]), posx, posy) {
		maze[posx][posy] = ":)"
		PrintMaze(maze)
		return true
	}
	maze[posx][posy] = "V"
	found := EscapeMaze(maze, posx-1, posy) ||
		EscapeMaze(maze, posx+1, posy) ||
		EscapeMaze(maze, posx, posy-1) ||
		EscapeMaze(maze, posx, posy+1)
	if found {
		PrintMaze(maze)
		return true
	} else {
		maze[posx][posy] = "D"
	}
	return false
}
