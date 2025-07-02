# GoAlex

**GoAlex** is a Go client library for the [OpenAlex API](https://docs.openalex.org/).

> [OpenAlex](https://openalex.org/) is a fully open catalog of the global research system, created by the nonprofit organization OurResearch. It's named after the ancient Library of Alexandria.

## Features

* Provides Go structs for OpenAlex entities: works, authors, sources, institutions, concepts, and venues
* Lightweight and easy-to-use API client
* Supports:

  * Pagination (standard and cursor)
  * Filtering and full-text search
  * Sorting, field selection, and random sampling
  * Grouping and aggregation
  * Autocomplete
* Polite pool support for higher rate limits
* Authentication support for premium access

## Roadmap

The project is in active development. The following features have been implemented:

* [x] Core data structures
* [x] Polite pool support
* [x] Basic API client
* [x] Authentication support
* [x] Pagination support
* [x] Filtering and searching
* [x] Sorting, selecting, and sampling
* [x] Random result retrieval
* [x] Grouping support
* [x] Cursor pagination
* [x] Autocomplete support
* [ ] N-gram support *(Not yet available in OpenAlex)*

Community contributions are welcome!

## Installation

```bash
go get -u github.com/Sunhill666/goalex
```

## Usage

### Client Initialization

Create a new client with:

```go
client := goalex.NewClient()
```

You can customize the client with options:

```go
client := goalex.NewClient(goalex.WithRetry(3, 2 * time.Second), goalex.WithTimeout(10 * time.Second))
```

To use a custom HTTP client, you can pass it as an option:

```go
client := goalex.NewClient(goalex.WithHTTPClient(&http.Client{ Timeout: 10 * time.Second }))
```

To use the polite pool (recommended for higher rate limits), provide an email:

```go
client := goalex.NewClient(goalex.PolitePool("you@example.com"))
```

To use authentication (for premium features):

```go
client := goalex.NewClient(goalex.Auth("your_api_key"))
```

---

### Fetch a Single Entity

Retrieve a work by ID:

```go
work, err := client.Works().Get("W2741809807")
```

---

### Fetch a Random Entity

Get a random work:

```go
work, err := client.Works().GetRandom()
```

---

### List Entities

Fetch a list of works:

```go
works, err := client.Works().List()
```

With pagination:

```go
works, err := client.Works().Page(1).PerPage(10).List()
```

With metadata:

```go
resultWithMeta, err := client.Works().ListWithMeta()
results, meta := resultWithMeta.Results, resultWithMeta.Meta
```

---

### Cursor Pagination

For large result sets:

```go
works, nextCursor, err := client.Works().Filter("publication_year", 2020).PerPage(100).Cursor()
```

To get the next page:

```go
nextWorks, _, err := client.Works().Filter("publication_year", 2020).PerPage(100).Cursor(nextCursor)
```

---

### Filtering and Searching

#### Filtering

```go
works, err := client.Works().FilterMap(map[string]any{
    "institutions.country_code": "fr+gb",
    "authors_count":             ">2",
}).List()
```

#### Searching

```go
works, err := client.Works().Search("machine learning").List()
```

#### Combined Search + Filter

```go
works, err := client.Works().SearchFilter(map[string]string{
    "display_name": "surgery",
    "title":        "surgery",
}, true).List()
```

---

### Sorting, Selecting, and Sampling

#### Sorting

```go
works, err := client.Works().SortMap(map[string]bool{
    "publication_year": true,
    "relevance_score":  true,
}).List()
```

#### Selecting Fields

```go
works, err := client.Works().Select("id", "doi", "display_name").List()
```

#### Sampling

```go
works, err := client.Works().Sample(2).Seed(42).List()
```

---

### Grouping

```go
grouped, err := client.Works().GroupBy("authorships.countries", true).ListGroupBy()
```

---

### Autocomplete

#### Institutions

```go
results, err := client.Institutions().AutoComplete("flori").List()
```

#### Works with filter and search

```go
results, err := client.Works().
    Filter("publication_year", 2010).
    Search("frogs").
    AutoComplete("greenhouse").
    List()
```

---

## License

Licensed under the [MIT License](LICENSE).

## Issues & Contributions

Found a bug or have a feature request? Feel free to open an issue or submit a PR on the [GitHub repo](https://github.com/Sunhill666/goalex/issues).
