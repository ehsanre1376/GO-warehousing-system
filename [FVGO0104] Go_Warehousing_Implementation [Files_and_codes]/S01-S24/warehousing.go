package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type a struct {
	user string
	pass string
}
type b struct {
	name      string
	inventory int
	price     int
}
type c struct {
	name      string
	last_name string
	id        int
	Money     int
}
type d struct {
	id     int
	name   string
	number int
	check  bool
}

var staff [1000]a
var warehouse [1000]b
var customer [1000]c
var transaction [1000]d

func read_staff() {
	file1, err := os.Open("staff.txt")
	if err != nil {
		fmt.Print(err)
	}
	data1 := bufio.NewScanner(file1)
	data1.Split(bufio.ScanLines)
	var s1 []string
	for data1.Scan() {
		s1 = append(s1, data1.Text())
	}
	for i := 0; i < len(s1)/2; i++ {
		staff[i].user = s1[i*2]
		staff[i].pass = s1[(i*2)+1]
	}
}
func read_warehouse() {
	file2, err := os.Open("warehouse.txt")
	if err != nil {
		fmt.Print(err)
	}
	data2 := bufio.NewScanner(file2)
	data2.Split(bufio.ScanLines)
	var s2 []string
	for data2.Scan() {
		s2 = append(s2, data2.Text())
	}
	for i := 0; i < len(s2)/3; i++ {
		warehouse[i].name = s2[i*3]
		warehouse[i].inventory, err = strconv.Atoi(s2[(i*3)+1])
		warehouse[i].price, err = strconv.Atoi(s2[(i*3)+2])
	}
}
func read_customer() {
	file3, err := os.Open("customer.txt")
	if err != nil {
		fmt.Print(err)
	}
	data3 := bufio.NewScanner(file3)
	data3.Split(bufio.ScanLines)
	var s3 []string
	for data3.Scan() {
		s3 = append(s3, data3.Text())
	}
	for i := 0; i < len(s3)/4; i++ {
		customer[i].name = s3[i*4]
		customer[i].last_name = s3[(i*4)+1]
		customer[i].id, err = strconv.Atoi(s3[(i*4)+2])
		customer[i].Money, err = strconv.Atoi(s3[(i*4)+3])
	}
}
func read_transaction() {
	file4, err := os.Open("transaction.txt")
	if err != nil {
		fmt.Print(err)
	}
	data4 := bufio.NewScanner(file4)
	data4.Split(bufio.ScanLines)
	var s4 []string
	for data4.Scan() {
		s4 = append(s4, data4.Text())
	}
	for i := 0; i < len(s4)/4; i++ {
		transaction[i].id, err = strconv.Atoi(s4[i*4])
		transaction[i].name = s4[(i*4)+1]
		transaction[i].number, err = strconv.Atoi(s4[(i*4)+2])
		transaction[i].check, err = strconv.ParseBool(s4[(i*4)+3])
	}
}
func write_warehouse() {
	file, err := os.Create("warehouse.txt")
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < 1000; i++ {
		if warehouse[i].name != "" && warehouse[i].inventory != 0 && warehouse[i].price != 0 {
			file.WriteString(warehouse[i].name)
			file.WriteString("\n")
			file.WriteString(strconv.Itoa(warehouse[i].inventory))
			file.WriteString("\n")
			file.WriteString(strconv.Itoa(warehouse[i].price))
			file.WriteString("\n")
		}
	}
}
func write_customer() {
	file, err := os.Create("customer.txt")
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < 1000; i++ {
		if customer[i].name != "" && customer[i].last_name != "" && customer[i].id != 0 && customer[i].Money != 0 {
			file.WriteString(customer[i].name)
			file.WriteString("\n")
			file.WriteString(customer[i].last_name)
			file.WriteString("\n")
			file.WriteString(strconv.Itoa(customer[i].id))
			file.WriteString("\n")
			file.WriteString(strconv.Itoa(customer[i].Money))
			file.WriteString("\n")
		}
	}
}
func write_transaction() {
	file, err := os.Create("transaction.txt")
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < 1000; i++ {
		if transaction[i].name != "" && transaction[i].id != 0 && transaction[i].number != 0 {
			file.WriteString(strconv.Itoa(transaction[i].id))
			file.WriteString("\n")
			file.WriteString(transaction[i].name)
			file.WriteString("\n")
			file.WriteString(strconv.Itoa(transaction[i].number))
			file.WriteString("\n")
			file.WriteString(strconv.FormatBool(transaction[i].check))
			file.WriteString("\n")
		}
	}
}
func main() {
	var kala = make(map[string]int)
	var moshtari = make(map[int]int)
	read_staff()
	read_warehouse()
	read_customer()
	read_transaction()
	var s, ss string
	sw := 0
	flag := 0
	for sw < 10 {
		fmt.Println("username va password ra vared konid")
		fmt.Scan(&s, &ss)
		for i := 0; i < 1000; i++ {
			if s == staff[i].user && ss == staff[i].pass {
				flag++
				code := staff[i].user[0]
				if code == '1' {
					var s int
					for sw < 5 {
						fmt.Println("1.moshahede anbar")
						fmt.Println("2.moshahede moshtari")
						fmt.Println("3.por sood tarin moshtari")
						fmt.Println("4.por sood tarin kala")
						fmt.Println("5.khoroj")
						fmt.Scan(&s)
						if s == 1 {
							for i := 0; i < 1000; i++ {
								if warehouse[i].name != "" && warehouse[i].inventory != 0 && warehouse[i].price != 0 {
									fmt.Println("nam kala barabar ast ba: " + warehouse[i].name)
									fmt.Print("tedad kala barabar ast ba: ")
									fmt.Println(warehouse[i].inventory)
									fmt.Print("gheymat har vahed kala barabar ast ba: ")
									fmt.Println(warehouse[i].price)
									fmt.Println()
									fmt.Println()
								}
							}
						} else if s == 2 {
							for i := 0; i < 1000; i++ {
								if customer[i].name != "" && customer[i].last_name != "" && customer[i].id != 0 && customer[i].Money != 0 {
									fmt.Println("nam moshtari barabar ast ba: " + customer[i].name)
									fmt.Println("nam khanevadegi moshtari barabar ast ba: " + customer[i].last_name)
									fmt.Print("id moshtari barabar ast ba: ")
									fmt.Println(customer[i].id)
									fmt.Print("mojoodi hesab moshtari barabar ast ba: ")
									fmt.Println(customer[i].Money)
									fmt.Println()
									fmt.Println()
								}
							}
						} else if s == 3 {
							fmt.Println(moshtari)

						} else if s == 4 {
							fmt.Println(kala)

						} else if s == 5 {
							break

						} else {
							fmt.Print("error")
						}
					}

				} else if code == '2' {
					var s int
					for sw < 5 {
						fmt.Println("1.moshahede anbar")
						fmt.Println("2.moshahede moshtari")
						fmt.Println("3.sabt tarakonesh")
						fmt.Println("4.khoroj")
						fmt.Scan(&s)
						if s == 1 {
							for i := 0; i < 1000; i++ {
								if warehouse[i].name != "" && warehouse[i].inventory != 0 && warehouse[i].price != 0 {
									fmt.Println("nam kala barabar ast ba: " + warehouse[i].name)
									fmt.Print("tedad kala barabar ast ba: ")
									fmt.Println(warehouse[i].inventory)
									fmt.Print("gheymat har vahed kala barabar ast ba: ")
									fmt.Println(warehouse[i].price)
									fmt.Println()
									fmt.Println()
								}
							}
						} else if s == 2 {
							for i := 0; i < 1000; i++ {
								if customer[i].name != "" && customer[i].last_name != "" && customer[i].id != 0 && customer[i].Money != 0 {
									fmt.Println("nam moshtari barabar ast ba: " + customer[i].name)
									fmt.Println("nam khanevadegi moshtari barabar ast ba: " + customer[i].last_name)
									fmt.Print("id moshtari barabar ast ba: ")
									fmt.Println(customer[i].id)
									fmt.Print("mojoodi hesab moshtari barabar ast ba: ")
									fmt.Println(customer[i].Money)
									fmt.Println()
									fmt.Println()
								}
							}
						} else if s == 3 {
							for i := 0; i < 1000; i++ {
								if transaction[i].id != 0 && transaction[i].name != "" && transaction[i].number != 0 {
									check1 := 0
									for j := 0; j < 1000; j++ {
										if transaction[i].id == customer[j].id {
											check1++
											check2 := 0
											for k := 0; k < 1000; k++ {
												if transaction[i].name == warehouse[k].name {
													check2++
													if transaction[i].number <= warehouse[k].inventory {
														if (customer[j].Money - (transaction[i].number * warehouse[k].price)) >= -200000 {
															if transaction[i].check == false {
																kala[transaction[i].name] += (transaction[i].number * warehouse[k].price)
																moshtari[transaction[i].id] += (transaction[i].number * warehouse[k].price)
																warehouse[k].inventory -= transaction[i].number
																customer[j].Money -= (transaction[i].number * warehouse[k].price)
																transaction[i].check = true
															}
														} else {
															fmt.Printf("mojodi hesab moshtari dar tarakomesh %v kafi nist", i+1)
														}
													} else {
														fmt.Printf("mojodi anbar dar tarakomesh %v kafi nist", i+1)
													}
												}

											}
											if check2 == 0 {
												fmt.Printf("kalayi ba in nam dar tarakomesh %v vojod nadarad", i+1)
											}
										}
									}
									if check1 == 0 {
										fmt.Printf("moshtari ba in id dar tarakomesh %v vojod nadarad", i+1)
									}
								}

							}
							write_customer()
							write_warehouse()
							write_transaction()
						} else if s == 4 {
							break
						} else {
							fmt.Print("error")
						}
					}
				} else {
					var s int
					for sw < 5 {
						fmt.Println("1.moshahede anbar")
						fmt.Println("2.afzoodan kala")
						fmt.Println("3.khoroj")
						fmt.Scan(&s)
						if s == 1 {
							for i := 0; i < 1000; i++ {
								if warehouse[i].name != "" && warehouse[i].inventory != 0 && warehouse[i].price != 0 {
									fmt.Println("nam kala barabar ast ba: " + warehouse[i].name)
									fmt.Print("tedad kala barabar ast ba: ")
									fmt.Println(warehouse[i].inventory)
									fmt.Print("gheymat har vahed kala barabar ast ba: ")
									fmt.Println(warehouse[i].price)
									fmt.Println()
									fmt.Println()
								}
							}
						} else if s == 2 {
							var a string
							var aa, aaa int
							fmt.Print("nam kala ra vared konid : ")
							fmt.Scan(&a)
							fmt.Print("tedad kala ra vared konid : ")
							fmt.Scan(&aa)
							fmt.Print("gheymat har vahed kala ra vared konid : ")
							fmt.Scan(&aaa)
							for i := 0; i < 1000; i++ {
								if warehouse[i].name == "" && warehouse[i].inventory == 0 && warehouse[i].price == 0 {
									fmt.Println(i)
									warehouse[i].name = a
									warehouse[i].inventory = aa
									warehouse[i].price = aaa
									break
								}
							}
							write_warehouse()
						} else if s == 3 {
							break
						} else {
							fmt.Print("error")
						}
					}
				}
			}
		}
		if flag == 0 {
			fmt.Println("not found")
		}
	}

}

//go run warehousing.go
