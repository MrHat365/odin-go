package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	// BaseURL Odin.fun API的基础URL
	BaseURL = "https://api.odin.fun/v1"
)

// Client Odin.fun API客户端结构
type Client struct {
	httpClient *http.Client
	Token      string // 用于授权的令牌
}

// NewClient 创建一个新的Odin.fun API客户端
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

// SetToken 设置授权令牌
func (c *Client) SetToken(token string) {
	c.Token = token
}

// Get 发送GET请求
func (c *Client) Get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", BaseURL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	return body, nil
}

// Post 发送POST请求，带有JSON数据
func (c *Client) Post(endpoint string, data interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s", BaseURL, endpoint)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON编码失败: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d, 错误: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// PostMultipart 发送带有表单数据的POST请求
func (c *Client) PostMultipart(endpoint string, formData map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", BaseURL, endpoint)

	// 创建表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range formData {
		if err := writer.WriteField(key, value); err != nil {
			return nil, fmt.Errorf("写入表单字段失败: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("关闭表单写入器失败: %w", err)
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d, 错误: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
