package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GraphQLResponse struct {
	Data struct {
		Organization struct {
			Repositories struct {
				Nodes []struct {
					Name         string `json:"name"`
					PullRequests struct {
						Count int `json:"totalCount"`
					} `json:"pullRequests"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"organization"`
	} `json:"data"`
}

type graphqlAPI struct {
	username string
	token    string
}

func NewGraphQLAPI(username, token string) GithubAPI {
	return &graphqlAPI{
		username: username,
		token:    token,
	}
}

func (g graphqlAPI) getPullRequestsByRepo(organization string) ([]RepoWithPRCount, time.Duration, int) {
	gql1 := time.Now()
	var jsonStr = []byte(`
        {"query":
            "query {
                organization(login: \"shipwallet\") {
                    repositories(first: 100, privacy: PRIVATE){
						totalCount
                        nodes {
                            name
                            pullRequests(first: 100, states: OPEN) {
                                totalCount
                            }
                        }
                    }
                }
            }"
        }`)
	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", g.token))

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var gqlStats []RepoWithPRCount
	var r GraphQLResponse
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println("json decode err")
		panic(err)
	}

	for _, r := range r.Data.Organization.Repositories.Nodes {
		if r.PullRequests.Count > 0 {
			gqlStats = append(gqlStats, RepoWithPRCount{
				Name:    r.Name,
				PRCount: r.PullRequests.Count,
			})
		}
	}
	gql2 := time.Now()

	return gqlStats, gql2.Sub(gql1), 1
}
