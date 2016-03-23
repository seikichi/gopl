package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	c := &Client{token}
	return c
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueCreateParams struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

type IssueEditParams struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

func (c *Client) request(method, path string, params interface{}, result interface{}) error {
	const endpoint = "https://api.github.com"
	url := endpoint + path

	var body io.Reader
	if params != nil {
		json, err := json.Marshal(params)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(json)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token "+c.token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetIssue(owner, repo, number string) (*Issue, error) {
	path := fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repo, number)
	var issue Issue
	if err := c.request("GET", path, nil, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (c *Client) SearchIssues(owner, repo string) ([]*Issue, error) {
	path := fmt.Sprintf("/repos/%s/%s/issues", owner, repo)
	var issues []*Issue
	if err := c.request("GET", path, nil, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (c *Client) CreateIssue(owner, repo string, params *IssueCreateParams) (*Issue, error) {
	path := fmt.Sprintf("/repos/%s/%s/issues", owner, repo)
	var issue Issue
	if err := c.request("POST", path, params, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (c *Client) EditIssue(owner, repo, number string, params *IssueEditParams) (*Issue, error) {
	path := fmt.Sprintf("/repos/%s/%s/issues/%s", owner, repo, number)
	var issue Issue
	if err := c.request("PATCH", path, params, &issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
