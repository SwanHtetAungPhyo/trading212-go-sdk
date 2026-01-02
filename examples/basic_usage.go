// Example: Basic usage of the Trading 212 Go SDK
// Run with: go run examples/basic_usage.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	trading212 "github.com/SwanHtetAungPhyo/trading212-go-sdk"
)

func main() {
	// Get API credentials from environment variables
	apiKey := os.Getenv("TRADING212_API_KEY")
	apiSecret := os.Getenv("TRADING212_API_SECRET")
	
	if apiKey == "" || apiSecret == "" {
		log.Fatal("Please set TRADING212_API_KEY and TRADING212_API_SECRET environment variables")
	}

	// Create client for demo environment
	client := trading212.NewClient(trading212.Demo, apiKey, apiSecret)
	ctx := context.Background()

	// Get account summary
	fmt.Println("=== Account Summary ===")
	summary, err := client.GetAccountSummary(ctx)
	if err != nil {
		log.Fatalf("Failed to get account summary: %v", err)
	}

	fmt.Printf("Account ID: %d\n", summary.ID)
	fmt.Printf("Currency: %s\n", summary.Currency)
	fmt.Printf("Free Cash: %.2f\n", summary.Cash.Free)
	fmt.Printf("Invested: %.2f\n", summary.Cash.Invested)
	fmt.Printf("Result: %.2f\n", summary.Cash.Result)
	fmt.Printf("Total: %.2f\n", summary.Cash.Total)

	// Get all positions
	fmt.Println("\n=== Open Positions ===")
	positions, err := client.GetPositions(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to get positions: %v", err)
	}

	if len(positions) == 0 {
		fmt.Println("No open positions")
	} else {
		for _, pos := range positions {
			fmt.Printf("Ticker: %s\n", pos.Ticker)
			fmt.Printf("  Quantity: %.4f\n", pos.Quantity)
			fmt.Printf("  Current Price: %.2f\n", pos.CurrentPrice)
			fmt.Printf("  Average Price: %.2f\n", pos.AveragePrice)
			fmt.Printf("  P&L: %.2f\n", pos.Ppl)
			fmt.Printf("  FX P&L: %.2f\n", pos.FxPpl)
			fmt.Printf("  Pie Quantity: %.4f\n", pos.PieQuantity)
			fmt.Printf("  Max Buy: %.4f\n", pos.MaxBuy)
			fmt.Printf("  Max Sell: %.4f\n", pos.MaxSell)
			fmt.Printf("  Initial Fill Date: %s\n", pos.InitialFillDate.Format("2006-01-02 15:04:05"))
			fmt.Println()
		}
	}

	// Get pending orders
	fmt.Println("=== Pending Orders ===")
	orders, err := client.GetOrders(ctx)
	if err != nil {
		log.Fatalf("Failed to get orders: %v", err)
	}

	if len(orders) == 0 {
		fmt.Println("No pending orders")
	} else {
		for _, order := range orders {
			fmt.Printf("Order ID: %d\n", order.ID)
			fmt.Printf("  Ticker: %s\n", order.Ticker)
			fmt.Printf("  Type: %s\n", order.Type)
			fmt.Printf("  Side: %s\n", order.Side)
			fmt.Printf("  Status: %s\n", order.Status)
			fmt.Printf("  Quantity: %.4f\n", order.Quantity)
			if order.LimitPrice != nil {
				fmt.Printf("  Limit Price: %.2f\n", *order.LimitPrice)
			}
			if order.StopPrice != nil {
				fmt.Printf("  Stop Price: %.2f\n", *order.StopPrice)
			}
			fmt.Printf("  Created: %s\n", order.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Println()
		}
	}

	// Get some instruments
	fmt.Println("=== Sample Instruments ===")
	instruments, err := client.GetInstruments(ctx)
	if err != nil {
		log.Fatalf("Failed to get instruments: %v", err)
	}

	fmt.Printf("Total instruments available: %d\n", len(instruments))
	fmt.Println("First 5 instruments:")
	for i, instrument := range instruments {
		if i >= 5 {
			break
		}
		fmt.Printf("  %s - %s (%s)\n", instrument.Ticker, instrument.Name, instrument.Type)
	}
}