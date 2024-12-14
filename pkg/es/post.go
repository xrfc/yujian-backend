package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"

	"yujian-backend/pkg/model"
)

const postIndex = "posts"

type PostES struct {
	client *elasticsearch.Client
}

var postESInstance *PostES

// GetPostES 获取PostES单例
func GetPostES() *PostES {
	return postESInstance
}

// CreatePost 创建帖子
func (p *PostES) CreatePost(post *model.PostDTO) error {
	data, err := json.Marshal(post)
	if err != nil {
		return err
	}

	res, err := p.client.Index(
		postIndex,
		bytes.NewReader(data),
		p.client.Index.WithDocumentID(fmt.Sprintf("%d", post.Id)),
		p.client.Index.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// GetPost 获取帖子
func (p *PostES) GetPost(id int64) (*model.PostDTO, error) {
	res, err := p.client.Get(
		postIndex,
		fmt.Sprintf("%d", id),
		p.client.Get.WithContext(context.Background()),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var post model.PostDTO
	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

// UpdatePost 更新帖子
func (p *PostES) UpdatePost(post *model.PostDTO) error {
	data, err := json.Marshal(post)
	if err != nil {
		return err
	}

	res, err := p.client.Update(
		postIndex,
		fmt.Sprintf("%d", post.Id),
		bytes.NewReader(data),
		p.client.Update.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// DeletePost 删除帖子
func (p *PostES) DeletePost(id int64) error {
	res, err := p.client.Delete(
		postIndex,
		fmt.Sprintf("%d", id),
		p.client.Delete.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// SearchPosts 搜索帖子
func (p *PostES) SearchPosts(keyword string, from, size int) ([]*model.PostDTO, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"title", "content"},
			},
		},
		"from": from,
		"size": size,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := p.client.Search(
		p.client.Search.WithContext(context.Background()),
		p.client.Search.WithIndex(postIndex),
		p.client.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	posts := make([]*model.PostDTO, 0)

	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		data, err := json.Marshal(source)
		if err != nil {
			continue
		}

		var post model.PostDTO
		if err := json.Unmarshal(data, &post); err != nil {
			continue
		}
		posts = append(posts, &post)
	}

	return posts, nil
}
