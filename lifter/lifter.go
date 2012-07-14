package main

import (
	"fmt"
	"os"
)

func main() {
	m, err := ReadMap(os.Stdin)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("\n\nSTARTING")

	for _, v := range m.m {
		fmt.Printf("%s\n", v)
	}
	fmt.Printf("Lift: %d, %d\n", m.Lift.x, m.Lift.y)
	fmt.Printf("Robot: %d, %d\n", m.Robot.x, m.Lift.y)
	fmt.Printf("Lambdas: %d\n", m.Lambdas)
	fmt.Printf("Water: %d\n", m.Water)
	fmt.Printf("Flooding: %d\n", m.Flooding)
	fmt.Printf("Waterproof: %d\n", m.Waterproof)
	fmt.Printf("Complete: %d\n", m.Complete)
	fmt.Printf("Steps: %d\n", m.Steps)

	test := []byte("DDDLLLLLLURRRRRRRRRRRRDDDDDDDLLLLLLLLLLLDDDRRRRRRRRRRRD")
	for _, c := range test {
		m.Step(c)
	}

	fmt.Println("\n\nENDING")

	for _, v := range m.m {
		fmt.Printf("%s\n", v)
	}

	fmt.Printf("Lift: %d, %d\n", m.Lift.x, m.Lift.y)
	fmt.Printf("Robot: %d, %d\n", m.Robot.x, m.Lift.y)
	fmt.Printf("Lambdas: %d\n", m.Lambdas)
	fmt.Printf("Water: %d\n", m.Water)
	fmt.Printf("Flooding: %d\n", m.Flooding)
	fmt.Printf("Waterproof: %d\n", m.Waterproof)
	fmt.Printf("Complete: %d\n", m.Complete)
	fmt.Printf("Steps: %d\n", m.Steps)
}
