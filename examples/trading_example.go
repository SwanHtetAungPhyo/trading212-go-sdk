// Example: Trading operations with the Trading 212 Go SDK
// Run with: go run examples/trading_example.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	trading212 "trading212-go-sdk"
)

func main() {
	// Get API credentials from environment variables
	apiKey := os.Getenv("TRADING212_API_KEY")
	apiSecret := os.Getenv("TRADING212_API_SECRET")
	
	if apiKey == "" || apiSecret == "" {
		log.Fatal("Please set TRADING212_API_KEY and TRADING212_API_SECRET environment variables")
	}

	// Create client for demo environment (ALWAYS use demo for testing!)
	client := trading212.NewClient(trading212.Demo, apiKey, apiSecret)
	ctx := context.Background()

	// Example 1: Place a market buy order
	fmt.Println("=== Placing Market Buy Order ===")
	marketBuyOrder := trading212.MarketOrderRequest{
		Ticker:        "AAPL_US_EQ", // Apple stock
		Quantity:      1.0,          // Buy 1 share (positive for buy)
		ExtendedHours: false,
	}

	order, err := client.PlaceMarketOrder(ctx, marketBuyOrder)
	if err != nil {
		log.Printf("Failed to place market buy order: %v", err)
	} else {
		fmt.Printf("Market buy order placed successfully!\n")
		fmt.Printf("Order ID: %d\n", order.ID)
		fmt.Printf("Status: %s\n", order.Status)
		fmt.Printf("Ticker: %s\n", order.Ticker)
		fmt.Printf("Quantity: %.4f\n", order.Quantity)
	}

	// Wait a moment before next order
	time.Sleep(3 * time.Second)

	// Example 2: Place a limit sell order
	fmt.Println("\n=== Placing Limit Sell Order ===")
	limitSellOrder := trading212.LimitOrderRequest{
		Ticker:       "AAPL_US_EQ",
		Quantity:     -0.5, // Sell 0.5 shares (negative for sell)
		LimitPrice:   200.00, // Sell at $200 or higher
		TimeValidity: trading212.TimeValidityDay,
	}

	limitOrder, err := client.PlaceLimitOrder(ctx, limitSellOrder)
	if err != nil {
		log.Printf("Failed to place limit sell order: %v", err)
	} else {
		fmt.Printf("Limit sell order placed successfully!\n")
		fmt.Printf("Order ID: %d\n", limitOrder.ID)
		fmt.Printf("Status: %s\n", limitOrder.Status)
		fmt.Printf("Limit Price: %.2f\n", *limitOrder.LimitPrice)
	}

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Example 3: Place a stop-loss order
	fmt.Println("\n=== Placing Stop-Loss Order ===")
	stopOrder := trading212.StopOrderRequest{
		Ticker:       "MSFT_US_EQ",
		Quantity:     -1.0, // Sell 1 share (negative for sell)
		StopPrice:    300.00, // Trigger when price hits $300
		TimeValidity: trading212.TimeValidityGoodTillCancel,
	}

	stopOrderResult, err := client.PlaceStopOrder(ctx, stopOrder)
	if err != nil {
		log.Printf("Failed to place stop order: %v", err)
	} else {
		fmt.Printf("Stop order placed successfully!\n")
		fmt.Printf("Order ID: %d\n", stopOrderResult.ID)
		fmt.Printf("Stop Price: %.2f\n", *stopOrderResult.StopPrice)
	}

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Example 4: Get all pending orders
	fmt.Println("\n=== Current Pending Orders ===")
	pendingOrders, err := client.GetOrders(ctx)
	if err != nil {
		log.Printf("Failed to get pending orders: %v", err)
	} else {
		if len(pendingOrders) == 0 {
			fmt.Println("No pending orders")
		} else {
			for _, order := range pendingOrders {
				fmt.Printf("Order ID: %d - %s %s %.4f shares of %s (Status: %s)\n",
					order.ID, order.Type, order.Side, order.Quantity, order.Ticker, order.Status)
			}
		}
	}

	// Example 5: Cancel the last order (if any)
	if len(pendingOrders) > 0 {
		lastOrder := pendingOrders[len(pendingOrders)-1]
		fmt.Printf("\n=== Cancelling Order %d ===\n", lastOrder.ID)
		
		err := client.CancelOrder(ctx, lastOrder.ID)
		if err != nil {
			log.Printf("Failed to cancel order: %v", err)
		} else {
			fmt.Printf("Order %d cancelled successfully!\n", lastOrder.ID)
		}
	}

	// Example 6: Get historical orders (last 10)
	fmt.Println("\n=== Recent Historical Orders ===")
	historyOpts := &trading212.HistoryOrdersOptions{
		Limit: 10,
	}
	
	historicalOrders, err := client.GetHistoricalOrders(ctx, historyOpts)
	if err != nil {
		log.Printf("Failed to get historical orders: %v", err)
	} else {
		if len(historicalOrders.Items) == 0 {
			fmt.Println("No historical orders found")
		} else {
			for _, histOrder := range historicalOrders.Items {
				fmt.Printf("Order ID: %d - %s %s %.4f of %s (Filled: %.4f)\n",
					histOrder.Order.ID,
					histOrder.Order.Type,
					histOrder.Order.Side,
					histOrder.Order.Quantity,
					histOrder.Order.Ticker,
					histOrder.Order.FilledQuantity)
			}
		}
	}

	fmt.Println("\n=== Trading Example Complete ===")
	fmt.Println("Note: This example uses the demo environment. No real money was involved.")
}