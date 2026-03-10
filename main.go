package main

func main() {
	cfg := &config{
		origBalance:      0.0,
		currBalance:      0.0,
		interestRate:     0.0,
		monthlyPayment:   0.0,
		escrow:           0.0,
		durationinMonths: 0,
	}
	startRepl(cfg)
}
