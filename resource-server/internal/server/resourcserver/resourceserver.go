package resourceserver

import (
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
