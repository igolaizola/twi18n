# twi18n

twi18n is a tool that allows you to post tweets in multiple languages using a single tweet.

## 📦 Installation

You can use the Golang binary to install **twi18n**:

```bash
go install github.com/igolaizola/twi18n/cmd/twi18n@latest
```

Or you can download the binary from the [releases](https://github.com/igolaizola/twi18n/releases)

## 📋 Requirements

### OpenAI API key

You need to obtain the api-key of openai or you can use instead a local model like llama3 using ollama.

#### openai settings

```
openai-key: YOUR_API_KEY
openai-model: gp-3.5-turbo
```

#### ollama and llama3 example

```
openai-model: llama3
openai-host: http://localhost:11434/v1
```

### Twitter API keys

For each account, the main one and the translated one you need the following Twitter API keys:

- `twitter-api-key`
- `twitter-api-secret-key`
- `twitter-access-token`
- `twitter-access-token-secret`

To obtain those keys you need to create a Twitter Developer account and create a new app.
You also need to enable "User authentication settings" with write permissions.

https://developer.twitter.com/en/portal/dashboard

## 🕹️ Usage

Launch commands using a yam file:

```bash
twi18n run --config twi18n.yaml
```

```bash
# twi18n.yaml
openai-key: YOUR_API_KEY # not needed for local models
openai-model: gp-3.5-turbo
openai-host: http://localhost:11434/v1 # required for local models
 # main account, here the original tweet will be posted
twitter-api-key-1: TWITTER_API_KEY
twitter-api-secret-key-1: TWITTER_API_SECRET_KEY
twitter-access-token-1: TWITTER_ACCESS_TOKEN
twitter-access-token-secret-1: TWITTER_ACCESS_TOKEN_SECRET
# translated account, here the translated tweet will be posted
twitter-api-key-2: TWITTER_API_KEY
twitter-api-secret-key-2: TWITTER_API_SECRET_KEY
twitter-access-token-2: TWITTER_ACCESS_TOKEN
twitter-access-token-secret-2: TWITTER_ACCESS_TOKEN_SECRET
# tweet to be translated
tweet: "Hola mundo"
# prompt template for the translation
template: "Translate the following tweet to English, return only the translation: {TWEET}"
```

## 💖 Support

If you have found my code helpful, please give the repository a star ⭐

Additionally, if you would like to support my late-night coding efforts and the coffee that keeps me going, I would greatly appreciate a donation.

You can invite me for a coffee at ko-fi (0% fees):

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/igolaizola)

Or at buymeacoffee:

[![buymeacoffee](https://user-images.githubusercontent.com/11333576/223217083-123c2c53-6ab8-4ea8-a2c8-c6cb5d08e8d2.png)](https://buymeacoffee.com/igolaizola)

Donate to my PayPal:

[paypal.me/igolaizola](https://www.paypal.me/igolaizola)

Sponsor me on GitHub:

[github.com/sponsors/igolaizola](https://github.com/sponsors/igolaizola)

Or donate to any of my crypto addresses:

- BTC `bc1qvuyrqwhml65adlu0j6l59mpfeez8ahdmm6t3ge`
- ETH `0x960a7a9cdba245c106F729170693C0BaE8b2fdeD`
- USDT (TRC20) `TD35PTZhsvWmR5gB12cVLtJwZtTv1nroDU`
- USDC (BEP20) / BUSD (BEP20) `0x960a7a9cdba245c106F729170693C0BaE8b2fdeD`
- Monero `41yc4R9d9iZMePe47VbfameDWASYrVcjoZJhJHFaK7DM3F2F41HmcygCrnLptS4hkiJARCwQcWbkW9k1z1xQtGSCAu3A7V4`

Thanks for your support!
