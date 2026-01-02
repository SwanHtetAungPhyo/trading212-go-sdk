// Simple test script to verify Trading212 API authentication
// Usage: TRADING212_API_KEY=your_key TRADING212_API_SECRET=your_secret go run examples/test_auth.go
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
		fmt.Println("Please set TRADING212_API_KEY and TRADING212_API_SECRET environment variables")
		fmt.Println("Example:")
		fmt.Println("  export TRADING212_API_KEY=your_api_key")
		fmt.Println("  export TRADING212_API_SECRET=your_api_secret")
		fmt.Println("  go run examples/test_auth.go")
		os.Exit(1)
	}

	// Create client for demo environment
	client := trading212.NewClient(trading212.Demo, apiKey, apiSecret)
	ctx := context.Background()

	fmt.Println("Testing Trading212 API authentication...")
	fmt.Printf("Using demo environment: %s\n", trading212.Demo)
	fmt.Println()

	// Test 1: Get account info
	fmt.Println("1. Testing account info...")
	info, err := client.GetAccountInfo(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to get account info: %v", err)
	}
	fmt.Printf("✅ Account ID: %d, Currency: %s\n", info.ID, info.Currency)

	// Test 2: Get account cash
	fmt.Println("\n2. Testing account cash...")
	cash, err := client.GetAccountCash(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to get account cash: %v", err)
	}
	fmt.Printf("✅ Free: %.2f, Invested: %.2f, Total: %.2f\n", cash.Free, cash.Invested, cash.Total)

	// Test 3: Get account summary (combined)
	fmt.Println("\n3. Testing account summary...")
	summary, err := client.GetAccountSummary(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to get account summary: %v", err)
	}
	fmt.Printf("✅ Account %d (%s): Total %.2f\n", summary.ID, summary.Currency, summary.Cash.Total)

	// Test 4: Get positions
	fmt.Println("\n4. Testing positions...")
	positions, err := client.GetPositions(ctx, nil)
	if err != nil {
		log.Fatalf("❌ Failed to get positions: %v", err)
	}
	fmt.Printf("✅ Found %d positions\n", len(positions))

	// Test 5: Get instruments (limited to first 5)
	fmt.Println("\n5. Testing instruments...")
	instruments, err := client.GetInstruments(ctx)
	if err != nil {
		log.Fatalf("❌ Failed to get instruments: %v", err)
	}
	fmt.Printf("✅ Found %d instruments\n", len(instruments))
	if len(instruments) > 0 {
		fmt.Printf("   First instrument: %s - %s\n", instruments[0].Ticker, instruments[0].Name)
	}

	fmt.Println("\n🎉 All tests passed! Your API credentials are working correctly.")
	fmt.Println("\nYou can now use the SDK with confidence.")
	fmt.Println("Remember to switch to trading212.Live environment for real trading.")
}