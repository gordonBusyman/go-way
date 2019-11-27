package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// checkPizza()
	// checkFunctionVariable()
	// checkDefer()
	// checkFunctionGenerator()

}

// multiplyInts function
func multiplyInts(x, y int) int {
	return x * y
}

// multiplyIntsV2 function
func multiplyIntsV2(x, y int) (z int) {
	z = x * y
	return
}

// getPositiveMultiplierV1 function
func getPositiveMultiplierV1(x, y int) (z int, err error) {
	if x < 0 || y < 0 {
		err = fmt.Errorf("works only with positive numbers")
	}
	z = x * y
	return
}

// getPositiveMultiplierV2 function
func getPositiveMultiplierV2(x, y int) (int, error) {
	var z int
	if x < 0 || y < 0 {
		err := fmt.Errorf("works only with positive numbers")
		return z, err
	}
	z = x * y
	return z, nil
}

// getPositiveMultiplierV3 function
func getPositiveMultiplierV3(x, y int) (functionToReturn func() int, err error) {
	if x < 0 || y < 0 {
		err = fmt.Errorf("works only with positive numbers")
	}
	functionToReturn = func() int {
		return x * y
	}
	return
}

func f() func() int {
	x := 2

	return func() int {
		x += x
		return x
	}
}

// getFileSizeInBytes function
func getFileSizeInBytes(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	var bytesReaded []byte

	for {
		byteReaded, err := reader.ReadByte()
		if err != nil {
			break
		}
		buffer.WriteByte(byteReaded)
		bytesReaded = append(bytesReaded, byteReaded)
	}
	fmt.Println(len(bytesReaded))
}

// sum function
func sum(vals ...int) (total int) {
	for _, v := range vals {
		total += v
	}
	return
}

// Pizza ...
type Pizza struct {
	IsPepperoni  bool
	HasPineapple bool
	Price        int
}

func callByValue(p Pizza) {
	p.Price = 666
}

func callByReference(p *Pizza) {
	p.Price = 666
}

func checkPizza() {
	pizza := Pizza{
		IsPepperoni:  false,
		HasPineapple: true,
		Price:        12,
	}

	p := &pizza

	fmt.Printf("\nPizza before callByValue: %v\n", p)
	callByValue(pizza)
	fmt.Printf("Pizza after callByValue: %v\n\n-------------------------\n\n", p)

	fmt.Printf("Pizza before callByReference: %v\n", p)
	callByReference(p)
	fmt.Printf("Pizza after callByReference: %v\n\n", p)

}

func checkFunctionVariable() {
	var sayHello = func(name string) {
		fmt.Println("Hello,", name)
	}

	sayHello("John")
}

func checkDefer() {
	fmt.Println("\nStart Demo:")
	defer fmt.Println("End Demo (1)")
	defer fmt.Println("End Demo (2)")

	fmt.Println("Do stuff...")
}

func checkFunctionGenerator() {
	firstArg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	secondArg, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	actualMultiplier, err := getPositiveMultiplierV3(firstArg, secondArg)
	if err != nil {
		fmt.Println(err)
	}
	res := actualMultiplier()
	fmt.Printf("\tResult: %v (%[1]T)\n", res)
}
