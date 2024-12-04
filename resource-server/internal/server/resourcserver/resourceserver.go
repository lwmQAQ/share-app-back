package resourceserver

import (
	"fmt"

	"resource-server/internal/constants"
	"resource-server/internal/models"
	"resource-server/utils"
)

type ResourceServer struct {
	contents []string
	esClient *utils.ESClient
}

func NewResourceServer(esClient *utils.ESClient) *ResourceServer {
	return &ResourceServer{
		esClient: esClient,
		contents: []string{"title", "tags"},
	}
}

func (s *ResourceServer) SelectByAll(input string) ([]models.Resource, error) {
	var resp []models.Resource
	err := s.esClient.MultiMatchSearch(s.contents, input, constants.ResourceIndex, &resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, t := range resp {
		fmt.Println(t)
	}
	return resp, nil
}
