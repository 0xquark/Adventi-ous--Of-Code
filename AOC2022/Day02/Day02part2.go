package main
import (
"bufio"
"fmt"
"os"
)
func main() {
input, _ := os.Open("input1.txt")
defer input.Close()
sc := bufio.NewScanner(input)
// A = X Rock 1
// B = Y Paper 2
// C = Z Scissors 3
// WIN = +6
// LOSE = +0
// DRAW = +3

var score int
scores := map[string]int{"B X":1, "C X":2, "A X":3, "A Y":4,"B Y":5,"C Y":6, "C Z":7, "A Z":8, "B Z":9}
for sc.Scan() {
score += scores[sc.Text()]
}
fmt.Println(score)
}