package tests

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/Sunhill666/goalex/pkg/core"
)

// TestServer provides a mock HTTP server for testing
type TestServer struct {
	*httptest.Server
	ResponseHandler func(req *http.Request) (status int, body string)
}

// NewTestServer creates a new test server with custom response handler
func NewTestServer() *TestServer {
	ts := &TestServer{}

	ts.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ts.ResponseHandler != nil {
			status, body := ts.ResponseHandler(r)
			w.WriteHeader(status)
			_, _ = w.Write([]byte(body))
		} else {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"results": [], "meta": {"count": 0, "page": 1, "per_page": 25}}`))
		}
	}))

	return ts
}

// SetResponse sets a simple response for the test server
func (ts *TestServer) SetResponse(status int, body string) {
	ts.ResponseHandler = func(req *http.Request) (int, string) {
		return status, body
	}
}

// NewTestClient creates a client configured for testing
func NewTestClient(baseURL string, opts ...core.Option) *core.Client {
	defaultOpts := []core.Option{
		core.WithTimeout(1 * time.Second),
		core.WithRetry(1, 100*time.Millisecond),
	}

	allOpts := append(defaultOpts, opts...)
	client := core.New(allOpts...)
	client.BaseURL = baseURL

	return client
}

// Sample response data for testing
const (
	SampleWorkResponse = `{
		"id": "https://openalex.org/W2741809807",
		"doi": "https://doi.org/10.7717/peerj.4375",
		"title": "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles",
		"display_name": "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles",
		"publication_year": 2018,
		"publication_date": "2018-02-13",
		"ids": {
			"openalex": "https://openalex.org/W2741809807",
			"doi": "https://doi.org/10.7717/peerj.4375",
			"mag": "2741809807",
			"pmid": "https://pubmed.ncbi.nlm.nih.gov/29456894"
		},
		"language": "en",
		"type": "article",
		"type_crossref": "journal-article",
		"open_access": {
			"is_oa": true,
			"oa_date": "2018-02-13",
			"oa_url": "https://peerj.com/articles/4375.pdf"
		}
	}`

	SampleAuthorResponse = `{
		"id": "https://openalex.org/A5023888391",
		"orcid": "https://orcid.org/0000-0002-3100-3734",
		"display_name": "Heather Piwowar",
		"display_name_alternatives": ["H. Piwowar", "Heather A. Piwowar"],
		"works_count": 38,
		"cited_by_count": 2415,
		"summary_stats": {
			"2yr_mean_citedness": 8.2,
			"h_index": 17,
			"i10_index": 21
		},
		"ids": {
			"openalex": "https://openalex.org/A5023888391",
			"orcid": "https://orcid.org/0000-0002-3100-3734"
		}
	}`

	SamplePaginatedResponse = `{
		"results": [` + SampleWorkResponse + `],
		"meta": {
			"count": 1,
			"db_response_time_ms": 123,
			"page": 1,
			"per_page": 25,
			"groups_count": null
		}
	}`

	SampleAutoCompleteResponse = `{
		"results": [
			{
				"id": "https://openalex.org/I33213144",
				"display_name": "Harvard University",
				"hint": "Cambridge, Massachusetts, United States",
				"cited_by_count": 12345678,
				"entity_type": "institution",
				"external_id": "https://ror.org/03vek6s52"
			}
		],
		"meta": {
			"count": 1,
			"db_response_time_ms": 45
		}
	}`
)
