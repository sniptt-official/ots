<p align="center">
  <b>👀 The web version has now moved to <a href="https://ots.sniptt.com">ots.sniptt.com</a> 👀</b>
</p>

<p align="center">
  Please upgrade the CLI to <a href="https://github.com/sniptt-official/ots/releases/tag/v0.1.0">version 0.1.0</a>
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

```sh
brew install ots
```

### Go

```sh
go get -u github.com/sniptt-official/ots
```

### Manual

Please refer to the [manual install](./docs/manual-install.md) doc.

## Usage

![render1628628123170](https://user-images.githubusercontent.com/778109/128932301-190388b3-171c-4e41-be5c-88ecf315beda.gif)

### Prompt

```sh
$ ots new -x 2h
Enter your secret: 
```

### Pipeline

You can also use pipes, for example

```sh
$ pbpaste | ots new
```

or

```sh
$ cat .env | ots new
```

### Data residency

Use `--region` to choose where the secrets reside.

```sh
$ ots new -x 24h --region eu-central-1
```

## FAQs

Please refer to our Q\&A discussions.

*   [Why should I trust you with my secrets?](https://github.com/sniptt-official/ots/discussions/13)
*   [Can I persist my secrets for later use?](https://github.com/sniptt-official/ots/discussions/15)
*   [What kind of limits are in place?](https://github.com/sniptt-official/ots/discussions/18)

## License

See [LICENSE](LICENSE)
