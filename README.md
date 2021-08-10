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

<p align="center">https://www.sniptt.com/ots</p>

***

## Install

### Homebrew

The recommended way to install `ots` on macOS is via Homebrew.

```sh
brew tap sniptt-official/tap
brew install ots
```

*NOTE: We need 30 forks, 30 watchers, and 75 stars to make it to Homebrew/core. Please help us get there 👀!*

### Go

```sh
go get -u github.com/sniptt-official/ots
```

### Manual

For manual installation instructions on macOS and Linux, please refer to the dedicated [install docs](./docs/manual-install.md).

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

## FAQs

Please refer to our Q&A discussions.

- [Why should I trust you with my secrets?](https://github.com/sniptt-official/ots/discussions/13)
- [Can I persist my secrets for later use?](https://github.com/sniptt-official/ots/discussions/15)
- [What kind of limits are in place?](https://github.com/sniptt-official/ots/discussions/18)

## License

See [LICENSE](LICENSE)
