# mnemonic
`mnemonic` is a diceware alike password generator written in Go.

## features

* specify number of words
* different languages: English, German
* specify a delimeter between words
* use an external file as word list

## how to use it

```
Usage of ./mnemonic:
  -d string
        delimeter as split element for the password (default "-")
  -f string
        path to external word list
  -l string
        language of word list: [en, de] (default "en")
  -n uint
        number of words (default 3)
```

## background
The history of mnemonic starts with the XKCD comic about easy rememberable passwords:

[https://xkcd.com/936/](https://xkcd.com/936/)

I am aware that other solutions exist, but most of them are websites.
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

## todo

* provide more word lists
* provide a `-m`, `--max` parameter for setting a characterlimit. Useful for websites with password character limit.
* tests, tests.. tests
* build release binaries magically via github actions
* option for generating random numbers as padding/delimeter