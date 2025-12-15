# Currency Converter Service

Standalone REST API service for currency conversion.

## API Endpoints

- `POST /convert` - Convert currency amount
- `GET /currencies` - Get supported currencies  
- `POST /rates` - Set exchange rate
- `DELETE /rates` - Reset rates to defaults

## Usage

```bash
go run cmd/service/main.go
```

Service runs on port 8080.