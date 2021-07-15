# Currency Converter
A simple web service to convert currencies.
it is able to convert the Ghanaian Cedi,Nigerian Naira and Kenyan Shilling from any one to the other

the server accepts a POST request with the amount to the converted, the currency and the target currency to be converted to and returns the converted amount 

 the service is started with :
 ```bash
    go run main.go
```

the endpoint :
```url
    http://127.0.0.1:3000/api/convert
```
example:
 request body:
```json
    {
        "amount":10.00,
        "currency":"GHS",
        "target":"NGN"
    }
```
response:
```json
    {
        "currency": "NGN",
        "amount": 693.1
    }
```

