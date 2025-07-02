package tests

import (
	"net/http"
	"testing"
)

func TestWorksQuery(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	t.Run("works list", func(t *testing.T) {
		server.SetResponse(http.StatusOK, SamplePaginatedResponse)

		works, err := client.Works().List()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(works) != 1 {
			t.Errorf("Expected 1 work, got %d", len(works))
		}

		if works[0].Title != "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles" {
			t.Errorf("Unexpected work title: %s", works[0].Title)
		}
	})

	t.Run("works get by id", func(t *testing.T) {
		server.SetResponse(http.StatusOK, SampleWorkResponse)

		work, err := client.Works().Get("W2741809807")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if work.Title != "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles" {
			t.Errorf("Unexpected work title: %s", work.Title)
		}
	})

	t.Run("works cursor pagination", func(t *testing.T) {
		responseWithCursor := `{
			"results": [` + SampleWorkResponse + `],
			"meta": {
				"count": 1,
				"next_cursor": "next-cursor-value"
			}
		}`
		server.SetResponse(http.StatusOK, responseWithCursor)

		works, nextCursor, err := client.Works().Cursor()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(works) != 1 {
			t.Errorf("Expected 1 work, got %d", len(works))
		}

		if nextCursor != "next-cursor-value" {
			t.Errorf("Expected next cursor 'next-cursor-value', got '%s'", nextCursor)
		}
	})

	t.Run("works with meta", func(t *testing.T) {
		server.SetResponse(http.StatusOK, SamplePaginatedResponse)

		response, err := client.Works().ListWithMeta()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if response.Meta == nil {
			t.Error("Expected meta information")
		}

		if len(response.Results) != 1 {
			t.Errorf("Expected 1 result, got %d", len(response.Results))
		}
	})
}

func TestAuthorsQuery(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	t.Run("authors list", func(t *testing.T) {
		authorPaginatedResponse := `{
			"results": [` + SampleAuthorResponse + `],
			"meta": {
				"count": 1,
				"page": 1,
				"per_page": 25
			}
		}`
		server.SetResponse(http.StatusOK, authorPaginatedResponse)

		authors, err := client.Authors().List()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(authors) != 1 {
			t.Errorf("Expected 1 author, got %d", len(authors))
		}

		if authors[0].DisplayName != "Heather Piwowar" {
			t.Errorf("Unexpected author name: %s", authors[0].DisplayName)
		}
	})

	t.Run("authors get by id", func(t *testing.T) {
		server.SetResponse(http.StatusOK, SampleAuthorResponse)

		author, err := client.Authors().Get("A5023888391")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if author.DisplayName != "Heather Piwowar" {
			t.Errorf("Unexpected author name: %s", author.DisplayName)
		}
	})
}

func TestAutoComplete(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	t.Run("autocomplete institutions", func(t *testing.T) {
		server.ResponseHandler = func(req *http.Request) (int, string) {
			// Verify the autocomplete endpoint is called
			if req.URL.Path != "/autocomplete/institutions" {
				t.Errorf("Expected path /autocomplete/institutions, got %s", req.URL.Path)
			}

			// Verify the query parameter
			q := req.URL.Query().Get("q")
			if q != "harvard" {
				t.Errorf("Expected q=harvard, got q=%s", q)
			}

			return http.StatusOK, SampleAutoCompleteResponse
		}

		completions, err := client.Institutions().AutoComplete("harvard").List()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(completions) != 1 {
			t.Errorf("Expected 1 completion, got %d", len(completions))
		}

		if completions[0].DisplayName != "Harvard University" {
			t.Errorf("Unexpected completion name: %s", completions[0].DisplayName)
		}
	})
}

func TestAllEndpoints(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	tests := []struct {
		name     string
		testFunc func() error
	}{
		{
			name: "works",
			testFunc: func() error {
				_, err := client.Works().List()
				return err
			},
		},
		{
			name: "authors",
			testFunc: func() error {
				_, err := client.Authors().List()
				return err
			},
		},
		{
			name: "sources",
			testFunc: func() error {
				_, err := client.Sources().List()
				return err
			},
		},
		{
			name: "institutions",
			testFunc: func() error {
				_, err := client.Institutions().List()
				return err
			},
		},
		{
			name: "topics",
			testFunc: func() error {
				_, err := client.Topics().List()
				return err
			},
		},
		{
			name: "keywords",
			testFunc: func() error {
				_, err := client.Keywords().List()
				return err
			},
		},
		{
			name: "publishers",
			testFunc: func() error {
				_, err := client.Publishers().List()
				return err
			},
		},
		{
			name: "funders",
			testFunc: func() error {
				_, err := client.Funders().List()
				return err
			},
		},
		{
			name: "concepts",
			testFunc: func() error {
				_, err := client.Concepts().List()
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.ResponseHandler = func(req *http.Request) (int, string) {
				expectedPath := "/" + tt.name
				if req.URL.Path != expectedPath {
					t.Errorf("Expected path %s, got %s", expectedPath, req.URL.Path)
				}
				return http.StatusOK, `{"results": [], "meta": {"count": 0}}`
			}

			err := tt.testFunc()
			if err != nil {
				t.Errorf("Unexpected error for %s: %v", tt.name, err)
			}
		})
	}
}

func TestGetRandom(t *testing.T) {
	server := NewTestServer()
	defer server.Close()

	client := NewTestClient(server.URL)

	server.ResponseHandler = func(req *http.Request) (int, string) {
		if req.URL.Path != "/works/random" {
			t.Errorf("Expected path /works/random, got %s", req.URL.Path)
		}
		return http.StatusOK, SampleWorkResponse
	}

	work, err := client.Works().GetRandom()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if work.Title != "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles" {
		t.Errorf("Unexpected work title: %s", work.Title)
	}
}
