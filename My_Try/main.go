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

// main is the entry point of the program. It opens the "staff.txt" file,
// reads its contents line by line, and stores each line in a slice of strings.
// If there is an error opening the file, the program will panic. Finally, it
// prints the contents of the slice and ensures the file is closed before the
// program exits.
func main() {
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
	for _, staff := range staffLines {
		fmt.Println(staff)
	}
	/////////////////////////////////////////////
	var warehouseList [4]Warehouse
	//warehouseList = append(warehouseList, Warehouse{name: "A4", inventory: 100, price: 1000})
	// Open the "warehouse.txt" file for reading.
	// If there is an error opening the file, the program will panic.
	warehouseFile, err := os.Open("warehouse.txt")
	if err != nil {
		panic(err)
	}
	defer staffFile.Close()
	datawarehouse := bufio.NewScanner(warehouseFile)
	var warehouseLines []string
	for datawarehouse.Scan() {
		warehouseLines = append(warehouseLines, datawarehouse.Text())
	}
	fmt.Println(warehouseLines)
	fmt.Println(warehouseList)
	// since we have I*3 we cant use this code below
	//First way
	// for i := 0; i < len(warehouseLines); i += 3 {

	// inventory, err := strconv.Atoi(warehouseLines[(i)+1])
	// if err != nil {
	// 	panic(err)
	// }
	// price, err := strconv.Atoi(warehouseLines[(i)+2])
	// if err != nil {
	// 	panic(err)
	// }
	// warehouseList = append(warehouseList, Warehouse{name: warehouseLines[i], inventory: inventory, price: price})
	//Second way
	for i := 0; i < len(warehouseLines)/3; i++ {
		warehouseList[i].inventory, err = strconv.Atoi(warehouseLines[(3*i)+1])
		warehouseList[i].price, err = strconv.Atoi(warehouseLines[(3*i)+2])
		warehouseList[i].name = warehouseLines[(3 * i)]
	}

	// fmt.Println("122222222222222222222222")
	// fmt.Println(warehouseList[0])
	// // for _, warehouse := range warehouseList {
	// 	fmt.Println(warehouse)
	// }
	// for i, warehouse := range warehouseLines {
	// 	fmt.Println(i, warehouse)
	// }
}
