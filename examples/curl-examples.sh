#!/bin/bash

# Currency Converter Service API Examples

BASE_URL="http://localhost:8085"

echo "=== Currency Converter Service API Examples ==="

echo -e "\n1. Get supported currencies:"
curl -X GET "$BASE_URL/currencies"

echo -e "\n\n2. Convert USD to EUR:"
curl -X POST "$BASE_URL/convert" \
  -H "Content-Type: application/json" \
  -d '{"amount": 100, "from": "USD", "to": "EUR"}'

echo -e "\n\n3. Convert EUR to JPY:"
curl -X POST "$BASE_URL/convert" \
  -H "Content-Type: application/json" \
  -d '{"amount": 50, "from": "EUR", "to": "JPY"}'

echo -e "\n\n4. Set custom exchange rate:"
curl -X POST "$BASE_URL/rates" \
  -H "Content-Type: application/json" \
  -d '{"from": "USD", "to": "EUR", "rate": 0.90}'

echo -e "\n\n5. Convert with custom rate:"
curl -X POST "$BASE_URL/convert" \
  -H "Content-Type: application/json" \
  -d '{"amount": 100, "from": "USD", "to": "EUR"}'

echo -e "\n\n6. Test error handling (invalid amount):"
curl -X POST "$BASE_URL/convert" \
  -H "Content-Type: application/json" \
  -d '{"amount": -10, "from": "USD", "to": "EUR"}'

echo -e "\n"