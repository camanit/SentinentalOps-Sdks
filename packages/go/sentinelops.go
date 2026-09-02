package sentinelops

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Config struct {
	Endpoint    string
	APIKey      string
	TenantID    string
	ServiceName string
}

type Incident struct {
	Title    string \json:"title"\
	Division string \json:"division"\
	Severity string \json:"severity"\
	Details  string \json:"details,omitempty"\
}

type Client struct {
	cfg Config
	hc  *http.Client
}

func NewClient(cfg Config) *Client {
	return &Client{
		cfg: cfg,
		hc:  &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *Client) ReportIncident(inc Incident) error {
	body, _ := json.Marshal(inc)
	req, err := http.NewRequest("POST", c.cfg.Endpoint+"/api/v1/incidents", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.cfg.APIKey)
	req.Header.Set("X-Sentinel-Tenant", c.cfg.TenantID)
	_, err = c.hc.Do(req)
	return err
}
