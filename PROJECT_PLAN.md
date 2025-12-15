# 📋 Currency Converter Component Project Plan (Go Implementation)

## 🎯 Project Overview
**Project Name:** Reusable Currency Converter Component  
**System:** Student Budget & Expense Tracking System (SBETS)  
**Language:** Go  
**Team Size:** 10 Members  
**Duration:** 8-10 weeks  

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
├── examples/                       # CROSS-LANGUAGE EXAMPLES
│   ├── python-client/
│   │   ├── converter_client.py     # Python client example
│   │   └── requirements.txt        # Python dependencies
│   ├── java-client/
│   │   ├── ConverterClient.java    # Java client example
│   │   └── pom.xml                 # Maven dependencies
│   ├── nodejs-client/
│   │   ├── converter-client.js     # Node.js client example
│   │   └── package.json            # NPM dependencies
│   └── curl-examples.sh            # Raw HTTP examples
├── docker-compose.yml              # Multi-service deployment
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

#### **REST API Interface (Language-Agnostic)**
```yaml
# OpenAPI 3.0 Specification
paths:
  /convert:
    post:
      summary: Convert currency amount
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                amount: { type: number }
                from: { type: string }
                to: { type: string }
      responses:
        200:
          content:
            application/json:
              schema:
                type: object
                properties:
                  convertedAmount: { type: number }
                  originalAmount: { type: number }
                  fromCurrency: { type: string }
                  toCurrency: { type: string }
                  exchangeRate: { type: number }
                  timestamp: { type: string }

  /currencies:
    get:
      summary: Get supported currencies
      responses:
        200:
          content:
            application/json:
              schema:
                type: object
                properties:
                  currencies: 
                    type: array
                    items: { type: string }

  /rates:
    post:
      summary: Set exchange rate
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
                from: { type: string }
                to: { type: string }
                rate: { type: number }
    delete:
      summary: Reset all rates to defaults
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

---

## 👥 Team Member Assignments

| Member | Primary Task | Deliverables | Estimated Hours |
|--------|--------------|--------------|-----------------|
| **Member 1** | Component Service Design | REST API design, OpenAPI spec, service architecture | 20-25 |
| **Member 2** | Component Implementation | `currency_converter.go`, `exchange_rates.go`, core logic | 20-25 |
| **Member 3** | REST API Development | API handlers, middleware, routing, HTTP service | 20-25 |
| **Member 4** | Cross-Language Clients | Python, Java, Node.js client examples | 20-25 |
| **Member 5** | SBETS System Development | Expense tracking, budget calculation, converter client | 20-25 |
| **Member 6** | Database & UI Design | SBETS database, web interface, templates | 20-25 |
| **Member 7** | Event System & WebSockets | Real-time events, WebSocket implementation | 15-20 |
| **Member 8** | Error Handling & Validation | API validation, error responses, HTTP status codes | 15-20 |
| **Member 9** | Documentation & Examples | API docs, client examples, deployment guides | 20-25 |
| **Member 10** | Testing & Integration | Unit tests, API tests, cross-language integration tests | 25-30 |

---

## 📅 Development Timeline

### **Week 1-2: Setup & Design**
- [ ] Project repository setup
- [ ] Go module initialization
- [ ] Component interface design
- [ ] REST API specification design
- [ ] Team role assignments

### **Week 3-4: Component Service Development**
- [ ] Currency converter core implementation
- [ ] REST API development
- [ ] OpenAPI specification
- [ ] Basic error handling and validation
- [ ] Unit tests for component

### **Week 5-6: Cross-Language Integration**
- [ ] Python client implementation
- [ ] Java client implementation
- [ ] Node.js client implementation
- [ ] API integration testing

### **Week 6-7: SBETS System Development**
- [ ] Expense tracking module
- [ ] Database integration
- [ ] Budget calculation logic
- [ ] Converter service client integration

### **Week 7-8: UI & Testing**
- [ ] Web interface development
- [ ] End-to-end testing
- [ ] Error handling refinement
- [ ] Documentation completion

### **Week 9-10: Final Polish**
- [ ] Code review and refactoring
- [ ] Performance optimization
- [ ] Final testing
- [ ] Project presentation preparation

---

## 🔧 Technical Requirements

### **Component Service Dependencies**
```go
// currency-converter-service/go.mod
module currency-converter-service

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/gorilla/websocket v1.5.0
    github.com/rs/cors v1.10.1
    github.com/stretchr/testify v1.8.4
)
```

### **SBETS System Dependencies**
```go
// sbets-system/go.mod
module sbets-system

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/mattn/go-sqlite3 v1.14.17
    github.com/stretchr/testify v1.8.4
)
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

