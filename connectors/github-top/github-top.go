package GithubTop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Repository struct {
	Name  string `json:"name"`
	Stars int    `json:"stargazers_count"`
}

func GetGitHubTopItem(userWant int) (string, error) {
	url := "https://api.github.com/search/repositories?q=stars:%3E1&sort=stars&order=desc"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("create request failed, err: %v", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "Go-GitHub-Client")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request failed, err: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed, status code: %d", resp.StatusCode)
	}

	var result struct {
		Items []Repository `json:"items"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("JSON decode failed, err: %v", err)
	}

	count := len(result.Items)
	if count == 0 {
		return "", fmt.Errorf("no result")
	}

	if userWant > count {
		userWant = count
	}

	tpl := []string{fmt.Sprintf("GitHub Top %d Repositories", userWant)}
	for i := 0; i < userWant; i++ {
		repo := result.Items[i]
		tpl = append(tpl, fmt.Sprintf("%d. %s - 星标：%d", i+1, repo.Name, repo.Stars))
	}

	return strings.Join(tpl, "\n"), nil
}

func HandleUserPrompt(userInput string) string {
	re := regexp.MustCompile(`\d+`)
	str := re.FindString(userInput)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return fmt.Sprintf("Error occured while parsing: %v", err)
	}

	ret, err := GetGitHubTopItem(int(num))
	if err != nil {
		return fmt.Sprintf("Error occured while fetching data: %v", err)
	}
	return ret
}
