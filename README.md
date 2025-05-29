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
- [ ] Basic API client.
- [ ] Authentication support.
- [ ] Pagination support.
- [ ] Filtering and searching.
- [ ] Sorting, selecting and sampling.
- [ ] Random results.
- [ ] N-gram support.
- [ ] Cursor pagination support.
- [ ] Autocomplete support.

All features will be implemented gradually, and contributions are welcome!

## Installation

```bash
go get github.com/Sunhill666/goalex
```

## Usage

```go
import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

func main() {
    // Set a email to use polite pool.
    // See: https://docs.openalex.org/how-to-use-the-api/rate-limits-and-authentication#the-polite-pool
    client := goalex.NewClient(goalex.WithMailto("you@example.com"))

    // or use common pool.
    // client := goalex.NewClient()

    params := &goalex.QueryParams{
        Pagination: &goalex.PaginationParams{
            Page:    1,
            PerPage: 20,
        },
    }
    // Example: List works by query
    works, err := client.Works().List(params)
    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the results
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println(string(workJSON))
    }
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Issues

If you find any issues or have feature requests, please open an issue on the [GitHub repository](https://github.com/Sunhill666/goalex/issues).
