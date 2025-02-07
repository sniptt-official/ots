<p align="center">
  <a href="https://sniptt.com">
    <img src=".github/assets/ots-social-cover.svg" alt="OTS Logo" />
  </a>
</p>

<p align="right">
  <i>If you use this repo, please star it ✨</i>
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

## FAQs

Please refer to our Q\&A discussions.

*   [Why should I trust you with my secrets?](https://github.com/sniptt-official/ots/discussions/13)
*   [Can I persist my secrets for later use?](https://github.com/sniptt-official/ots/discussions/15)
*   [What kind of limits are in place?](https://github.com/sniptt-official/ots/discussions/18)

## License

See [LICENSE](LICENSE)
