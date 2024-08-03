// Closure: Binds the parent environment variables with the embedded environment
package main

import "fmt"

func activateGiftCard() func(int) int {
	balance := 100
	return func(debitAmount int) int {
		balance -= debitAmount
		return balance
	}
}

func main() {
	useGiftCard1 := activateGiftCard()
	fmt.Println(useGiftCard1(25)) // 75

	useGiftCard2 := activateGiftCard()
	fmt.Println(useGiftCard2(50)) // 50
}
