# Middleware & Utility Library for Go (Gin + JWT)

This library provides reusable middleware and utility components for Go applications using the Gin web framework. It includes:

- JWT Authentication (Sign & Verify)
- CORS configuration
- Custom error handling
- Environment configuration using Viper

---

## Installation

```bash
go get github.com/basputtipong/library
or 
go get github.com/basputtipong/library@latest
or
go get github.com/basputtipong/library@v1.x.x
```

Replace with your actual repository path.

---

## Package Overview

### libenv – Environment Initialization

- Reads from `configs/config.yaml`
- Supports environment variables with automatic binding

Usage:
```go
libenv.InitEnv()
```

---

### libmiddleware – JWT & CORS Middleware

#### Initialize Configuration:
```go
libmiddleware.Init()
libmiddleware.InitCorsConfig()
```

#### CORS Middleware
```go
r.Use(libmiddleware.CORSMiddleware())
```

#### JWT Generator
```go
jwtGen := libmiddleware.NewJWTGenerator()
token, err := jwtGen.Generate("user_id")
```

#### JWT Verification Middleware
```go
r.Use(libmiddleware.JWTVerify())
```

Custom claims used:
```go
type CustomClaims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}
```

---

### liberror – Structured Error Handling

Return structured errors like:
```go
return liberror.ErrorBadRequest("Invalid input", "missing email")
```

Available methods:
- ErrorBadRequest
- ErrorUnauthorized
- ErrorForbidden
- ErrorNotFound
- ErrorInternalServerError
- ErrorConflict

#### Error Handling Middleware
```go
r.Use(liberror.ErrorHandler())
```

This catches and serializes all errors into a consistent JSON response.

---

## Configuration Example

Example `configs/config.yaml`:
```yaml
internal:
  private:
    key: |
      -----BEGIN RSA PRIVATE KEY-----
      ...
      -----END RSA PRIVATE KEY-----
  public:
    key: |
      -----BEGIN PUBLIC KEY-----
      ...
      -----END PUBLIC KEY-----

whitelist:
  domain:
    - http://localhost:3000
    - https://yourdomain.com
```

---

## Example Usage in main.go

```go
func init() {
    libenv.InitEnv()
    libmiddleware.Init()
    libmiddleware.InitCorsConfig()
}

func main() {
    r := gin.Default()
    r.Use(libmiddleware.CORSMiddleware())
    r.Use(liberror.ErrorHandler())
    r.Use(libmiddleware.JWTVerify())

    // Define your routes...

    r.Run()
}
```

---

## Author

Developed by Puttipong Thammachart
