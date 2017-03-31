package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Repo struct {
	Name string `json:"name"`
}

type PR struct {
	UpdatedAt string `json:"updated_at"`
}

type restAPI struct {
	username string
	token    string
}

func NewRestAPI(username, token string) GithubAPI {
	return &restAPI{
		username: username,
		token:    token,
	}
}

func (r restAPI) getPullRequestsByRepo(organization string) ([]RepoWithPRCount, time.Duration, int) {
	reqCount := 0

	rest1 := time.Now()
	reposURL := fmt.Sprintf("https://api.github.com/orgs/%s/repos?per_page=100", organization)
	req, err := http.NewRequest(http.MethodGet, reposURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(username, token)

	c := http.Client{}
	reqCount++
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var reposResp []Repo
	json.NewDecoder(resp.Body).Decode(&reposResp)

	var stats []RepoWithPRCount
	for _, rr := range reposResp {
		URL := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", organization, rr.Name)
		req, err = http.NewRequest(http.MethodGet, URL, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.SetBasicAuth(username, token)

		var prs []PR
		reqCount++
		resp, err := c.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		json.NewDecoder(resp.Body).Decode(&prs)
		count := len(prs)

		if count > 0 {
			stats = append(stats, RepoWithPRCount{
				Name:    rr.Name,
				PRCount: count,
			})
		}
	}
	rest2 := time.Now()

	return stats, rest2.Sub(rest1), reqCount
}
