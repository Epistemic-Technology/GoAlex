package service

import (
	"fmt"

	"github.com/Sunhill666/goalex/pkg/client"
	"github.com/Sunhill666/goalex/pkg/model"
)

type WorksService struct {
	c *client.Client
}

func NewWorksService(c *client.Client) *WorksService {
	return &WorksService{c: c}
}

func (w *WorksService) Get(id string) (*model.Work, error) {
    var work model.Work
    err := w.c.Get(fmt.Sprintf("/works/%s", id), &work)
    if err != nil {
        return nil, err
    }
    return &work, nil
}

func (w *WorksService) List(params *QueryParams) ([]model.Work, error) {
    return ListEntities[model.Work](w.c, "/works", params)
}
