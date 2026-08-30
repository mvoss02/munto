package main

import (
	"fmt"
	"os"

	"github.com/mvoss02/munto/internal/enablebanking"
	"github.com/mvoss02/munto/internal/money"
)

func main() {
	appId, ok := os.LookupEnv("MUNTO_EB_APP_ID")
	if !ok {
		fmt.Fprintln(os.Stderr, "MUNTO_EB_APP_ID not set")
		os.Exit(1)
	}
	fmt.Printf("App id is: %s \n", appId)

	keyPath, ok := os.LookupEnv("MUNTO_EB_KEY_PATH")
	if !ok {
		fmt.Fprintln(os.Stderr, "MUNTO_EB_KEY_PATH not set")
		os.Exit(1)
	}
	fmt.Printf("EB key path is: %s \n", keyPath)

	key, err := enablebanking.LoadKey(keyPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "load key:", err)
		os.Exit(1)
	}
	fmt.Printf("Key is: %d \n", key.N.BitLen())

	formattedMoney := money.Format(500, "EUR")
	fmt.Printf("Formatted money amount: %s \n", formattedMoney)
}
