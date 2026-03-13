package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/EricContino/mortgage-tui/internal/amortization"
)

func commandPrint(cfg *config, args ...string) error {
	// Detect terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default to 80 if error
	}

	monthlyPayment := amortization.CalcMonthlyPayment(cfg.origBalance, cfg.interestRate, cfg.durationinMonths)

	seperator := strings.Repeat("=", width)
	fmt.Println(seperator)
	fmt.Println("Original Balance: ", cfg.origBalance)
	fmt.Printf("Interest Rate: %g%%\n", cfg.interestRate*100)
	fmt.Println("Duration in Months: ", cfg.durationinMonths)
	fmt.Println("Escrow: ", cfg.monthlyEscrow)
	fmt.Println("Monthly Payment: ", monthlyPayment)
	fmt.Println(seperator)
	return nil
}
