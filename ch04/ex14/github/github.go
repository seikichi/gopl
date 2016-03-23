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

type Milestone struct {
	URL          string
	HTMLURL      string `json:"html_url"`
	LabelsURL    string `json:"labels_url"`
	Number       int
	State        string
	Title        string
	Description  string
	Creator      *User
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DueOn        time.Time `json:"due_on"`
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

func (c *Client) ListIssues(owner, repo string) ([]*Issue, error) {
	path := fmt.Sprintf("/repos/%s/%s/issues", owner, repo)
	var issues []*Issue
	if err := c.request("GET", path, nil, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (c *Client) ListMilestones(owner, repo string) ([]*Milestone, error) {
	path := fmt.Sprintf("/repos/%s/%s/milestones", owner, repo)
	var milestones []*Milestone
	if err := c.request("GET", path, nil, &milestones); err != nil {
		return nil, err
	}
	return milestones, nil
}

func (c *Client) ListContributors(owner, repo string) ([]*User, error) {
	path := fmt.Sprintf("/repos/%s/%s/contributors", owner, repo)
	var users []*User
	if err := c.request("GET", path, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}
