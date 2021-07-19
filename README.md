<p align="center" style="text-align:center;">
  <a href="https://sniptt.com">
    <img src=".github/assets/readme-hero-logo.svg" alt="Sniptt Logo" />
  </a>
</p>

<p align="right">
  <i>If you use this repo, star it ✨</i>
</p>

***

<div align="center">🔐 Share end-to-end encrypted secrets with others via a one-time URL</div>

***

![render1626708858371](https://user-images.githubusercontent.com/84327906/126186752-156fe5bd-129a-4152-9dff-d7c01e581687.gif)

## Install

### Homebrew

The recommended way to install `ots` on macOS is via Homebrew.

```sh
brew install sniptt-official/ots/ots
```

#### Update

To update to latest version of `ots`, use:

```sh
brew upgrade sniptt-official/ots/ots
```

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

All secrets are **end-to-end encrypted**, which means the plaintext values **never leave your device**. We do *not* log, track, share, or store the encryption key that protects your secret. You can check the client code to learn more about how we create the encryption key as well as what data is being sent to our servers.

### Is sharing via URL really secure?

Secrets created using the `ots new` command are what we refer to as "one-time secrets". Once they are retrieved by the recipient, they can no longer be viewed even if someone got hold of the URL. Furthermore, each one-time secret gets automatically deleted after specified duration if not viewed. By default, this is 24 hours but you can set yours, for example `ots new -x 2h`.

However, it goes without saying that URL-accessible one-time secrets should be shared with **intended recipients only**.

### Can I persist my secrets for later use?

Please use the [snip-cli](https://github.com/sniptt-official/snip-cli) instead.

## License

See [LICENSE](LICENSE)
