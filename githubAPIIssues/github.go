package githubAPIIssues

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items          []*Issue
}

type CreateIssueRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}


type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Milestone *Milestone
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	HTMLURL     string `json:"html_url"`
	Title       string
	Description string
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func CreateIssue(createIssueRequest CreateIssueRequest) (*Issue, error) {
	client := &http.Client{}
	createIssueRequestBytes, err := json.Marshal(createIssueRequest)
	url := "https://api.github.com/repos/XiaobinZhao/fastAPI/issues"
	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(createIssueRequestBytes))
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Basic eGlhb2Jpbi56aGFvOmdocF9WY2Nwa3l3aDBWSTJlMjQ1RTB0VXNKY29aYVBWNEIycFc2ZnU=")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > http.StatusBadRequest {
		resp.Body.Close()
		return nil, fmt.Errorf("create issue failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func UpdateIssue(issueNumber int64, updateIssueRequest CreateIssueRequest) (*Issue, error) {
	client := &http.Client{}
	updateIssueRequestBytes, err := json.Marshal(updateIssueRequest)
	url := "https://api.github.com/repos/XiaobinZhao/fastAPI/issues/" +  strconv.FormatInt(issueNumber, 10)
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(updateIssueRequestBytes))
	if err != nil {
		return nil, fmt.Errorf("request for update issue failed: %s", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Basic eGlhb2Jpbi56aGFvOmdocF9WY2Nwa3l3aDBWSTJlMjQ1RTB0VXNKY29aYVBWNEIycFc2ZnU=")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > http.StatusBadRequest {
		resp.Body.Close()
		return nil, fmt.Errorf("update issue failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}


func CloseIssue(issueNumber int64) (*Issue, error) {
	client := &http.Client{}
	updateIssueRequestBytes := []byte(`{"state": "closed"}`)
	url := "https://api.github.com/repos/XiaobinZhao/fastAPI/issues/" +  strconv.FormatInt(issueNumber, 10)
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(updateIssueRequestBytes))
	if err != nil {
		return nil, fmt.Errorf("request for update issue failed: %s", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Basic eGlhb2Jpbi56aGFvOmdocF9WY2Nwa3l3aDBWSTJlMjQ1RTB0VXNKY29aYVBWNEIycFc2ZnU=")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > http.StatusBadRequest {
		resp.Body.Close()
		return nil, fmt.Errorf("update issue failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}