# mnemonic
`mnemonic` is a diceware alike password generator written in Go.

## introduction
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

## how to use it

```
Usage of ./mnemonic:
  -d string
        delimeter as split element for the password (default " ")
  -n uint
        the number of words (default 3)
```

On default `mnemonic` will return 3 random words splitted by a space.

## todo

* provide more word lists (english is a must have)
* provide a `-m`, `--max` parameter for setting a characterlimit. Useful for websites with password character limit.
* tests, tests.. tests
* build release binaries magically via github actions
* provide a flag for choosing the language of the used word list