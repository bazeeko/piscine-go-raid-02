package main

import (
	"fmt"
	"os"
	"github.com/01-edu/z01"
)

func checkCorrect(args []string) bool {
	if len(args) != 10 {
		return false
	}
	for i := 1; i <= 9; i++ {
		str := args[i]
		if len(str) != 9 {
			return false
		}
		for j := range str {
			if !(rune(str[j]) >= '1' && rune(str[j]) <= '9') && rune(str[j]) != '.' {
				return false
			}
			for k := j + 1; k <= 8; k++ {
				if rune(str[j]) == rune(str[k]) && rune(str[j]) != '.' {
					return false
				}
			}
		}
	}
	return true
}

func main(){
	args := os.Args
	if checkCorrect(args) == false {
		fmt.Println("Error")
		return
	}

	runes := make([][]rune, 9)
	for i := 1; i < len(args); i++ {
		runes[i-1] = []rune(args[i])
	}
	if solution(runes) {
		for i := range runes {
			for j:= range runes[i] {
				z01.PrintRune(runes[i][j])
				if j<len(runes)-1{
					z01.PrintRune(' ')
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}
}

func solution(runes [][]rune) bool {
	var a, b int
	if !getEmpty(runes, &a, &b) {
		return true
	}
	

	for i:='1';i<='9';i++{
		if check(runes, a, b, i) {
			runes[a][b] = i
			if solution(runes) {
				return true
			}
			runes[a][b]='.'
		}
	}
	return false
}

func getEmpty(runes [][]rune, a *int, b *int) bool {
	for *a=0; *a<9;*a++ {
		for *b=0;*b<9;*b++{
			if runes[*a][*b]=='.' {
				return true
			}
		}
	}
	return false
}

func checkRow(runes [][]rune, a int, num rune) bool {
	for i:=0; i<9; i++ {
		if runes[a][i] == num {
			return true
		}
	}
	return false
}

func checkColumn(runes [][]rune, a int, num rune) bool {
	for i:=0; i<9;i++{
		if runes[i][a] == num {
			return true
		}
	}
	return false
}

func checkSquare(runes [][]rune, a,b int, num rune) bool {
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			if runes[i+a][j+b]==num {
				return true
			}
		}
	}
	return false
}

func check(runes [][]rune, a,b int, num rune) bool {
	return !checkRow(runes, a, num) && !checkColumn(runes, b, num) && !checkSquare(runes, a-a%3, b-b%3, num) && runes[a][b]=='.'
}