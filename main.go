package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("create a new bill: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	b := newBill(name)

	fmt.Println("created a bill - ", b.name)

	return b
}

func main() {
	myBill := createill()

	fmt.Println(myBill)

}

// func main() {
// 	mybill := newBill("mario's bill")
// 	mybill.addItem("veger", 10.3)
// 	mybill.addItem("veerg", 70.3)
// 	mybill.addItem("vrg", 15.3)
// 	mybill.updateTip(10)
// 	fmt.Println(mybill.format())
// }