---

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

### **API Level Errors**
- HTTP 400: Bad Request (invalid JSON, missing fields)
- HTTP 404: Not Found (unsupported currency)
- HTTP 500: Internal Server Error (service failures)

### **Error Types**
```go
type ConversionError struct {
    Code    string
    Message string
    Amount  float64
    From    string
    To      string
}

func (e ConversionError) Error() string {
    return fmt.Sprintf("conversion error [%s]: %s", e.Code, e.Message)
}
```

---

## 🧪 Testing Strategy

### **Unit Tests**
- Currency conversion accuracy
- Exchange rate management
- Error handling scenarios
- Input validation

### **API Tests**
- REST endpoint functionality
- Request/response validation
- Error response formats
- Cross-language client compatibility

### **Integration Tests**
- Component-system integration
- Database operations
- End-to-end expense tracking

### **Test Coverage Goals**
- Component code: 90%+
- API code: 85%+
- System code: 80%+
- Error paths: 100%

---

## 📝 Deliverables Checklist

### **Code Deliverables**
- [ ] Currency converter service (standalone)
- [ ] REST API with OpenAPI specification
- [ ] Cross-language client examples
- [ ] Expense tracking system
- [ ] Database layer
- [ ] Web interface
- [ ] Test suite

### **Documentation Deliverables**
- [ ] Component API documentation
- [ ] Client integration guides
- [ ] System user guide
- [ ] Installation instructions
- [ ] Code comments and README

### **Demonstration Materials**
- [ ] Live demo scenarios
- [ ] Cross-language client demos
- [ ] Test data sets
- [ ] Error handling examples
- [ ] Performance metrics

---

## 🚀 Build and Run Instructions

### **Development Setup**
```bash
# Clone repository
git clone <repository-url>
cd currency-converter-project

# Start component service
cd currency-converter-service
go mod tidy
make run-service

# In another terminal, start SBETS system
cd sbets-system
go mod tidy
make run-sbets

# Or use Docker Compose
docker-compose up
```

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
- Demonstrated with Python, Java, and Node.js clients
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

---

## 🎯 Final Presentation Structure

1. **Component Overview** (5 min)
   - Standalone service architecture
   - REST API design
   - Properties, methods, events, interfaces

2. **Cross-Language Demonstration** (10 min)
   - Python client consuming the service
   - Java client example
   - Node.js integration
   - SBETS system using the component

3. **Code Walkthrough** (10 min)
   - Component implementation
   - API endpoints
   - Reusability examples
   - Testing approach

4. **Q&A Session** (5 min)
   - Technical questions
   - Design decisions
   - Future improvements

---

## 📋 Risk Management

### **Technical Risks**
- **Risk:** Complex currency conversion logic
- **Mitigation:** Use simple hardcoded rates, focus on component design

### **Team Risks**
- **Risk:** Uneven workload distribution
- **Mitigation:** Regular check-ins, flexible task reassignment

### **Timeline Risks**
- **Risk:** Feature scope creep
- **Mitigation:** Strict adherence to original requirements, no new features

### **Integration Risks**
- **Risk:** Cross-language compatibility issues
- **Mitigation:** Early API testing, standard HTTP/JSON protocols

---

## 🌐 Cross-Language Usage Examples

### **Python Client**
```python
import requests

def convert_currency(amount, from_curr, to_curr):
    response = requests.post('http://localhost:8080/convert', json={
        'amount': amount,
        'from': from_curr,
        'to': to_curr
    })
    return response.json()
```

### **Java Client**
```java
public class ConverterClient {
    public ConversionResult convert(double amount, String from, String to) {
        // HTTP client implementation
    }
}
```

### **Node.js Client**
```javascript
const axios = require('axios');

async function convertCurrency(amount, from, to) {
    const response = await axios.post('http://localhost:8080/convert', {
        amount, from, to
    });
    return response.data;
}
```

---

This project plan provides a clear roadmap for implementing a truly reusable Currency Converter Component that can be consumed by applications written in any programming language through its REST API interface.