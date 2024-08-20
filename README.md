Go version: 1.22

# Usage
Run a Open Telemetry collector and trace viewing UI
```
docker-compose up -d
```

Run the example:
```
go run main.go
```

View the trace at http://localhost:16686/search, the service will be called `unknown_service`.
