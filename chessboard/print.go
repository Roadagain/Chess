package chessboard

import "fmt"

func PrintChessboard() {
	for i := 0; i < 8; i++ {
		fmt.Print("+")
		for j := 0; j < 8; j++ {
			fmt.Print("-+")
		}
		fmt.Println()
		fmt.Print("|")
		for j := 0; j < 8; j++ {
			fmt.Print(" |")
		}
		fmt.Println()
	}
	fmt.Print("+")
	for i := 0; i < 8; i++ {
		fmt.Print("-+")
	}
	fmt.Println()
}
