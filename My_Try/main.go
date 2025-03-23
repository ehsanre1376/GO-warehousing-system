package main

import (
	"bufio"
	"fmt"
	"os"
)

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











}
