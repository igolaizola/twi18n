package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"

	"github.com/igolaizola/twi18n"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/peterbourgon/ff/v3/ffyaml"
)

// Build flags
var version = ""
var commit = ""
var date = ""

func main() {
	// Create signal based context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Launch command
	cmd := newCommand()
	if err := cmd.ParseAndRun(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func newCommand() *ffcli.Command {
	fs := flag.NewFlagSet("twi18n", flag.ExitOnError)

	return &ffcli.Command{
		ShortUsage: "twi18n [flags] <subcommand>",
		FlagSet:    fs,
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
		Subcommands: []*ffcli.Command{
			newVersionCommand(),
			newRunCommand(),
		},
	}
}

func newVersionCommand() *ffcli.Command {
	return &ffcli.Command{
		Name:       "version",
		ShortUsage: "twi18n version",
		ShortHelp:  "print version",
		Exec: func(ctx context.Context, args []string) error {
			v := version
			if v == "" {
				if buildInfo, ok := debug.ReadBuildInfo(); ok {
					v = buildInfo.Main.Version
				}
			}
			if v == "" {
				v = "dev"
			}
			versionFields := []string{v}
			if commit != "" {
				versionFields = append(versionFields, commit)
			}
			if date != "" {
				versionFields = append(versionFields, date)
			}
			fmt.Println(strings.Join(versionFields, " "))
			return nil
		},
	}
}

func newRunCommand() *ffcli.Command {
	cmd := "run"
	fs := flag.NewFlagSet(cmd, flag.ExitOnError)
	_ = fs.String("config", "", "config file (optional)")
	var cfg twi18n.Config
	fs.BoolVar(&cfg.Debug, "debug", false, "debug mode")

	fs.StringVar(&cfg.TwitterAPIKey1, "twitter-api-key-1", "", "twitter api key for main account")
	fs.StringVar(&cfg.TwitterAPISecret1, "twitter-api-secret-1", "", "twitter api secret for main account")
	fs.StringVar(&cfg.TwitterAccessToken1, "twitter-access-token-1", "", "twitter access token for main account")
	fs.StringVar(&cfg.TwitterAccessTokenSecret1, "twitter-access-token-secret-1", "", "twitter access token secret for main account")
	fs.StringVar(&cfg.TwitterAPIKey2, "twitter-api-key-1", "", "twitter api key for second account")
	fs.StringVar(&cfg.TwitterAPISecret2, "twitter-api-secret-2", "", "twitter api secret for second account")
	fs.StringVar(&cfg.TwitterAccessToken2, "twitter-access-token-2", "", "twitter access token for second account")
	fs.StringVar(&cfg.TwitterAccessTokenSecret2, "twitter-access-token-secret-2", "", "twitter access token secret for second account")

	fs.StringVar(&cfg.OpenAIKey, "openai-key", "", "openai key")
	fs.StringVar(&cfg.OpenAIModel, "openai-model", "", "openai model")
	fs.StringVar(&cfg.OpenAIHost, "openai-host", "", "openai host")

	fs.StringVar(&cfg.Post, "post", "", "post")
	fs.StringVar(&cfg.Template, "template", twi18n.DefaultTemplate, "template for translation prompt (use {TWEET} for tweet placeholder)")

	return &ffcli.Command{
		Name:       cmd,
		ShortUsage: fmt.Sprintf("twi18n %s [flags] <key> <value data...>", cmd),
		Options: []ff.Option{
			ff.WithConfigFileFlag("config"),
			ff.WithConfigFileParser(ffyaml.Parser),
			ff.WithEnvVarPrefix("TWI18N"),
		},
		ShortHelp: fmt.Sprintf("twi18n %s command", cmd),
		FlagSet:   fs,
		Exec: func(ctx context.Context, args []string) error {
			return twi18n.Run(ctx, &cfg)
		},
	}
}
