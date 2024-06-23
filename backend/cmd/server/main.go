package main

import "fmt"

func Run() error {
	fmt.Println("Starting Our Backend API")

	return nil
}

func main() {
	fmt.Println("Piko Blog API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}