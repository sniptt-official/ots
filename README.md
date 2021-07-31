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

<p align="center">Secrets are descructed 💥 once viewed, or after specified expiry</p>

***

![render1626708858371](https://user-images.githubusercontent.com/84327906/126186752-156fe5bd-129a-4152-9dff-d7c01e581687.gif)

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

## Security

### Why should I trust you with my secrets?

All secrets are encrypted end-to-end, which means the plaintext values never leave your device. We do *not* log, track, share, or store the encryption key that protects your secret. You can check the client code to learn more about how we create the encryption key as well as what data is being sent to our servers.

### Is sharing via URL really secure?

Secrets created using the `ots new` command are what we refer to as "one-time secrets". Once they are retrieved by the recipient, they can no longer be viewed even if someone got hold of the URL. Furthermore, each one-time secret gets automatically deleted after specified duration if not viewed. By default, this is 24 hours but you can change this as required, for example `ots new -x 2h`.

It goes without saying that URL-accessible one-time secrets should be shared with **intended recipients only**.

### Can I persist my secrets for later use?

Please use [snip](https://github.com/sniptt-official/snip) instead.

## License

See [LICENSE](LICENSE)
