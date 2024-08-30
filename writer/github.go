package writer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type respErr struct {
	Message string `json:"message"`
}

type GitHub struct {
	key    string
	user   string
	repo   string
	branch string

	path string
}

func (github GitHub) SetCode(code string) Writer {
	github.path = code + "/index.html"
	return github
}

func (github GitHub) Write(content []byte) (n int, err error) {
	encodedContent := base64.StdEncoding.EncodeToString(content)

	body := strings.NewReader(fmt.Sprintf("{\"message\": \"create %s\", \"content\": \"%s\", \"branch\": \"%s\"}", github.path, encodedContent, github.branch))

	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", github.user, github.repo, github.path), body)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+github.key)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		var respErr respErr
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return 0, fmt.Errorf("failed to decode response: %w", err)
		}
		return 0, fmt.Errorf("failed to upload content: %s", respErr.Message)
	}

	return len(content), nil
}
