# SBETS - Student Budget & Expense Tracking System

A web application that demonstrates the usage of the Currency Converter Service.

## Features

- Add expenses in different currencies
- Automatic conversion to USD base currency
- View expense history
- Budget summary with total expenses

## Usage

1. Start the Currency Converter Service first:
```bash
cd ../currency-converter-service
make run-service
```

2. Start SBETS:
```bash
make run-sbets
```

3. Open http://localhost:8081 in your browser

## API Endpoints

- `POST /api/expenses` - Add new expense
- `GET /api/expenses` - Get all expenses
- `GET /api/budget` - Get budget summary