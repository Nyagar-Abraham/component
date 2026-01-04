# 📋 Currency Converter Component Project Plan (Go Implementation)

## 🎯 Project Overview
**Project Name:** Reusable Currency Converter Component  
**System:** Student Budget & Expense Tracking System (SBETS)  
**Language:** Go

---

## 📁 Project Structure
```
currency-converter-project/
├── currency-converter-service/     # STANDALONE REUSABLE COMPONENT
│   ├── cmd/
│   │   └── service/
│   │       └── main.go             # Component service entry point
│   ├── pkg/
│   │   ├── converter/
│   │   │   ├── interfaces.go       # ICurrencyConverter, IExchangeRateProvider
│   │   │   ├── currency_converter.go # Core component implementation
│   │   │   └── exchange_rates.go   # Hardcoded exchange rates
│   │   ├── api/
│   │   │   ├── handlers.go         # REST API handlers
│   │   │   ├── middleware.go       # API middleware
│   │   │   └── routes.go           # API routes
│   │   └── models/
│   │       ├── request.go          # API request models
│   │       └── response.go         # API response models
│   ├── docs/
│   │   ├── api_spec.yaml           # OpenAPI specification
│   │   └── component_design.md     # Component documentation
│   ├── tests/
│   │   ├── converter_test.go       # Component unit tests
│   │   └── api_test.go             # API integration tests
│   ├── go.mod                      # Component module file
│   ├── go.sum                      # Component dependencies
│   ├── Dockerfile                  # Container for deployment
│   └── README.md                   # Component README
├── sbets-system/                   # SYSTEM USING THE COMPONENT
│   ├── cmd/
│   │   └── sbets/
│   │       └── main.go             # SBETS application entry point
│   ├── pkg/
│   │   ├── expense/
│   │   │   ├── expense.go          # Expense model and operations
│   │   │   └── budget.go           # Budget calculation logic
│   │   ├── database/
│   │   │   ├── models.go           # Database models
│   │   │   └── repository.go       # Database operations
│   │   ├── client/
│   │   │   └── converter_client.go # Client to consume converter service
│   │   └── ui/
│   │       ├── handlers.go         # HTTP handlers
│   │       └── templates/          # HTML templates
│   ├── web/
│   │   ├── static/                 # CSS, JS files
│   │   └── templates/              # HTML templates
│   ├── tests/
│   │   ├── expense_test.go         # Expense module tests
│   │   └── integration_test.go     # End-to-end tests
│   ├── go.mod                      # SBETS module file
│   └── README.md                   # SBETS README
└── README.md                       # Project overview
```

---

## 🧩 Component Specifications

### **Currency Converter Component (Standalone Service)**

#### **Go Implementation (Internal)**
```go
type CurrencyConverter struct {
    baseCurrency           string
    targetCurrency         string
    exchangeRate          float64
    supportedCurrencies   []string
    lastConversionResult  ConversionResult
    exchangeRates         map[string]map[string]float64
}

// Internal Go Methods
func (c *CurrencyConverter) Convert(amount float64) (float64, error)
func (c *CurrencyConverter) SetExchangeRate(from, to string, rate float64) error
func (c *CurrencyConverter) GetSupportedCurrencies() []string
func (c *CurrencyConverter) ResetRates()
```


#### **Events (WebSocket/HTTP Callbacks)**
```go
type ConversionEvent struct {
    Type      string    `json:"type"`      // "success" or "failure"
    Timestamp time.Time `json:"timestamp"`
    Data      interface{} `json:"data"`
}
```

#### **Go Interfaces (Internal)**
```go
type ICurrencyConverter interface {
    Convert(amount float64) (float64, error)
    SetExchangeRate(from, to string, rate float64) error
    GetSupportedCurrencies() []string
    ResetRates()
}

type IExchangeRateProvider interface {
    GetRate(from, to string) (float64, error)
    SetRate(from, to string, rate float64) error
}
```

---

## 📊 Database Schema

### **expenses table**
```sql
CREATE TABLE expenses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    amount DECIMAL(10,2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    converted_amount DECIMAL(10,2) NOT NULL,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### **exchange_rates table**
```sql
CREATE TABLE exchange_rates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    from_currency VARCHAR(3) NOT NULL,
    to_currency VARCHAR(3) NOT NULL,
    rate DECIMAL(10,6) NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```





### **Supported Currencies (Hardcoded)**
- USD (US Dollar) - Base currency
- EUR (Euro)
- GBP (British Pound)
- JPY (Japanese Yen)
- CAD (Canadian Dollar)
- AUD (Australian Dollar)

### **Hardcoded Exchange Rates**
```go
var defaultRates = map[string]map[string]float64{
    "USD": {"EUR": 0.85, "GBP": 0.73, "JPY": 110.0, "CAD": 1.25, "AUD": 1.35},
    "EUR": {"USD": 1.18, "GBP": 0.86, "JPY": 129.5, "CAD": 1.47, "AUD": 1.59},
    // ... other currency pairs
}
```


## 🧪 Error Handling Scenarios

### **Component Level Errors**
- Negative amount conversion
- Unsupported currency codes
- Missing exchange rates
- Invalid input formats

### **System Level Errors**
- Database connection failures
- Invalid expense entries
- Budget calculation errors
- UI input validation


### **Component Service Makefile**
```makefile
# currency-converter-service/Makefile
.PHONY: build test run-service clean docker

build:
	go build -o bin/converter-service cmd/service/main.go

test:
	go test ./...

run-service:
	go run cmd/service/main.go

docker:
	docker build -t currency-converter-service .

clean:
	rm -rf bin/
```

### **SBETS System Makefile**
```makefile
# sbets-system/Makefile
.PHONY: build test run-sbets clean

build:
	go build -o bin/sbets cmd/sbets/main.go

test:
	go test ./...

run-sbets:
	CONVERTER_SERVICE_URL=http://localhost:8080 go run cmd/sbets/main.go

clean:
	rm -rf bin/
```

---

## 📊 Success Metrics

### **Component Reusability**
- Component accessible via REST API from any language
- Demonstrated with Python client
- Clear API specification (OpenAPI)
- Containerized for easy deployment
- Zero coupling with consuming applications

### **Code Quality**
- All tests passing
- Code coverage > 85%
- No critical linting issues
- Clear documentation

### **Functionality**
- Accurate currency conversions
- Proper error handling
- Responsive user interface
- Data persistence

