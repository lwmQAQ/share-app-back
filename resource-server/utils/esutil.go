package utils

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"resource-server/config"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	esapi "github.com/elastic/go-elasticsearch/v8/esapi"
)

type ESClient struct {
	client *elasticsearch.Client
}

// SearchResponse 是 Elasticsearch 的查询响应结构
// Elasticsearch 通用响应结构
type ESResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string          `json:"_index"`
			ID     string          `json:"_id"`
			Score  float64         `json:"_score"`
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func NewESlient(config *config.ESConfig) *ESClient {
	var err error
	// 创建自定义的 Transport，忽略证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 忽略证书验证
		},
	}
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{config.Url},
		Username:  config.User,
		Password:  config.Password,
		Transport: tr, // 使用自定义 Transport
	})
	if err != nil {
		panic(err)
	}
	return &ESClient{
		client: client,
	}
}

// 添加索引
func (e *ESClient) AddIndex(model interface{}, index string) {
	// 生成索引映射
	mapping := generateMapping(model)
	// 将映射转换为 JSON
	body, err := json.Marshal(map[string]interface{}{
		"settings": map[string]interface{}{
			"number_of_shards":   1,
			"number_of_replicas": 1,
		},
		"mappings": mapping,
	})
	if err != nil {
		log.Fatalf("Error marshaling index settings: %s", err)
		return
	}

	// 创建索引
	req := esapi.IndicesCreateRequest{
		Index: index,
		Body:  bytes.NewReader(body),
	}

	// 执行请求
	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
		return
	}
	defer res.Body.Close()

	// 输出创建索引的结果
	if res.IsError() {
		log.Printf("Error creating index %s: %s", index, res.Status())
	} else {
		log.Printf("Index %s created successfully", index)
	}
}

// 删除索引
func (e *ESClient) DeleteIndex(index string) error {
	// 发送删除索引的请求
	res, err := e.client.Indices.Delete([]string{index})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// 检查返回的状态码
	if res.IsError() {
		return fmt.Errorf("Failed to delete index with status: %s", res.Status())
	}

	// 打印响应内容
	log.Printf("Index %s deleted successfully.", index)
	return nil
}

// 查找文档索引
func (e *ESClient) SearchByID(index string, docID string, result interface{}) error {
	res, err := e.client.Get(index, "1")
	if err != nil {
		log.Fatalf("Error getting document: %s", err)
	}
	defer res.Body.Close()

	// 检查是否成功
	if res.IsError() {
		log.Printf("[%s] Error: %s", res.Status(), res.String())
		return fmt.Errorf("查找失败")
	}
	// 将文档解析为 Map 类型
	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// 获取 _source 字段
	source, ok := doc["_source"].(map[string]interface{})
	if !ok {
		log.Println("Error: _source is not available in the response")
		return fmt.Errorf("查找失败")
	}

	// 通过反射将 _source 解析到传入的结构体中
	if err := mapToStruct(source, result); err != nil {
		log.Fatalf("Error mapping to struct: %s", err)
		return err
	}

	return nil
}

// 向索引插入数据
// 向索引插入数据
func (e *ESClient) InsertDocument(index string, document interface{}, docID string) error {
	fmt.Println(index, docID)

	// 将结构体编码为 JSON
	docJSON, err := json.Marshal(document)
	if err != nil {
		// 这里返回错误而不是直接终止程序
		return fmt.Errorf("Error marshaling document: %s", err)
	}
	// 输出 JSON 字符串以查看内容
	fmt.Printf("json: %s\n", docJSON)

	response, err := e.client.Index(
		index,
		bytes.NewReader(docJSON),
		e.client.Index.WithDocumentID(docID),
		e.client.Index.WithRefresh("true"),
	)
	if err != nil {
		fmt.Printf("插入文档失败: %s", err)
		return err
	}
	defer response.Body.Close()

	if response.IsError() {
		fmt.Printf("Elasticsearch 错误: %s", response.String())
		return fmt.Errorf("elasticsearch error: %s", response.String())
	}

	fmt.Println("文档插入成功")
	return nil
}

