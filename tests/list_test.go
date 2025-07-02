package test

import (
	"testing"

	"github.com/Sunhill666/goalex/internal/model"
	"github.com/Sunhill666/goalex/pkg/core"
)

type listTester[T any] struct {
	name    string
	builder func(*core.Client) *core.QueryBuilder[T]
}

func runListTests[T any](t *testing.T, testers []listTester[T]) {
	client := core.New()

	for _, tester := range testers {
		t.Run(tester.name, func(t *testing.T) {
			t.Parallel()
			list, err := tester.builder(client).List()

			if err != nil {
				t.Fatalf("Failed to list %s: %v", tester.name, err)
			}
			if len(list) != 25 {
				t.Errorf("Expected 25 items in list for %s, got %d", tester.name, len(list))
			}
		})
	}
}

func TestWorksList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Work]{
		{
			name: "works basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Work] {
				return c.Works()
			},
		},
	})
}

func TestAuthorsList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Author]{
		{
			name: "authors basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Author] {
				return c.Authors()
			},
		},
	})
}

func TestSourcesList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Source]{
		{
			name: "sources basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Source] {
				return c.Sources()
			},
		},
	})
}

func TestInstitutionsList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Institution]{
		{
			name: "institutions basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Institution] {
				return c.Institutions()
			},
		},
	})
}

func TestTopicsList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Topic]{
		{
			name: "topics basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Topic] {
				return c.Topics()
			},
		},
	})
}

func TestKeywordsList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Keyword]{
		{
			name: "keywords basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Keyword] {
				return c.Keywords()
			},
		},
	})
}

func TestPublishersList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Publisher]{
		{
			name: "publishers basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Publisher] {
				return c.Publishers()
			},
		},
	})
}

func TestFundersList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Funder]{
		{
			name: "funders basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Funder] {
				return c.Funders()
			},
		},
	})
}

func TestConceptsList(t *testing.T) {
	t.Parallel()
	runListTests(t, []listTester[model.Concept]{
		{
			name: "concepts basic",
			builder: func(c *core.Client) *core.QueryBuilder[model.Concept] {
				return c.Concepts()
			},
		},
	})
}
