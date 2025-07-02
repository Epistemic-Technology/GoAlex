# Tests

This package contains comprehensive tests for the GoAlex library.

## Test Structure

### Core Test Files

- **`common_test.go`** - Common testing utilities and test helpers
  - `TestServer` - Mock HTTP server for testing
  - `NewTestClient` - Test client factory
  - Sample response data constants

- **`client_test.go`** - Tests for the core HTTP client functionality
  - Client creation and configuration
  - HTTP request handling
  - Option application (PolitePool, Auth, etc.)

- **`params_test.go`** - Tests for query parameter handling
  - Pagination parameters
  - Query parameter serialization
  - Filter and sort parameter ordering

- **`builder_test.go`** - Tests for the query builder
  - Method chaining
  - Filter operations
  - Sort operations
  - Pagination
  - Search functionality

- **`query_test.go`** - Tests for entity-specific queries
  - Works, Authors, Sources, etc.
  - Entity retrieval by ID
  - Random entity retrieval
  - Autocomplete functionality
  - Cursor-based pagination

- **`integration_test.go`** - Integration and error handling tests
  - Retry mechanism
  - Timeout handling
  - Error response handling
  - Invalid JSON handling

- **`model_test.go`** - Tests for data model JSON serialization/deserialization
  - Work model
  - Author model
  - Completion model
  - Paginated response model

- **`goalex_test.go`** - Tests for the main package exports
  - Package-level API
  - Type aliases
  - Example usage patterns

- **`benchmark_test.go`** - Performance benchmarks
  - Client creation
  - Query building
  - HTTP requests
  - JSON decoding
  - Memory allocation profiling

## Running Tests

### All Tests

```bash
go test ./tests
```

### Specific Test Files

```bash
go test ./tests -run TestClient
go test ./tests -run TestQueryBuilder
```

### Benchmarks

```bash
go test ./tests -bench=. -benchmem
go test ./tests -bench=Benchmark -benchmem
```

### Verbose Output

```bash
go test ./tests -v
```

### Coverage

```bash
go test ./tests -cover
go test ./tests -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Patterns

### Mock Server Usage

```go
server := NewTestServer()
defer server.Close()

server.SetResponse(http.StatusOK, `{"message": "success"}`)
client := NewTestClient(server.URL)
```

### Custom Response Handler

```go
server.ResponseHandler = func(req *http.Request) (int, string) {
    // Custom logic based on request
    return http.StatusOK, `{"custom": "response"}`
}
```

### Error Testing

```go
server.SetResponse(http.StatusInternalServerError, `{"error": "server error"}`)
_, err := client.Get("/test", &result)
if err == nil {
    t.Error("Expected error but got nil")
}
```

## Test Coverage

The tests aim to cover:

- ✅ All public API methods
- ✅ Error conditions and edge cases
- ✅ Configuration options
- ✅ JSON serialization/deserialization
- ✅ HTTP retry logic
- ✅ Query parameter generation
- ✅ Method chaining
- ✅ Performance characteristics

## Best Practices

1. **Test Isolation** - Each test is independent and uses fresh mock servers
2. **Error Testing** - All error paths are tested with appropriate scenarios
3. **Performance** - Benchmarks ensure the library performs well
4. **Documentation** - All tests have clear names and comments
5. **Maintainability** - Common functionality is extracted to helpers
