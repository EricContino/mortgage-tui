package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func commandNew(cfg *config, args ...string) error {
	reader := bufio.NewScanner(os.Stdin)

	valid := false
	for !valid {
		fmt.Print("Enter Original Mortgage Balance > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			fmt.Println("Nothing entered")
			continue
		}

		num, err := strconv.ParseFloat(words[0], 32)
		if err != nil {
			fmt.Printf("Error converting string to float: %v\n", err)
			continue
		}
		cfg.origBalance = num
		valid = true
	}

	valid = false
	for !valid {
		fmt.Print("Enter Interest Rate(in percentage) > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			fmt.Println("Nothing entered")
			continue
		}

		num, err := strconv.ParseFloat(words[0], 32)
		if err != nil {
			fmt.Printf("Error converting string to float: %v\n", err)
			continue
		}

		cfg.interestRate = num / 100
		valid = true
	}

	valid = false
	for !valid {
		fmt.Print("Enter Loan Durtation in months > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			fmt.Println("Nothing entered")
			continue
		}

		num, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			continue
		}

		cfg.durationinMonths = num
		valid = true
	}

	valid = false
	for !valid {
		fmt.Print("Enter Montly Escrow > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			fmt.Println("Nothing entered")
			continue
		}

		num, err := strconv.ParseFloat(words[0], 32)
		if err != nil {
			fmt.Printf("Error converting string to float: %v\n", err)
			continue
		}

		cfg.monthlyEscrow = num
		valid = true
	}

	return nil

}
