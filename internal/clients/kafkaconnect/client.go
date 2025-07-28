package kafkaconnect

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

// Client is a Kafka Connect API client
type Client struct {
    baseURL    string
    httpClient *http.Client
    username   string
    password   string
}

// NewClient creates a new Kafka Connect client
func NewClient(baseURL string, options ...ClientOption) *Client {
    c := &Client{
        baseURL: baseURL,
        httpClient: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
    
    for _, opt := range options {
        opt(c)
    }
    
    return c
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(client *http.Client) ClientOption {
    return func(c *Client) {
        c.httpClient = client
    }
}

// WithBasicAuth sets basic authentication credentials
func WithBasicAuth(username, password string) ClientOption {
    return func(c *Client) {
        c.username = username
        c.password = password
    }
}

// ConnectorConfig represents a connector configuration
type ConnectorConfig struct {
    Name   string            `json:"name"`
    Config map[string]string `json:"config"`
}

// ConnectorStatus represents the status of a connector
type ConnectorStatus struct {
    Name      string                 `json:"name"`
    Connector ConnectorState         `json:"connector"`
    Tasks     []TaskState           `json:"tasks"`
    Type      string                 `json:"type"`
}

// ConnectorState represents the state of a connector
type ConnectorState struct {
    State    string `json:"state"`
    WorkerID string `json:"worker_id"`
    Trace    string `json:"trace,omitempty"`
}

// TaskState represents the state of a task
type TaskState struct {
    ID       int    `json:"id"`
    State    string `json:"state"`
    WorkerID string `json:"worker_id"`
    Trace    string `json:"trace,omitempty"`
}

// CreateConnector creates a new connector
func (c *Client) CreateConnector(ctx context.Context, config ConnectorConfig) (*ConnectorInfo, error) {
    body, err := json.Marshal(config)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal connector config: %w", err)
    }
    
    req, err := c.newRequest(ctx, http.MethodPost, "/connectors", bytes.NewReader(body))
    if err != nil {
        return nil, err
    }
    
    var info ConnectorInfo
    if err := c.doRequest(req, &info); err != nil {
        return nil, fmt.Errorf("failed to create connector: %w", err)
    }
    
    return &info, nil
}

// GetConnector gets a connector by name
func (c *Client) GetConnector(ctx context.Context, name string) (*ConnectorInfo, error) {
    req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/connectors/%s", name), nil)
    if err != nil {
        return nil, err
    }
    
    var info ConnectorInfo
    if err := c.doRequest(req, &info); err != nil {
        return nil, fmt.Errorf("failed to get connector: %w", err)
    }
    
    return &info, nil
}

// UpdateConnector updates a connector configuration
func (c *Client) UpdateConnector(ctx context.Context, name string, config map[string]string) (*ConnectorInfo, error) {
    body, err := json.Marshal(config)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal connector config: %w", err)
    }
    
    req, err := c.newRequest(ctx, http.MethodPut, fmt.Sprintf("/connectors/%s/config", name), bytes.NewReader(body))
    if err != nil {
        return nil, err
    }
    
    var info ConnectorInfo
    if err := c.doRequest(req, &info); err != nil {
        return nil, fmt.Errorf("failed to update connector: %w", err)
    }
    
    return &info, nil
}

// DeleteConnector deletes a connector
func (c *Client) DeleteConnector(ctx context.Context, name string) error {
    req, err := c.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/connectors/%s", name), nil)
    if err != nil {
        return err
    }
    
    if err := c.doRequest(req, nil); err != nil {
        return fmt.Errorf("failed to delete connector: %w", err)
    }
    
    return nil
}

// GetConnectorStatus gets the status of a connector
func (c *Client) GetConnectorStatus(ctx context.Context, name string) (*ConnectorStatus, error) {
    req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/connectors/%s/status", name), nil)
    if err != nil {
        return nil, err
    }
    
    var status ConnectorStatus
    if err := c.doRequest(req, &status); err != nil {
        return nil, fmt.Errorf("failed to get connector status: %w", err)
    }
    
    return &status, nil
}

// ConnectorInfo represents connector information
type ConnectorInfo struct {
    Name   string            `json:"name"`
    Config map[string]string `json:"config"`
    Tasks  []TaskInfo        `json:"tasks"`
    Type   string            `json:"type"`
}

// TaskInfo represents task information
type TaskInfo struct {
    Connector string            `json:"connector"`
    Task      int               `json:"task"`
    Config    map[string]string `json:"config"`
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
    url := c.baseURL + path
    req, err := http.NewRequestWithContext(ctx, method, url, body)
    if err != nil {
        return nil, err
    }
    
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")
    
    if c.username != "" && c.password != "" {
        req.SetBasicAuth(c.username, c.password)
    }
    
    return req, nil
}

func (c *Client) doRequest(req *http.Request, v interface{}) error {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        body, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
    }
    
    if v != nil && resp.StatusCode != http.StatusNoContent {
        if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
            return fmt.Errorf("failed to decode response: %w", err)
        }
    }
    
    return nil
}
