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
- [x] Authentication support.
- [x] Pagination support.
- [x] Filtering and searching.
- [x] Sorting, selecting and sampling.
- [x] Random result.
- [x] Grouping support.
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

<details>

<summary>Click to expand</summary>

To use the GoAlex library, you need to create a new client. You can do this by calling the `goalex.NewClient()` function. By default, it uses the common pool with a rate limit. Read more about the [OpenAlex rate limits](https://docs.openalex.org/how-to-use-the-api/rate-limits-and-authentication).

```go
package main

import "github.com/Sunhill666/goalex"

func main() {
    // Create a new client with common pool
    client := goalex.NewClient()
}
```

</details>

#### Polite Pool

<details>

<summary>Click to expand</summary>

To use the polite pool, you need to set an email address. This is required by OpenAlex to allow higher rate limits. You can do this by using the `goalex.PolitePool` function when creating a new client.

```go
package main

import "github.com/Sunhill666/goalex"

func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))
}
```

</details>

#### Authentication

<details>

<summary>Click to expand</summary>

For OpenAlex premium users, you can use the `goalex.Auth` function to authenticate your client. This will allow you to access premium features and higher rate limits.

```go
package main

import "github.com/Sunhill666/goalex"

func main() {
    // Create a new client with authentication
    client := goalex.NewClient(goalex.Auth("your_api_key"))
}
```

</details>

### Get single entity

<details>

<summary>Click to expand</summary>

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

</details>

### Get random entity

<details>

<summary>Click to expand</summary>

You can use the client to fetch a random entity like a work, author, or source. The following example demonstrates how to fetch a random work.

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
    client := goalex.NewClient()

    // Fetch a work by its ID
    work, err := client.Works().GetRandom()
    if err != nil {
        fmt.Printf("Error fetching work: %v\n", err)
        return
    }

    // Print the work
    workJSON, _ := json.MarshalIndent(work, "", "  ")
    fmt.Println("Work:", string(workJSON))
}
```

</details>

### Get list entities

<details>

<summary>Click to expand</summary>

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

</details>

#### Pagination

<details>

<summary>Click to expand</summary>

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

</details>

#### Metadata

<details>

<summary>Click to expand</summary>

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

</details>

### Filtering and Searching

The library supports filtering and searching for works, authors, sources, etc.

#### Filter Example

<details>

<summary>Click to expand</summary>

You can filter works by multiple conditions using the `FilterMap()` method, and filter by a single condition using the `Filter()` method. The following example demonstrates how to filter works by country code and authors count.

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
    works, err := client.Works().FilterMap(map[string]any{
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
    works, err = client.Works().Filter("institutions.country_code", "fr+gb").List()

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

</details>

#### Search Example

<details>

<summary>Click to expand</summary>

You can search for works using the `Search()` method. The following example demonstrates how to search for works with a query.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Search works
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))
    // Search for works with a query
    works, err := client.Works().Search("machine learning").List()
    // Or, if you want to exact search
    // works, err = client.Works().Search("\"machine learning\"").List()
    if err != nil {
        fmt.Printf("Error searching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Search results:", string(workJSON))
    }
}
```

</details>

#### Search Filter

<details>

<summary>Click to expand</summary>

You can also combine search with filtering. The following example demonstrates how to search for works with a query and filter by country code.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Search works with filter
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // Search for works with a query and filter by country code
    works, err := client.Works().SearchFilter(map[string]string{
        "display_name": "surgery",
        "title":        "surgery",
    }, true).List() // Search without stemming
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

</details>

### Sorting, Selecting and Sampling

The library supports sorting, selecting, and sampling of works.

#### Sort Example

<details>

<summary>Click to expand</summary>

You can sort the results by multiple fields using the `SortMap()` method, or by a single field using the `Sort()` method. The following example demonstrates how to sort works by publication year and relevance score.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

// Example: Sort works
func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    // Fetch a list of works and sort them by publication year and relevance score
    works, err := client.Works().SearchFilter(map[string]string{
        "display_name": "bioplastics",
    }, false).SortMap(map[string]bool{
        "publication_year": true, // Descending order
        "relevance_score":  true, // Descending order
    }).List()
    if err != nil {
        fmt.Printf("Error fetching work: %v\n", err)
        return
    }

    // Print the sorted works
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Sorted works:", string(workJSON))
    }

    // Or sort by a single field
    works, err = client.Works().Sort("publication_year", true).List()
    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the sorted works
    for _, work := range works {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Sorted works by single field:", string(workJSON))
    }
}
```

</details>

#### Select Example

<details>

<summary>Click to expand</summary>

You can select specific fields from the results using the `Select()` method. The following example demonstrates how to select the ID, DOI, and display name of works.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    selectedWorks, err := client.Works().Select("id", "doi", "display_name").List()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range selectedWorks {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Works:", string(workJSON))
    }
}
```

</details>

#### Sample Example

<details>

<summary>Click to expand</summary>

You can sample a specific number of works using the `Sample()` method. And you can also set a seed for reproducibility using the `Seed()` method. The following example demonstrates how to sample 2 works with a seed of 42.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

func main() {
    // Create a new client with polite pool
    client := goalex.NewClient(goalex.PolitePool("you@example.com"))

    sampleWorks, err := client.Works().Sample(2).Seed(42).List()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range sampleWorks {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Works:", string(workJSON))
    }
}
```

</details>

### Grouping Example

<details>

<summary>Click to expand</summary>

You can group works by a specific field using the `GroupBy()` method. The following example demonstrates how to group works by authorship countries and include unknown countries.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/Sunhill666/goalex"
)

func main() {
    // Create a new client with polite pool
    client := goalex.NewClient()

    // Group works by authorship countries and include unknown
    groupWorks, err := client.Works().GroupBy("authorships.countries", true).ListGroupBy()

    if err != nil {
        fmt.Printf("Error fetching works: %v\n", err)
        return
    }

    // Print the works
    for _, work := range groupWorks {
        workJSON, _ := json.MarshalIndent(work, "", "  ")
        fmt.Println("Group Work:", string(workJSON))
    }
}
```

</details>

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Issues

If you find any issues or have feature requests, please open an issue on the [GitHub repository](https://github.com/Sunhill666/goalex/issues).
