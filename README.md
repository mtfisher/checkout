# Checkout

This is a checkout example that has been writen using GO 1.14 using go mod

## Requirements

- Go 1.14 or above

## To run

The only way to run this code is via `go test` to run use the following command

```bash
go test pricing/*
go test checkout/*
go test integration/*
```

## Breakdown

- `checkout/` contains the checkout and cart logic
- `pricing/` contains the pricing logic
- `integration/` contains an integration test