/*
分页查询
page: 第几页
pagesize: 页大小
result: 查询结果
orderby：根据哪个字段排序
*/
func (e *ESClient) PageSearch(page int, pagesize int, input string, index string, result interface{}, orderby string) error {
	// 计算分页查询的起始位置
	from := (page - 1) * pagesize
	if input == "" {
		input = "*"
	}
	// 构建查询体
	var query map[string]interface{}
	query = map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": input, // 匹配 input 在所有字段中的内容
			},
		},
		"from": from,
		"size": pagesize,
	}

	// 处理排序
	if orderby != "" {
		query["sort"] = []interface{}{
			map[string]interface{}{
				orderby: map[string]interface{}{
					"order": "asc", // 默认升序，或者根据需要改成 desc
				},
			},
		}
	}
	// 将查询结构体编码为 JSON 格式
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("failed to marshal query: %v", err)
	}

	// 创建请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(queryJSON),
	}

	// 执行查询
	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute range search: %v", err)
	}
	defer res.Body.Close()

	// 检查响应状态码
	if res.IsError() {
		return fmt.Errorf("search query failed with status: %s", res.Status())
	}

	// 读取响应体
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// 解析响应数据
	if err := ParseESResponse(string(body), result); err != nil {
		return fmt.Errorf("error parsing search response: %v", err)
	}

	return nil

}

/*
范围查询
content 字段
index 索引
result 返回结果
*/
func (e *ESClient) RangeSearch(content string, index string, min int, max int, result interface{}) error {
	// 构建查询体
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				content: map[string]interface{}{
					"gte": min,
					"lte": max,
				},
			},
		},
	}

	// 将查询结构体编码为 JSON 格式
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("failed to marshal query: %v", err)
	}

	// 创建请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(queryJSON),
	}

	// 执行查询
	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute range search: %v", err)
	}
	defer res.Body.Close()

	// 检查响应状态码
	if res.IsError() {
		return fmt.Errorf("search query failed with status: %s", res.Status())
	}

	// 读取响应体
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// 解析响应数据
	if err := ParseESResponse(string(body), result); err != nil {
		return fmt.Errorf("error parsing search response: %v", err)
	}

	return nil
}

/*
匹配查询
content 字段
input 查询内容
index 索引
result 返回结果
*/
func (e *ESClient) MatchSearch(content string, input string, index string, result interface{}) error {

	// 构建查询体
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				content: input, // 将 content 参数动态传入
			},
		},
	}

	// 将查询结构体编码为 JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error marshaling query: %s", err)
		return err
	}
	res, err := e.client.Search(
		e.client.Search.WithIndex(index),
		e.client.Search.WithBody(bytes.NewReader(queryJSON)),
		e.client.Search.WithPretty(), // 格式化输出
	)
	// 检查是否成功
	if err != nil || res.IsError() {
		log.Printf("[%s] Error: %s", res.Status(), res.String())
		return fmt.Errorf("查找失败")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = ParseESResponse(string(body), result)
	return err
}

/*
多字段匹配查询
contents 字段
input 查询内容
index 索引
result 返回结果
*/
func (e *ESClient) MultiMatchSearch(contents []string, input string, index string, result interface{}) error {
	// 构建 multi_match 查询
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     input,    // 查询的内容
				"fields":    contents, // 查询的字段列表
				"operator":  "and",    // 使用 AND 运算符来连接多个字段的匹配，默认为 "or"
				"fuzziness": "AUTO",   // 可以设置模糊匹配的程度
			},
		},
	}

	// 将查询结构体编码为 JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return err
	}

	// 创建请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  strings.NewReader(string(queryJSON)),
	}

	// 执行查询
	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return fmt.Errorf("failed to execute multi-match search: %v", err)
	}
	defer res.Body.Close()

	// 检查响应状态码
	if res.IsError() {
		return fmt.Errorf("search query failed with status: %s", res.Status())
	}

	// 读取响应内容
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	fmt.Println(string(body))
	// 解析响应内容到 result 中
	err = ParseESResponse(string(body), result)
	if err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	return nil
}

