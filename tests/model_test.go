package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Sunhill666/goalex/pkg/model"
)

func TestModel(t *testing.T) {
	entities := map[string]interface{}{
		"author":      model.Author{},
		"work":        model.Work{},
		"institution": model.Institution{},
		"source":      model.Source{},
		"topic":       model.Topic{},
		"keyword":     model.Keyword{},
		"publisher":   model.Publisher{},
		"funder":      model.Funder{},
	}

	for name, entity := range entities {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile("model_json/" + name + ".json")
			if err != nil {
				t.Errorf("Error reading %s.json: %v\n", name, err)
				return
			}
			err = json.Unmarshal(data, &entity)
			if err != nil {
				t.Errorf("Error unmarshalling JSON for %s: %v\n", name, err)
				return
			}
			t.Logf("Successfully processed %s.\n", name)
		})
	}
}
