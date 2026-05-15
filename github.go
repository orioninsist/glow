package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/charmbracelet/glow/v2/utils"
)

// findGitHubMarkdownFile fetches a markdown file linked from the GitHub web UI.
func findGitHubMarkdownFile(u *url.URL) (*source, bool, error) {
	parts := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
	if len(parts) < 5 || (parts[2] != "blob" && parts[2] != "raw") {
		return nil, false, nil
	}

	filePath := strings.Join(parts[4:], "/")
	if path.Ext(filePath) == "" || !utils.IsMarkdownFile(filePath) {
		return nil, false, nil
	}

	rawURL := (&url.URL{
		Scheme: "https",
		Host:   "raw.githubusercontent.com",
		Path:   strings.Join([]string{parts[0], parts[1], parts[3], filePath}, "/"),
	}).String()

	resp, err := http.Get(rawURL) //nolint: noctx,bodyclose
	if err != nil {
		return nil, true, fmt.Errorf("unable to get url: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, true, fmt.Errorf("HTTP status %d", resp.StatusCode)
	}
	return &source{resp.Body, rawURL}, true, nil
}

// findGitHubREADME tries to find the correct README filename in a repository using GitHub API.
func findGitHubREADME(u *url.URL) (*source, error) {
	if src, ok, err := findGitHubMarkdownFile(u); ok {
		return src, err
	}

	owner, repo, ok := strings.Cut(strings.TrimPrefix(u.Path, "/"), "/")
	if !ok {
		return nil, fmt.Errorf("invalid url: %s", u.String())
	}

	type readme struct {
		DownloadURL string `json:"download_url"`
	}

	apiURL := fmt.Sprintf("https://api.%s/repos/%s/%s/readme", u.Hostname(), owner, repo)

	//nolint:bodyclose
	// it is closed on the caller
	res, err := http.Get(apiURL) //nolint: gosec,noctx
	if err != nil {
		return nil, fmt.Errorf("unable to get url: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read http response body: %w", err)
	}

	var result readme
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unable to parse json: %w", err)
	}

	if res.StatusCode == http.StatusOK {
		//nolint:bodyclose
		// it is closed on the caller
		resp, err := http.Get(result.DownloadURL) //nolint: noctx
		if err != nil {
			return nil, fmt.Errorf("unable to get url: %w", err)
		}

		if resp.StatusCode == http.StatusOK {
			return &source{resp.Body, result.DownloadURL}, nil
		}
	}

	return nil, errors.New("can't find README in GitHub repository")
}
