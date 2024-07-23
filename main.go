package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func deposit() {
	fmt.Println("Deposite amount...")
}

func withDraw() {
	fmt.Println("Withdrawing amount...")
}

func main() {

	var list = []string{"012671", "aavashh"}

	fmt.Println("Enter account number")
	accountHolder := bufio.NewReader(os.Stdin)
	holder, _ := accountHolder.ReadString('\n')

	value := strings.TrimSpace(holder)

	if value == list[0] {
		fmt.Println("checking")
	}

}
