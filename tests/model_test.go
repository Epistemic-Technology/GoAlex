package tests

import (
	"encoding/json"
	"testing"

	"github.com/Sunhill666/goalex/internal/model"
)

func TestWorkModel(t *testing.T) {
	// Test Work JSON unmarshaling
	var work model.Work
	err := json.Unmarshal([]byte(SampleWorkResponse), &work)
	if err != nil {
		t.Fatalf("Failed to unmarshal work: %v", err)
	}

	// Verify basic fields
	if work.Title != "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles" {
		t.Errorf("Unexpected title: %s", work.Title)
	}

	if work.PublicationYear != 2018 {
		t.Errorf("Expected publication year 2018, got %d", work.PublicationYear)
	}

	if work.Type != "article" {
		t.Errorf("Expected type 'article', got '%s'", work.Type)
	}

	// Verify IDs
	if work.IDs.DOI != "https://doi.org/10.7717/peerj.4375" {
		t.Errorf("Unexpected DOI: %s", work.IDs.DOI)
	}

	if work.IDs.OpenAlex != "https://openalex.org/W2741809807" {
		t.Errorf("Unexpected OpenAlex ID: %s", work.IDs.OpenAlex)
	}

	// Verify open access info
	if work.OpenAccess == nil {
		t.Error("Expected open access information")
	} else {
		if !work.OpenAccess.IsOA {
			t.Error("Expected work to be open access")
		}
		if work.OpenAccess.OAURL != "https://peerj.com/articles/4375.pdf" {
			t.Errorf("Unexpected OA URL: %s", work.OpenAccess.OAURL)
		}
	}
}

func TestAuthorModel(t *testing.T) {
	// Test Author JSON unmarshaling
	var author model.Author
	err := json.Unmarshal([]byte(SampleAuthorResponse), &author)
	if err != nil {
		t.Fatalf("Failed to unmarshal author: %v", err)
	}

	// Verify basic fields
	if author.DisplayName != "Heather Piwowar" {
		t.Errorf("Unexpected display name: %s", author.DisplayName)
	}

	if author.WorksCount != 38 {
		t.Errorf("Expected works count 38, got %d", author.WorksCount)
	}

	if author.CitedByCount != 2415 {
		t.Errorf("Expected cited by count 2415, got %d", author.CitedByCount)
	}

	// Verify ORCID
	if author.ORCID != "https://orcid.org/0000-0002-3100-3734" {
		t.Errorf("Unexpected ORCID: %s", author.ORCID)
	}

	// Verify summary stats
	if author.SummaryStats == nil {
		t.Error("Expected summary stats")
	} else {
		if author.SummaryStats.HIndex != 17 {
			t.Errorf("Expected h-index 17, got %d", author.SummaryStats.HIndex)
		}
		if author.SummaryStats.I10Index != 21 {
			t.Errorf("Expected i10-index 21, got %d", author.SummaryStats.I10Index)
		}
	}

	// Verify display name alternatives
	if len(author.DisplayNameAlternatives) == 0 {
		t.Error("Expected display name alternatives")
	}
}

func TestCompletionModel(t *testing.T) {
	// Parse the autocomplete response
	var response model.PaginatedResponse[model.Completion]
	err := json.Unmarshal([]byte(SampleAutoCompleteResponse), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal autocomplete response: %v", err)
	}

	if len(response.Results) != 1 {
		t.Fatalf("Expected 1 completion result, got %d", len(response.Results))
	}

	completion := response.Results[0]

	// Verify completion fields
	if completion.DisplayName != "Harvard University" {
		t.Errorf("Unexpected completion display name: %s", completion.DisplayName)
	}

	if completion.EntityType != "institution" {
		t.Errorf("Expected entity type 'institution', got '%s'", completion.EntityType)
	}

	if completion.Hint != "Cambridge, Massachusetts, United States" {
		t.Errorf("Unexpected hint: %s", completion.Hint)
	}

	if completion.CitedByCount != 12345678 {
		t.Errorf("Expected cited by count 12345678, got %d", completion.CitedByCount)
	}
}

func TestPaginatedResponse(t *testing.T) {
	// Test paginated response unmarshaling
	var response model.PaginatedResponse[model.Work]
	err := json.Unmarshal([]byte(SamplePaginatedResponse), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal paginated response: %v", err)
	}

	// Verify meta information
	if response.Meta == nil {
		t.Fatal("Expected meta information")
	}

	if response.Meta.Count != 1 {
		t.Errorf("Expected count 1, got %d", response.Meta.Count)
	}

	if response.Meta.Page != 1 {
		t.Errorf("Expected page 1, got %d", response.Meta.Page)
	}

	if response.Meta.PerPage != 25 {
		t.Errorf("Expected per page 25, got %d", response.Meta.PerPage)
	}

	// Verify results
	if len(response.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(response.Results))
	}

	if response.Results[0].Title != "The state of OA: a large-scale analysis of the prevalence and impact of Open Access articles" {
		t.Errorf("Unexpected work title in results: %s", response.Results[0].Title)
	}
}

func TestEmptyPaginatedResponse(t *testing.T) {
	emptyResponse := `{
		"results": [],
		"meta": {
			"count": 0,
			"page": 1,
			"per_page": 25
		}
	}`

	var response model.PaginatedResponse[model.Work]
	err := json.Unmarshal([]byte(emptyResponse), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal empty response: %v", err)
	}

	if len(response.Results) != 0 {
		t.Errorf("Expected 0 results, got %d", len(response.Results))
	}

	if response.Meta.Count != 0 {
		t.Errorf("Expected count 0, got %d", response.Meta.Count)
	}
}

func TestModelJSONTags(t *testing.T) {
	// Test that JSON omitempty works correctly
	work := &model.Work{
		Title: "Test Work",
		// Leave other fields empty to test omitempty
	}

	data, err := json.Marshal(work)
	if err != nil {
		t.Fatalf("Failed to marshal work: %v", err)
	}

	// Parse back to verify only non-empty fields are included
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal work data: %v", err)
	}

	// Should contain title but not empty fields
	if result["title"] != "Test Work" {
		t.Error("Expected title to be present")
	}

	// Should not contain empty/zero value fields due to omitempty
	if _, exists := result["publication_year"]; exists {
		t.Error("Expected publication_year to be omitted when zero")
	}
}
