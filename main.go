// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program calculates your actual pay and what the government takes

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	var hoursWorked float64
	var hourlyRate float64

	// input
	fmt.Println("This program gets a user's actual pay and what the government takes.")
	fmt.Println()
	fmt.Print("Enter the hours you've worked: ")
	fmt.Scanln(&hoursWorked)
	fmt.Print("Enter your hourly rate: ")
	fmt.Scanln(&hourlyRate)

	// process
	var pay = (hoursWorked * hourlyRate) * 0.82
	var government = (hoursWorked * hourlyRate) * 0.18

	// This function displays currency
	accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println("The pay will be:", accountingFormater.FormatMoney(pay))
	accountingFormater2 := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println("The government will take:", accountingFormater2.FormatMoney(government))
}
