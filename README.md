# Trading 212 Go SDK

A comprehensive Go SDK for the Trading 212 Public API, supporting both demo (paper trading) and live trading environments.

## Features

- **Complete API Coverage**: All Trading 212 API endpoints
- **Type Safety**: Strongly typed Go structs for all API responses
- **Environment Support**: Both demo and live trading environments
- **Authentication**: Built-in HTTP Basic Auth with API key/secret
- **Error Handling**: Comprehensive error handling with detailed messages
- **Context Support**: All methods support Go context for cancellation and timeouts
- **Rate Limiting**: Respects API rate limits as documented

## Installation

```bash
go get github.com/SwanHtetAungPhyo/trading212-go-sdk
```

## Getting Started

### 1. Get API Keys
First, generate your API keys from the Trading 212 app. Follow the instructions in the [Trading 212 Help Centre](https://helpcentre.trading212.com/hc/en-us/articles/14584770928157-Trading-212-API-key).

### 2. Test Your Credentials
Use the included test script to verify your API credentials:

```bash
export TRADING212_API_KEY=your_api_key
export TRADING212_API_SECRET=your_api_secret
go run examples/test_auth.go
```

This will test all basic API endpoints and confirm your authentication is working.

### 3. Start Building
Once authentication is confirmed, you can start using the SDK in your applications.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    trading212 "github.com/SwanHtetAungPhyo/trading212-go-sdk"
)

func main() {
    // Create client for demo environment
    client := trading212.NewClient(
        trading212.Demo, // or trading212.Live for real trading
        "your-api-key",
        "your-api-secret",
    )
    
    ctx := context.Background()
    
    // Get account summary
    summary, err := client.GetAccountSummary(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Account ID: %d\n", summary.ID)
    fmt.Printf("Currency: %s\n", summary.Currency)
    fmt.Printf("Free Cash: %.2f\n", summary.Cash.Free)
    fmt.Printf("Total: %.2f\n", summary.Cash.Total)
}
```

## API Coverage

### Account Management
- `GetAccountInfo()` - Get account metadata (ID, currency)
- `GetAccountCash()` - Get account cash balances
- `GetAccountSummary()` - Get combined account information (convenience method)

### Orders
- `GetOrders()` - Get all pending orders
- `GetOrderByID(orderID)` - Get specific order by ID
- `PlaceMarketOrder(request)` - Place market order
- `PlaceLimitOrder(request)` - Place limit order
- `PlaceStopOrder(request)` - Place stop order
- `PlaceStopLimitOrder(request)` - Place stop-limit order
- `CancelOrder(orderID)` - Cancel pending order

### Positions
- `GetPositions(options)` - Get all open positions with optional ticker filter

### Instruments & Exchanges
- `GetInstruments()` - Get all tradable instruments
- `GetExchanges()` - Get all exchanges with working schedules

### Historical Data
- `GetHistoricalOrders(options)` - Get historical orders with pagination
- `GetDividends(options)` - Get dividend history with pagination
- `GetTransactions(options)` - Get transaction history with pagination

### Reports
- `RequestReport(request)` - Request CSV report generation
- `GetReports()` - Get status of all requested reports

## Examples

### Placing Orders

```go
// Market order (buy)
marketOrder := trading212.MarketOrderRequest{
    Ticker:        "AAPL_US_EQ",
    Quantity:      10.0, // positive for buy
    ExtendedHours: false,
}
order, err := client.PlaceMarketOrder(ctx, marketOrder)

// Market order (sell)
sellOrder := trading212.MarketOrderRequest{
    Ticker:   "AAPL_US_EQ", 
    Quantity: -5.0, // negative for sell
}
order, err := client.PlaceMarketOrder(ctx, sellOrder)

// Limit order
limitOrder := trading212.LimitOrderRequest{
    Ticker:       "MSFT_US_EQ",
    Quantity:     5.0,
    LimitPrice:   150.00,
    TimeValidity: trading212.TimeValidityDay,
}
order, err := client.PlaceLimitOrder(ctx, limitOrder)
```

### Getting Positions

```go
// Get all positions
positions, err := client.GetPositions(ctx, nil)

// Get positions for specific ticker
opts := &trading212.GetPositionsOptions{
    Ticker: "AAPL_US_EQ",
}
positions, err := client.GetPositions(ctx, opts)

for _, pos := range positions {
    fmt.Printf("Ticker: %s, Quantity: %.2f, Current Price: %.2f, P&L: %.2f\n",
        pos.Ticker, pos.Quantity, pos.CurrentPrice, pos.Ppl)
}
```

### Historical Data with Pagination

```go
// Get historical orders with pagination
opts := &trading212.HistoryOrdersOptions{
    Limit: 50,
}

for {
    result, err := client.GetHistoricalOrders(ctx, opts)
    if err != nil {
        log.Fatal(err)
    }
    
    // Process orders
    for _, order := range result.Items {
        fmt.Printf("Order ID: %d, Ticker: %s\n", 
            order.Order.ID, order.Order.Ticker)
    }
    
    // Check if there are more pages
    if result.NextPagePath == nil {
        break
    }
    
    // Extract cursor from next page path for next iteration
    // Implementation depends on parsing the NextPagePath URL
}
```

### Generating Reports

```go
// Request a CSV report
reportReq := trading212.PublicReportRequest{
    TimeFrom: time.Now().AddDate(0, -1, 0), // 1 month ago
    TimeTo:   time.Now(),
    DataIncluded: trading212.ReportDataIncluded{
        IncludeOrders:       true,
        IncludeDividends:    true,
        IncludeTransactions: true,
        IncludeInterest:     false,
    },
}

response, err := client.RequestReport(ctx, reportReq)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Report requested with ID: %d\n", response.ReportID)

// Check report status
reports, err := client.GetReports(ctx)
for _, report := range reports {
    if report.ReportID == response.ReportID {
        fmt.Printf("Report Status: %s\n", report.Status)
        if report.DownloadLink != nil {
            fmt.Printf("Download Link: %s\n", *report.DownloadLink)
        }
    }
}
```

## Environment Configuration

```go
// Demo environment (paper trading)
demoClient := trading212.NewClient(
    trading212.Demo,
    "demo-api-key",
    "demo-api-secret",
)

// Live environment (real money)
liveClient := trading212.NewClient(
    trading212.Live,
    "live-api-key", 
    "live-api-secret",
)
```

## Error Handling

The SDK provides detailed error information:

```go
order, err := client.PlaceMarketOrder(ctx, request)
if err != nil {
    // Error includes HTTP status code and response body
    fmt.Printf("Order failed: %v\n", err)
    return
}
```

## Rate Limiting

The SDK respects Trading 212's rate limits. The API will return rate limit errors if exceeded:

- Account summary: 1 req / 5s
- Orders: Various limits per endpoint
- Historical data: 6 req / 1m
- Market orders: 50 req / 1m

## Important Notes

### Order Limitations
- Orders can only be executed in the **main account currency**
- Only **Market Orders** are supported in the live environment
- Multi-currency accounts are not supported

### Order Direction
- **Buy orders**: Use positive quantity values
- **Sell orders**: Use negative quantity values

### Authentication
- API keys must be generated from the Trading 212 app
- Keys can be restricted to specific IP addresses for security
- Use HTTP Basic Auth with API Key as username and API Secret as password

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This SDK is not officially affiliated with Trading 212. Use at your own risk. Always test thoroughly in the demo environment before using with real money.