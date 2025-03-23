package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Warehouse struct {
	// Name of the warehouse.
	name string
	// Inventory of the warehouse.
	inventory int
	price     int
}

type Customer struct {
	// Name of the warehouse.
	name string
	// Inventory of the warehouse.
	Finally_name string
	ID           int
	mandeh       int
}

type Transaction struct {
	ID        int
	Name      string
	Inventory int
}

type Staff struct {
	Username string
	Password string
}

// main is the entry point of the program. It opens the "staff.txt" file,
// reads its contents line by line, and stores each line in a slice of strings.
// If there is an error opening the file, the program will panic. Finally, it
// prints the contents of the slice and ensures the file is closed before the
// program exits.

func Read_Files() {
	////////////////////////////////////////////
	// Open the "staff.txt" file for reading.
	// If there is an error opening the file, the program will panic.
	staffFile, err := os.Open("staff.txt")
	if err != nil {
		panic(err)
	}
	defer staffFile.Close()
	dataStaff := bufio.NewScanner(staffFile)
	var staffLines []string
	for dataStaff.Scan() {
		staffLines = append(staffLines, dataStaff.Text())
	}
	fmt.Println(staffLines)

	// Process the staff lines
	var staffList []Staff
	for i := 0; i < len(staffLines); i += 2 {
		staffList = append(staffList, Staff{Username: staffLines[i], Password: staffLines[i+1]})
	}
	fmt.Println(staffList)
	/////////////////////////////////////////////

	// Open the "warehouse.txt" file for reading.
	// If there is an error opening the file, the program will panic.
	warehouseFile, err := os.Open("warehouse.txt")
	if err != nil {
		panic(err)
	}
	defer warehouseFile.Close()
	datawarehouse := bufio.NewScanner(warehouseFile)
	var warehouseLines []string
	for datawarehouse.Scan() {
		warehouseLines = append(warehouseLines, datawarehouse.Text())
	}
	fmt.Println(warehouseLines)

	// since we have I*3 we cant use this code below
	//First way
	var warehouseList []Warehouse
	for i := 0; i < len(warehouseLines); i += 3 {

		inventory, err := strconv.Atoi(warehouseLines[(i)+1])
		if err != nil {
			panic(err)
		}
		price, err := strconv.Atoi(warehouseLines[(i)+2])
		if err != nil {
			panic(err)
		}
		warehouseList = append(warehouseList, Warehouse{name: warehouseLines[i], inventory: inventory, price: price})
	}
	//Second way
	// var warehouseList [100]Warehouse
	// for i := 0; i < len(warehouseLines)/3; i++ {
	// 	warehouseList[i].inventory, err = strconv.Atoi(warehouseLines[(3*i)+1])
	// 	warehouseList[i].price, err = strconv.Atoi(warehouseLines[(3*i)+2])
	// 	warehouseList[i].name = warehouseLines[(3 * i)]
	// }

	// fmt.Println("122222222222222222222222")
	// fmt.Println(warehouseList[0])
	// // for _, warehouse := range warehouseList {
	// 	fmt.Println(warehouse)
	// }
	// for i, warehouse := range warehouseLines {
	// 	fmt.Println(i, warehouse)
	// }
	////////////////////////////////////////////////////////////////////////////
	// Open the "customer.txt" file for reading.
	// If there is an error opening the file, the program will panic.
	customerFile, err := os.Open("customer.txt")
	if err != nil {
		panic(err)
	}
	defer customerFile.Close()
	datacustomer := bufio.NewScanner(customerFile)
	var customerLines []string
	for datacustomer.Scan() {
		customerLines = append(customerLines, datacustomer.Text())
	}
	fmt.Println(customerLines)

	// since we have I*3 we cant use this code below
	//First way
	var customerList []Customer
	for i := 0; i < len(customerLines); i += 4 {

		ID, err := strconv.Atoi(customerLines[(i)+2])
		if err != nil {
			panic(err)
		}
		mandeh, err := strconv.Atoi(customerLines[(i)+3])
		if err != nil {
			panic(err)
		}
		customerList = append(customerList, Customer{name: customerLines[i], Finally_name: customerLines[i+1], ID: ID, mandeh: mandeh})
	}

	////////////////////////////////////////////////////////////////////////////
	// Open the "Transaction.txt" file for reading.
	// If there is an error opening the file, the program will panic.
	transactionFile, err := os.Open("Transaction.txt")
	if err != nil {
		panic(err)
	}
	defer transactionFile.Close()
	dataTransaction := bufio.NewScanner(transactionFile)
	var transactionLines []string
	for dataTransaction.Scan() {
		transactionLines = append(transactionLines, dataTransaction.Text())
	}
	fmt.Println(transactionLines)

	// Process the transaction lines
	var transactionList []Transaction
	for i := 0; i < len(transactionLines); i += 3 {
		ID, err := strconv.Atoi(transactionLines[i])
		if err != nil {
			panic(err)
		}
		Inventory, err := strconv.Atoi(transactionLines[i+2])
		if err != nil {
			panic(err)
		}
		transactionList = append(transactionList, Transaction{ID: ID, Name: transactionLines[i+1], Inventory: Inventory})
	}
	fmt.Println(transactionList)
}

func main() {
	Read_Files()
}
