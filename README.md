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
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // or use common pool.
    // client := goalex.NewClient()

    // Example: List works by default pagination(page=1, per_page=25)
    works, err := client.Works().List()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // If you want to get the paginated response with metadata
    worksWithMeta, err := client.Works().ListWithMeta()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Get results and metadata
    results, meta := worksWithMeta.Results, worksWithMeta.Meta

    // If you want to use custom pagination, you can do it like this:
    worksPaged, err := client.Works().
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
        fmt.Println("Works:", string(workJSON))
    }

    // Print the works with metadata
    for _, work := range results {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Work results from `ListWithMeta()`:", string(workJSON))
    }
    // Print the metadata
    fmt.Printf("Metadata: %+v\n", meta)

    // Print the paginated works
    for _, work := range worksPaged {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Works with custom pagination:", string(workJSON))
    }
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Issues

If you find any issues or have feature requests, please open an issue on the [GitHub repository](https://github.com/Sunhill666/goalex/issues).
