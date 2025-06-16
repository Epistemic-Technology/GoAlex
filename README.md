# GoAlex

GoAlex is a Go library for [OpenAlex API](https://docs.openalex.org/).

> [OpenAlex](https://openalex.org/) is a fully open catalog of the global research system. It's named after the ancient Library of Alexandria and made by the nonprofit OurResearch.

## Features

- Provides data structures for OpenAlex entities like works, authors, sources, etc.
- Easy to use lightweight API client for OpenAlex.
- Supports pagination, filtering, searching, sorting, selecting and sampling.
- Supports polite pool for higher rate limits.

## Roadmap

As of now, the library is in its early stages. The following features are planned:

- [x] Basic data structures.
- [x] Polite pool support.
- [x] Basic API client.
- [ ] Authentication support.
- [x] Pagination support.
- [ ] Filtering and searching.
- [ ] Sorting, selecting and sampling.
- [ ] Random results.
- [ ] N-gram support.
- [ ] Cursor pagination support.
- [ ] Autocomplete support.

All features will be implemented gradually, and contributions are welcome!

## Installation

```bash
go get -u github.com/Sunhill666/goalex
```

## Usage

### Client Initialization

To use the GoAlex library, you need to create a new client. You can do this by calling the `goalex.NewClient()` function. By default, it uses the common pool with a rate limit. Read more about the [OpenAlex rate limits](https://docs.openalex.org/how-to-use-the-api/rate-limits-and-authentication).

```go
package main

import "github.com/Sunhill666/goalex"

func main() {
    // Create a new client with common pool
    client := goalex.NewClient()
}
```

#### Polite Pool

To use the polite pool, you need to set an email address. This is required by OpenAlex to allow higher rate limits. You can do this by using the `goalex.PolitePool` function when creating a new client.

```go
package main

import "github.com/Sunhill666/goalex"

func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))
}
```

### Get single entity

You can use the client to fetch a single entity like a work, author, or source. The following example demonstrates how to fetch a work by its ID.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Get a single work by ID
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // Fetch a work by its ID
    work, err := client.Works().Get("W2741809807")
    if err != nil {
        fmt.Printf("Error fetching work: %v\n", err)
        return
    }

    // Print the work
    workJSON, _ := json.MarshalIndent(work, "", "  ")
    fmt.Println("Work:", string(workJSON))
}
```

### Get list entities

You can use the client to fetch a list of works, authors, sources, etc. The following example demonstrates how to fetch a list of works using the `Works()` method.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Get list of works
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // Fetch a list of works
    works, err := client.Works().List()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Works:", string(workJSON))
    }
}
```

#### Pagination

You can also paginate the results by using the `Page()` and `PerPage()` methods. By default, the API returns the first page with 25 items per page.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Get list of works with custom pagination
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // Fetch a list of works with custom pagination
    works, err := client.Works().
        Page(1).     // Set the page number
        PerPage(10). // Set the number of items per page
        List()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Works with custom pagination:", string(workJSON))
    }
}
```

#### Metadata

The `Works()` method also supports metadata. You can use the `ListWithMeta()` method to get the results along with metadata.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Get list of works with metadata
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // Fetch a list of works
    worksWithMeta, err := client.Works().ListWithMeta()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Get results and metadata
    results, meta := worksWithMeta.Results, worksWithMeta.Meta

    // Print the works with metadata
    for _, work := range results {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Work results from `ListWithMeta()`:", string(workJSON))
    }
    // Print the metadata
    fmt.Printf("Metadata: %+v\n", meta)
}
```

### Filtering and Searching

The library supports filtering and searching for works, authors, sources, etc. You can use the `Filter()` and `Search()` methods to apply filters and search queries.

#### Filter Example

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Filter and search works
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // Fetch a list of works with filtering with multiple conditions.
    works, err := client.Works().Filter(map[string]any{
        "institutions.country_code": "fr+gb",
        "authors_count":             ">2",
    }).List()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Works:", string(workJSON))
    }

    // Or filter with a single condition.
    works, err := client.Works().FilterField("institutions.country_code", "fr+gb").List()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Works with single filter condition:", string(workJSON))
    }
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Issues

If you find any issues or have feature requests, please open an issue on the [GitHub repository](https://github.com/Sunhill666/goalex/issues).
