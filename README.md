# Universal Unit Converter API

This project provides a robust, scalable, and modular HTTP API built in Go for performing common unit conversions across Temperature, Length, and Weight categories.

## Features

RESTful Endpoints: Supports both GET (via query parameters) and POST (via JSON body) requests for maximum flexibility.

Modular Architecture: Organized using standard Go project structure (cmd, internal/handler, internal/conversion, internal/model) for enhanced maintainability and testability.

Core Conversions:

Temperature: Celsius (C, celsius), Fahrenheit (F, fahrenheit), Kelvin (K, kelvin).

Length: Meter (m), Foot (ft), Inch (in), Kilometer (km), Centimeter (cm), Millimeter (mm), Yard (yd), Mile (mi).

Weight: Kilogram (kg), Pound (lb), Gram (g), Milligram (mg), Ounce (oz).

Error Handling: Provides clear JSON error responses for invalid input values or unsupported units.

## Getting Started

Prerequisites

You need Go installed on your system.

Running the Server

Clone the Repository (If applicable):

# Assuming you are in the root directory where the files are located

Run the Server:
The application entry point is located at cmd/server/main.go.

```
go run cmd/server/main.go
```

The API will start on port 8080.

Starting Unit Converter API server on port :8080


## API Endpoints

The single public endpoint for all conversions is: http://localhost:8080/api/v1/convert

1. GET Request (Query Parameters)

Format: Use URL query parameters to specify the value and units.
```
 ------------- --------------- ------------------- ---------------
| Parameter   | Required      | Description       | Example Value |
| value       | Yes           | The numeric value | 32            |
|             |               | to convert.       |               |
 ------------- --------------- ------------------- ---------------
| from        | Yes           | The unit to       | fahrenheit    |
|             |               | convert from.     |               |
 ------------- --------------- ------------------- ---------------
| to          | Yes           | The unit to       | celsius       |
|             |               | convert to.       |               |
 ------------- --------------- ------------------- ---------------
```
Example: Convert 10 miles to kilometers.
```
GET /api/v1/convert?value=10&from=mile&to=km
```
Successful Response (200 OK):
```
{
  "originalValue": 10,
  "fromUnit": "mile",
  "result": 16.0934,
  "toUnit": "km",
  "conversionType": "Length"
}
```

2. POST Request (JSON Body)

Format: Send a JSON object in the request body.

Example: Convert 15 pounds to kilograms.
```
POST /api/v1/convert
Content-Type: application/json

{
  "value": 15,
  "unitFrom": "pound",
  "unitTo": "kg"
}
```

Successful Response (200 OK):
```
{
  "originalValue": 15,
  "fromUnit": "pound",
  "result": 6.803875550275816,
  "toUnit": "kg",
  "conversionType": "Weight"
}
```

## Error Handling

The API returns a 400 Bad Request with a descriptive message if the input is invalid.

Example Error (Invalid Unit):
```
GET /api/v1/convert?value=10&from=mile&to=furlong
```

Error Response (400 Bad Request):
```
{
  "error": "invalid length unit 'mile' or 'furlong'"
}
```

## Project Structure

The project follows a standard Go layout to separate concerns:
```
unit-converter/
├── cmd/
│   └── server/
│       └── main.go       # Server initialization and routing.
├── internal/
│   ├── conversion/
│   │   └── logic.go      # Core business logic: conversion factors and algorithms.
│   ├── handler/
│   │   └── handler.go    # HTTP layer: request parsing, error/JSON response generation.
│   └── model/
│       └── types.go      # Data models: structs for API requests and responses.
```

## Acknowledgement
https://roadmap.sh/projects/unit-converter
