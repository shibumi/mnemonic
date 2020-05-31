package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func newWordList(language string) (list []string, err error) {
	var assetFile string
	switch language {
	case "de":
		assetFile = "lists/frequency.txt"
	default:
		assetFile = "lists/tothink.txt"
	}

	asset, err := Asset(assetFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(asset))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text() != "" {
			list = append(list, scanner.Text())
		}
	}
	return list, nil
}

func generatePassword(n uint, d string, list []string) (pw string) {
	seed := rand.NewSource(time.Now().UnixNano())
	for i := uint(0); i < n; i++ {
		generator := rand.New(seed)
		pw += list[generator.Intn(len(list))]
		if i < n-1 {
			pw += d
		}
	}
	return pw
}

func main() {
	numberOfWords := flag.Uint("n", 3, "the number of words")
	delimeter := flag.String("d", "-", "delimeter as split element for the password")
	language := flag.String("l", "en", "language of word list: [en, de]")
	flag.Parse()

	list, err := newWordList(*language)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(generatePassword(*numberOfWords, *delimeter, list))
}
