# 🛡️ VV-Captcha-Solver

[![License: ISC](https://img.shields.io/badge/License-ISC-blue.svg)](https://opensource.org/licenses/ISC)
[![GitHub stars](https://img.shields.io/github/stars/alexabrahall/VV-Captcha-Solver.svg)](https://github.com/alexabrahall/VV-Captcha-Solver/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/alexabrahall/VV-Captcha-Solver.svg)](https://github.com/alexabrahall/VV-Captcha-Solver/network)
[![GitHub issues](https://img.shields.io/github/issues/alexabrahall/VV-Captcha-Solver.svg)](https://github.com/alexabrahall/VV-Captcha-Solver/issues)

> **A powerful, multi-language library for solving VV-Captcha challenges with ease**

VV-Captcha-Solver is a comprehensive solution for automating VV-Captcha token generation across multiple programming languages. Whether you're building web scrapers, automation tools, or need to bypass VV-Captcha protection, this library provides a simple and reliable interface.

## ✨ Features

- 🌐 **Multi-Language Support**: Python, TypeScript/JavaScript, and Go implementations
- 🚀 **High Performance**: Optimized for speed and reliability
- 🛡️ **Robust Error Handling**: Comprehensive error messages and status code handling
- 📦 **Easy Integration**: Simple API with minimal dependencies
- 🔧 **Cross-Platform**: Works on Windows, macOS, and Linux
- 📚 **Well Documented**: Clear examples and comprehensive documentation

## 🚀 Quick Start

### Python

```bash
pip install requests
```

```python
from src.python.main import solve_vv_captcha

try:
    token = solve_vv_captcha("https://example.com/protected-page")
    print(f"VVCaptcha token: {token}")
except Exception as e:
    print(f"Error: {e}")
```

### TypeScript/JavaScript

```bash
npm install axios
```

```typescript
import { solveVVCaptcha } from "./src/typescript/index";

async function main() {
  try {
    const token = await solveVVCaptcha("https://example.com/protected-page");
    console.log(`VV Captcha token: ${token}`);
  } catch (error) {
    console.error(error);
  }
}
```

### Go

```go
package main

import (
    "fmt"
    "log"
    "your-module/src/golang"
)

func main() {
    token, err := captcha.SolveVVCaptcha("https://example.com/protected-page")
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    fmt.Printf("VVCaptcha token: %s\n", token)
}
```

## 📖 API Reference

### Python API

```python
solve_vv_captcha(url: str) -> str
```

**Parameters:**

- `url` (str): The target URL containing the VV-Captcha

**Returns:**

- `str`: The captcha token

**Raises:**

- `Exception`: Various exceptions for different error conditions

### TypeScript API

```typescript
solveVVCaptcha(url: string): Promise<string>
```

**Parameters:**

- `url` (string): The target URL containing the VV-Captcha

**Returns:**

- `Promise<string>`: A promise that resolves to the captcha token

**Throws:**

- `Error`: Various errors for different failure conditions

### Go API

```go
func SolveVVCaptcha(targetURL string) (string, error)
```

**Parameters:**

- `targetURL` (string): The target URL containing the VV-Captcha

**Returns:**

- `string`: The captcha token
- `error`: Error if the operation fails

## 🔧 Installation

### Prerequisites

- **Python**: 3.6+ with `requests` library
- **Node.js**: 14+ with `axios` library
- **Go**: 1.16+

### Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/alexabrahall/VV-Captcha-Solver.git
   cd VV-Captcha-Solver
   ```

2. **Install dependencies:**

   **For Python:**

   ```bash
   pip install requests
   ```

   **For TypeScript/JavaScript:**

   ```bash
   npm install
   ```

   **For Go:**

   ```bash
   go mod tidy
   ```

## 📁 Project Structure

```
VV-Captcha-Solver/
├── src/
│   ├── python/
│   │   ├── __init__.py
│   │   └── main.py              # Python implementation
│   ├── typescript/
│   │   └── index.ts             # TypeScript implementation
│   └── golang/
│       └── captcha.go           # Go implementation
├── examples/
│   ├── example.py               # Python usage example
│   ├── example.ts               # TypeScript usage example
│   └── example.go               # Go usage example
├── package.json                 # Node.js dependencies
└── README.md                    # This file
```

## 🎯 How It Works

The VV-Captcha-Solver works by:

1. **URL Parsing**: Extracts the domain from the provided URL
2. **Request Generation**: Creates a properly formatted HTTP request to the VV-Captcha endpoint
3. **Header Configuration**: Sets appropriate User-Agent and Referer headers
4. **Response Processing**: Handles various HTTP status codes and error conditions
5. **Token Extraction**: Uses regex pattern matching to extract the captcha token
6. **Error Handling**: Provides detailed error messages for troubleshooting

## 🛠️ Error Handling

The library provides comprehensive error handling for various scenarios:

| Status Code | Error Type                  | Description                          |
| ----------- | --------------------------- | ------------------------------------ |
| 403         | Domain Verification Failed  | The domain is not VV-Captcha enabled |
| 500         | Captcha Verification Failed | Error in domain handling             |
| 200         | Success                     | Token extracted successfully         |
| Other       | Unexpected Error            | Check domain settings                |

## 🔍 Examples

### Basic Usage

```python
# Python
from src.python.main import solve_vv_captcha

token = solve_vv_captcha("https://uk-umg.com/um-forms/48709-1272284.html")
print(f"Token: {token}")
```

```typescript
// TypeScript
import { solveVVCaptcha } from "./src/typescript/index";

const token = await solveVVCaptcha(
  "https://uk-umg.com/um-forms/48709-1272284.html"
);
console.log(`Token: ${token}`);
```

```go
// Go
import "your-module/src/golang"

token, err := captcha.SolveVVCaptcha("https://uk-umg.com/um-forms/48709-1272284.html")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Token: %s\n", token)
```

### Error Handling

```python
try:
    token = solve_vv_captcha("https://invalid-domain.com")
except Exception as e:
    print(f"Failed to solve captcha: {e}")
```



## ⚠️ Disclaimer

This tool is for educational and research purposes only. Please ensure you comply with the terms of service of any websites you interact with and respect their rate limits and usage policies.

