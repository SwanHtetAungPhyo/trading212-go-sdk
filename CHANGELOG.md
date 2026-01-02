# Changelog

## [v1.1.0] - 2026-01-02 - Fixed API Endpoints and Authentication

### 🔧 Fixed Issues
- **401 Authentication Error**: Fixed incorrect API endpoints that were causing authentication failures
- **Account Summary Endpoint**: Changed from `/api/v0/equity/account/summary` to correct endpoints:
  - `/api/v0/equity/account/info` for account metadata
  - `/api/v0/equity/account/cash` for cash balances
- **Positions Endpoint**: Changed from `/api/v0/equity/positions` to `/api/v0/equity/portfolio`

### 📝 Updated Data Structures
- **AccountInfo**: New struct for account metadata (ID, currency)
- **AccountCash**: New struct for cash balances (free, invested, result, total)
- **AccountSummary**: Updated to combine AccountInfo and AccountCash
- **Position**: Updated fields to match actual API response:
  - `averagePrice`, `currentPrice`, `quantity`, `ticker`
  - `ppl` (profit/loss), `fxPpl` (FX profit/loss)
  - `pieQuantity`, `maxBuy`, `maxSell`

### ✨ New Features
- **Separate Methods**: Added `GetAccountInfo()` and `GetAccountCash()` methods
- **Test Script**: Added `examples/test_auth.go` for easy API credential testing
- **Better Error Handling**: Improved error messages for authentication issues

### 📚 Documentation Updates
- Updated README.md with correct API usage examples
- Added getting started guide with credential testing
- Updated code examples to reflect new data structures

### 🧪 Testing
To test your API credentials:
```bash
export TRADING212_API_KEY=your_api_key
export TRADING212_API_SECRET=your_api_secret
go run examples/test_auth.go
```

### ⚠️ Breaking Changes
- `AccountSummary` struct fields have changed
- `Position` struct fields have changed
- Old field names like `AvailableToTrade`, `InPies`, etc. are no longer available

### 🔄 Migration Guide
If you were using the old structure:

**Before:**
```go
fmt.Printf("Available: %.2f\n", summary.Cash.AvailableToTrade)
fmt.Printf("Ticker: %s\n", pos.Instrument.Ticker)
```

**After:**
```go
fmt.Printf("Free: %.2f\n", summary.Cash.Free)
fmt.Printf("Ticker: %s\n", pos.Ticker)
```