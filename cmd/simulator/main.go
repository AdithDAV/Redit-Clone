package main

import (
	"Go_project/internal/client"
	"fmt"
)

const (
	serverAddr = "localhost:50051" // Address of the gRPC server
	userCount  = 1000              // Number of users to simulate
	subredditCount = 50            // Number of subreddits to simulate
)

func main() {
	fmt.Println("Reddit Clone Simulator Starting")

	// Initialize the simulator with the gRPC server address
	simulator := client.NewSimulator(serverAddr, userCount, subredditCount)

	fmt.Println("Starting simulation...")
	simulator.Run()

	fmt.Println("Simulation completed")
}
