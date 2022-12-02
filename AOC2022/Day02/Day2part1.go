package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	// Read input file
	input,_:=os.Open("input1.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	// A = X Rock 1
	// B = Y Paper 2
	// C = Z Scissors 3
	// WIN = +6
	// LOSE = +0
	// DRAW = +3

	var score int
	scores := map[string]int{"B X":1, "C Y":2, "A Z":3, "A X":4, "B Y":5, "C Z":6, "C X":7, "A Y":8, "B Z":9 }

	for sc.Scan(){
		score += scores[sc.Text()]
	}
	fmt.Println(score)
}