package test

import (
	"fmt"
	"path"
	"reflect"
	"testing"

	"github.com/Sunhill666/goalex/pkg/core"
	"github.com/Sunhill666/goalex/internal/model"
)

type getTester[T any] struct {
	id   string
	name string
	get  func(*core.Client, string) (*T, error)
}

func extractID[T any](entity *T) (string, error) {
	v := reflect.ValueOf(entity).Elem() // 取消指针
	idField := v.FieldByName("ID")
	if !idField.IsValid() || idField.Kind() != reflect.String {
		return "", fmt.Errorf("entity does not have an ID string field")
	}
	return path.Base(idField.String()), nil
}

func runGetTests[T any](t *testing.T, testers []getTester[T]) {
	client := core.New()

	for _, tester := range testers {
		t.Run(tester.name, func(t *testing.T) {
			t.Parallel()
			obj, err := tester.get(client, tester.id)
			if err != nil {
				t.Fatalf("Failed to get %s (%s): %v", tester.name, tester.id, err)
			}
			gotID, err := extractID(obj)
			if err != nil {
				t.Fatalf("Failed to extract ID from %s: %v", tester.name, err)
			}
			if gotID != tester.id {
				t.Errorf("Expected ID %s, got %s for %s", tester.id, gotID, tester.name)
			}
		})
	}
}

func TestWorksGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Work]{
		{
			id:   "W2741809807",
			name: "works basic",
			get: func(c *core.Client, id string) (*model.Work, error) {
				return c.Works().Get(id)
			},
		},
	})
}

func TestAuthorsGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Author]{
		{
			id:   "A5317838346",
			name: "authors basic",
			get: func(c *core.Client, id string) (*model.Author, error) {
				return c.Authors().Get(id)
			},
		},
	})
}

func TestSourcesGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Source]{
		{
			id:   "S137773608",
			name: "sources basic",
			get: func(c *core.Client, id string) (*model.Source, error) {
				return c.Sources().Get(id)
			},
		},
	})
}

func TestInstitutionsGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Institution]{
		{
			id:   "I27837315",
			name: "institutions basic",
			get: func(c *core.Client, id string) (*model.Institution, error) {
				return c.Institutions().Get(id)
			},
		},
	})
}

func TestTopicsGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Topic]{
		{
			id:   "T11636",
			name: "topics basic",
			get: func(c *core.Client, id string) (*model.Topic, error) {
				return c.Topics().Get(id)
			},
		},
	})
}

func TestKeywordsGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Keyword]{
		{
			id:   "cardiac-imaging",
			name: "keywords basic",
			get: func(c *core.Client, id string) (*model.Keyword, error) {
				return c.Keywords().Get(id)
			},
		},
	})
}

func TestPublishersGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Publisher]{
		{
			id:   "P4310319965",
			name: "publishers basic",
			get: func(c *core.Client, id string) (*model.Publisher, error) {
				return c.Publishers().Get(id)
			},
		},
	})
}

func TestFundersGet(t *testing.T) {
	t.Parallel()
	runGetTests(t, []getTester[model.Funder]{
		{
			id:   "F4320332161",
			name: "funders basic",
			get: func(c *core.Client, id string) (*model.Funder, error) {
				return c.Funders().Get(id)
			},
		},
	})
}
