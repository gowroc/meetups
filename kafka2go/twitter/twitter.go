package twitter

import (
	"fmt"
	"net/http"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

// Connect initializes a connection via Twitter API.
func Connect() *twittergo.Client {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    Key,
		ConsumerSecret: Secret,
	}
	return twittergo.NewClient(config, nil)
}

// Tweets returns a timeline with fetched tweet list.
func Tweets(client *twittergo.Client, handle string, lastTweetID uint64) (twittergo.Timeline, error) {
	url := getURL(handle, lastTweetID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build service request: %v", err)
	}

	resp, err := client.SendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send API request: %v", err)
	}

	timeline := &twittergo.Timeline{}

	if err = resp.Parse(timeline); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	return *timeline, nil
}

func getURL(handle string, lastTweetID uint64) string {
	baseURL := fmt.Sprintf("https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=%s", handle)
	if lastTweetID == 0 {
		return baseURL
	}
	return fmt.Sprintf("%s&since_id=%d", baseURL, lastTweetID)
}

// PrintTweets prints a list of tweets to the standard output.
func PrintTweets(tl twittergo.Timeline) {
	fmt.Printf("Tweets found: %d\n", len(tl))
	for _, t := range tl {
		fmt.Printf("#%d: %s\n", t.Id(), t.Text())
	}
}
