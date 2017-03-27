package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	username, token, organization string
)

func init() {
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&token, "token", "", "token")
	flag.StringVar(&organization, "organization", "", "organization")
}

type RepoWithPRCount struct {
	Name    string
	PRCount int
}

type GithubAPI interface {
	getPullRequestsByRepo(organization string) ([]RepoWithPRCount, time.Duration, int)
}

func main() {
	flag.Parse()

	if username == "" || token == "" || organization == "" {
		flag.Usage()
		log.Fatalln("Please provide all flag attributes.")
	}

	rest := NewRestAPI(username, token)
	graph := NewGraphQLAPI(username, token)

	resRes, resDur, resReqCount := rest.getPullRequestsByRepo(organization)
	gqlRes, gqlDur, gqlReqCount := graph.getPullRequestsByRepo(organization)

	fmt.Printf("GraphQL: %d requests in %v\n", gqlReqCount, gqlDur)
	fmt.Printf("REST: %d requests in %v\n", resReqCount, resDur)

	fmt.Println("REST: Pull requests by repo:")
	for _, s := range resRes {
		fmt.Printf("github.com/shipwallet/%s -> %d\n", s.Name, s.PRCount)
	}

	fmt.Println("GQL: Pull requests by repo:")
	for _, s := range gqlRes {
		fmt.Printf("github.com/shipwallet/%s -> %d\n", s.Name, s.PRCount)
	}
}
