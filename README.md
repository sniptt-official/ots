<p align="center">
  <b>Looking to self-host? Use <a href="https://github.com/sniptt-official/ots-aws">the official CDK construct</a></b>
</p>

<p align="center">
  <a href="https://sniptt.com">
    <img src=".github/assets/ots-social-cover.svg" alt="Ots Logo" />
  </a>
</p>

<p align="right">
  <i>If you use this repo, star it ✨</i>
</p>

***

<p align="center">🔐 <b>Share end-to-end encrypted secrets with others via a one-time URL</b></p>

<p align="center">Use to securely share API Keys, Signing secrets, Passwords, etc. with 3rd parties or with your team</p>

<p align="center">Secrets are destructed 💥 once viewed, or after specified expiry</p>

***

<p align="center"><b>Looking for the web version?</b></p>

<p align="center">https://ots.sniptt.com</p>

***

## Install

### Homebrew

The recommended way to install `ots` on macOS is via Homebrew.

```
brew install ots
```

### Go

```
go get -u github.com/sniptt-official/ots
```

## Usage

![render1628628123170](https://user-images.githubusercontent.com/778109/128932301-190388b3-171c-4e41-be5c-88ecf315beda.gif)

### Prompt

```
> ots new -x 2h
Enter your secret: 
```

### Pipeline

You can also use pipes, for example

```
pbpaste | ots new
```

or

```
cat .env | ots new
```

### Data residency

Use `--region` to choose where the secrets reside.

```
ots new -x 24h --region eu-central-1
```

### Self-hosting

Please refer to [the official CDK construct](https://github.com/sniptt-official/ots-aws) for detailed instructions.

Grab your API Gateway URL, API key and configure `~/.ots.yaml` (or whatever you provide to `--config`):

```yaml
apiUrl: https://YOUR_API_ID.execute-api.YOUR_REGION.amazonaws.com/prod/secrets
apiKey: YOUR_API_KEY
```

Use `ots` as before:

```
> ots new -x 2h
Using config file: /Users/xxx/.ots.yaml
Enter your secret: ***
Your secret is now available on the below URL.

https://my-ots-web-view.com/burn-secret?id=xxx&ref=ots-cli&region=us-east-1&v=debug#xxx

You should only share this URL with the intended recipient.

Please note that once retrieved, the secret will no longer
be available for viewing. If not viewed, the secret will
automatically expire at approximately xx xxx xxxx xx:xx:xx.
```

## FAQs

Please refer to our Q\&A discussions.

*   [Why should I trust you with my secrets?](https://github.com/sniptt-official/ots/discussions/13)
*   [Can I persist my secrets for later use?](https://github.com/sniptt-official/ots/discussions/15)
*   [What kind of limits are in place?](https://github.com/sniptt-official/ots/discussions/18)

## License

See [LICENSE](LICENSE)
