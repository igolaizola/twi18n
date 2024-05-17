package twi18n

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/igolaizola/twi18n/pkg/openai"
	"github.com/igolaizola/twi18n/pkg/twitter"
)

type Config struct {
	Debug                     bool
	TwitterAPIKey1            string
	TwitterAPISecret1         string
	TwitterAccessToken1       string
	TwitterAccessTokenSecret1 string
	TwitterAPIKey2            string
	TwitterAPISecret2         string
	TwitterAccessToken2       string
	TwitterAccessTokenSecret2 string
	OpenAIKey                 string
	OpenAIModel               string
	OpenAIHost                string
	Template                  string
	Post                      string
}

var DefaultTemplate = "Translate this tweet to English and return only the translation, nothing else:\n\n{TWEET}"

// Run runs the twi18n process.
func Run(ctx context.Context, cfg *Config) error {
	if cfg.Post == "" {
		return fmt.Errorf("twi18n: missing post")
	}
	if cfg.TwitterAPIKey1 == "" || cfg.TwitterAPISecret1 == "" || cfg.TwitterAccessToken1 == "" || cfg.TwitterAccessTokenSecret1 == "" {
		return fmt.Errorf("twi18n: missing twitter credentials for main account")
	}
	if cfg.TwitterAPIKey2 == "" || cfg.TwitterAPISecret2 == "" || cfg.TwitterAccessToken2 == "" || cfg.TwitterAccessTokenSecret2 == "" {
		return fmt.Errorf("twi18n: missing twitter credentials for second account")
	}
	if cfg.OpenAIKey == "" && cfg.OpenAIHost == "" {
		return fmt.Errorf("twi18n: missing openai key")
	}
	if cfg.OpenAIModel == "" {
		return fmt.Errorf("twi18n: missing openai model")
	}

	ai := openai.New(&openai.Config{
		Debug: cfg.Debug,
		Token: cfg.OpenAIKey,
		Host:  cfg.OpenAIHost,
		Model: cfg.OpenAIModel,
	})
	tw1 := twitter.New(&twitter.Config{
		Debug:             cfg.Debug,
		APIKey:            cfg.TwitterAPIKey1,
		APIKeySecret:      cfg.TwitterAPISecret1,
		AccessToken:       cfg.TwitterAccessToken1,
		AccessTokenSecret: cfg.TwitterAccessTokenSecret1,
	})
	tw2 := twitter.New(&twitter.Config{
		Debug:             cfg.Debug,
		APIKey:            cfg.TwitterAPIKey2,
		APIKeySecret:      cfg.TwitterAPISecret2,
		AccessToken:       cfg.TwitterAccessToken2,
		AccessTokenSecret: cfg.TwitterAccessTokenSecret2,
	})

	template := cfg.Template
	if template == "" {
		template = DefaultTemplate
	}
	msg := strings.Replace(template, "{TWEET}", cfg.Post, 1)

	// Use AI to translate the tweet
	translation, err := ai.ChatCompletion(ctx, msg)
	if err != nil {
		return fmt.Errorf("twi18n: couldn't chat completion: %w", err)
	}

	// Tweet the original tweet
	if err := tw1.Tweet(ctx, cfg.Post); err != nil {
		return fmt.Errorf("twi18n: couldn't tweet: %w", err)
	}
	log.Println("tweeted:", cfg.Post)

	// Tweet the translation
	if err := tw2.Tweet(ctx, translation); err != nil {
		return fmt.Errorf("twi18n: couldn't tweet: %w", err)
	}
	log.Println("tweeted:", translation)

	return nil
}