/*
精确匹配查询
*/
func (e *ESClient) TermSearch(content string, input string, index string, result interface{}) error {
	// 构建查询体
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				fmt.Sprintf("%s.keyword", content): input, // 精确匹配输入的字段值
			},
		},
	}

	// 将查询结构体编码为 JSON 格式
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("failed to marshal query: %v", err)
	}

	// 创建请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(queryJSON),
	}

	// 执行查询
	res, err := req.Do(context.Background(), e.client)
	// 检查是否成功
	if err != nil || res.IsError() {
		log.Printf("[%s] Error: %s", res.Status(), res.String())
		return fmt.Errorf("查找失败")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = ParseESResponse(string(body), result)
	return err
}

/*
模糊查询
*/
func (e *ESClient) FuzzySearch(content string, input string, index string, result interface{}) error {
	// 创建一个查询结构体
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"fuzzy": map[string]interface{}{},
		},
	}

	// 使用 content 作为字段名，构建 fuzzy 查询
	query["query"].(map[string]interface{})["fuzzy"].(map[string]interface{})[content] = map[string]interface{}{
		"value":     input,
		"fuzziness": "AUTO", // 或者你可以使用 "AUTO"
	}
	// 将查询体编码成 JSON 格式
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("failed to marshal query: %v", err)
	}
	res, err := e.client.Search(
		e.client.Search.WithIndex(index),
		e.client.Search.WithBody(bytes.NewReader(queryJSON)),
		e.client.Search.WithPretty(), // 格式化输出
	)
	// 检查是否成功
	if err != nil || res.IsError() {
		log.Printf("[%s] Error: %s", res.Status(), res.String())
		return fmt.Errorf("查找失败")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = ParseESResponse(string(body), result)
	return err
}

/*==============================================工具==================================================*/

// 将 map 转换为结构体
func mapToStruct(m map[string]interface{}, result interface{}) error {
	// 获取传入结构体的反射值
	v := reflect.ValueOf(result).Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("result must be a pointer to a struct")
	}

	// 遍历 map，映射到结构体字段
	for key, value := range m {
		field := v.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			// 将值转换为结构体字段的类型
			val := reflect.ValueOf(value)
			if field.Type() == val.Type() {
				field.Set(val)
			} else {
				// 如果类型不匹配，可以进行类型转换（这里假设字段类型匹配）
				field.Set(reflect.ValueOf(value))
			}
		}
	}
	return nil
}

// 动态生成映射 将结构体转换成map
func generateMapping(model interface{}) map[string]interface{} {
	t := reflect.TypeOf(model)
	properties := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// 获取字段的标签（es 类型）
		esType := field.Tag.Get("elasticsearch")
		if esType == "" {
			esType = "text" // 默认类型为 text
		}
		properties[field.Name] = map[string]string{
			"type": esType,
		}
	}

	return map[string]interface{}{
		"properties": properties,
	}
}

// 通用的解析方法，接收一个结构体类型的切片
func ParseESResponse(response string, result interface{}) error {
	// 解析 Elasticsearch 响应
	var esResponse ESResponse
	err := json.Unmarshal([]byte(response), &esResponse)
	if err != nil {
		return fmt.Errorf("failed to parse Elasticsearch response: %v", err)
	}

	// 获取 result 的类型并确保它是一个切片
	val := reflect.ValueOf(result)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result must be a pointer to a slice")
	}

	// 获取切片的元素类型
	elemType := val.Elem().Type().Elem()

	// 解析每个 hit 的 _source 字段，并将其映射到结构体中
	for _, hit := range esResponse.Hits.Hits {
		// 创建结构体实例
		elemPtr := reflect.New(elemType).Interface()

		// 将 _source 字段解析为结构体
		err := json.Unmarshal(hit.Source, &elemPtr)
		if err != nil {
			return fmt.Errorf("failed to unmarshal hit _source: %v", err)
		}

		// 将解析出来的元素添加到结果切片中
		val.Elem().Set(reflect.Append(val.Elem(), reflect.ValueOf(elemPtr).Elem()))
	}

	return nil
}
