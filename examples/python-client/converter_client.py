#!/usr/bin/env python3

import requests
import json

class CurrencyConverterClient:
    def __init__(self, base_url="http://localhost:8080"):
        self.base_url = base_url
    
    def convert(self, amount, from_currency, to_currency):
        """Convert currency amount"""
        url = f"{self.base_url}/convert"
        data = {
            "amount": amount,
            "from": from_currency,
            "to": to_currency
        }
        
        response = requests.post(url, json=data)
        if response.status_code == 200:
            return response.json()
        else:
            raise Exception(f"Conversion failed: {response.json()}")
    
    def get_currencies(self):
        """Get supported currencies"""
        url = f"{self.base_url}/currencies"
        response = requests.get(url)
        if response.status_code == 200:
            return response.json()["currencies"]
        else:
            raise Exception(f"Failed to get currencies: {response.json()}")
    
    def set_rate(self, from_currency, to_currency, rate):
        """Set custom exchange rate"""
        url = f"{self.base_url}/rates"
        data = {
            "from": from_currency,
            "to": to_currency,
            "rate": rate
        }
        
        response = requests.post(url, json=data)
        if response.status_code != 200:
            raise Exception(f"Failed to set rate: {response.json()}")

def main():
    client = CurrencyConverterClient()
    
    print("=== Currency Converter Python Client ===")
    
    # Get supported currencies
    print("\n1. Supported currencies:")
    currencies = client.get_currencies()
    print(currencies)
    
    # Convert USD to EUR
    print("\n2. Convert $100 USD to EUR:")
    result = client.convert(100, "USD", "EUR")
    print(f"${result['originalAmount']} USD = €{result['convertedAmount']:.2f} EUR")
    print(f"Exchange rate: {result['exchangeRate']}")
    
    # Convert EUR to JPY
    print("\n3. Convert €50 EUR to JPY:")
    result = client.convert(50, "EUR", "JPY")
    print(f"€{result['originalAmount']} EUR = ¥{result['convertedAmount']:.2f} JPY")
    
    # Set custom rate and convert
    print("\n4. Set custom USD->EUR rate to 0.90:")
    client.set_rate("USD", "EUR", 0.90)
    
    print("\n5. Convert $100 USD to EUR with custom rate:")
    result = client.convert(100, "USD", "EUR")
    print(f"${result['originalAmount']} USD = €{result['convertedAmount']:.2f} EUR")
    print(f"Custom exchange rate: {result['exchangeRate']}")

if __name__ == "__main__":
    main()