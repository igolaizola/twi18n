package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dghubble/oauth1"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type Client struct {
	debug  bool
	client *twitter.Client
}

type authorize struct {
}

func (a authorize) Add(_ *http.Request) {}

type Config struct {
	Debug             bool
	APIKey            string
	APIKeySecret      string
	AccessToken       string
	AccessTokenSecret string
}

func New(cfg *Config) *Client {
	oauth1Config := oauth1.NewConfig(cfg.APIKey, cfg.APIKeySecret)
	httpClient := oauth1Config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       cfg.AccessToken,
		TokenSecret: cfg.AccessTokenSecret,
	})
	httpClient.Timeout = 30 * time.Second

	client := &twitter.Client{
		Authorizer: authorize{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}

	return &Client{
		debug:  cfg.Debug,
		client: client,
	}
}

func (c *Client) Tweet(ctx context.Context, msg string) error {
	// Tweet message
	req := twitter.CreateTweetRequest{
		Text: msg,
	}
	resp, err := c.client.CreateTweet(ctx, req)
	if err != nil {
		return fmt.Errorf("twitter: couldn't create tweet: %w", err)
	}
	if c.debug {
		js, _ := json.MarshalIndent(resp, "", "  ")
		log.Println("openai: req:", string(js))
	}
	return nil
}
