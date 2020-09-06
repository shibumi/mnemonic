# mnemonic
[![build](https://github.com/shibumi/mnemonic/workflows/build/badge.svg)](https://github.com/shibumi/mnemonic/actions?query=workflow%3Abuild) [![Coverage Status](https://coveralls.io/repos/github/shibumi/mnemonic/badge.svg)](https://coveralls.io/github/mnemonic/go-pathspec) [![PkgGoDev](https://pkg.go.dev/badge/github.com/shibumi/mnemonic)](https://pkg.go.dev/github.com/shibumi/mnemonic)  
`mnemonic` is a diceware alike password generator written in Go.

## features

* specify number of words
* support for different languages
* specify a delimiter between words
* use an external file as word list

### supported languages

* German
* English
* Japanese
* French
* Korean
* Italian
* Spanish
* Czech
* Chinese (simplified)
* Chinese (traditional)

### roadmap (upcoming features)

* Provide a library to work with (useful for password generators, password stores, etc)
* Set a max limit for length of words/chars
* Random numbers as delimiter
* Leetspeak support eg: `hello-dear` to `h3ll0-d34r`
* Makefile and automated builds/releases on Github

## how to use it

```
Usage of ./mnemonic:
  -d string
        delimiter as split element for the password (default "-")
  -f string
        path to external word list
  -l string
        language of word list: [en, de, cz, jp, kr, it, es, fr, zh-cn, zh-tw] (default "en")
  -n uint
        number of words (default 3)
```

## background
The history of mnemonic starts with the XKCD comic about easy rememberable passwords:

[![https://xkcd.com/936/](https://imgs.xkcd.com/comics/password_strength.png)](https://xkcd.com/936/)

I am aware, that other solutions exist, but most of them are websites.
Websites are not trustable, even if password generation happens via javascript (evil plugins, that can read your website content).
With `Go` we finally have a programming language that allows us to build a static binary with embedded data.

## how to install
Just download the binary from the release page.

## building from source
For building from source you need to install `Go`.
Just clone this repository and execute: `go build -o mnemonic`.
This should drop a binary in the project directory named `mnemonic`.

## building static binaries for different platforms

### Windows

`CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mnemonic-windows-amd64.ex`

### Macbook

`CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mnemonic-apple-amd64`

### Linux

`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mnemonic-linux-amd64`