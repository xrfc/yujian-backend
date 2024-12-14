package es

import (
	"yujian-backend/pkg/log"

	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

func InitESClient() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.GetLogger().Error("创建ES客户端失败: %v", err)
		return
	}

	postESInstance = &PostES{
		client: esClient,
	}
}
