package main

import (
	"fmt"
	"log"

	captcha "../src/golang"
)

func main() {
	token, err := captcha.SolveVVCaptcha("https://uk-umg.com/um-forms/48709-1272284.html")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("VVCaptcha token: %s\n", token)
}