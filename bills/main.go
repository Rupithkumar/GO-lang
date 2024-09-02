package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}
func createbill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Bill created as: ", b.name)
	return b
}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a-add item, s-save bill, t-add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItems(name, p)
		fmt.Println("Item added -", name, p)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("Bill saved!", b.name)
		fmt.Println(b.format())
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added -", tip)
		promptOptions(b)
	default:
		fmt.Println("Enter a valid option")
		promptOptions(b)
	}
}

func main() {
	mybill := createbill()
	promptOptions(mybill)
}
