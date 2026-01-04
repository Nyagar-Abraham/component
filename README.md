# Currency Converter Component Project

A reusable currency converter component implemented as a standalone REST API service in Go, with a Student Budget & Expense Tracking System (SBETS) as a demonstration system.

## Project Structure

- `currency-converter-service/` - Standalone reusable component service
- `sbets-system/` - System using the component

## Quick Start

```bash
# Start the converter service
cd currency-converter-service
go run cmd/service/main.go

# In another terminal, start SBETS
cd sbets-system
go run cmd/sbets/main.go
```

