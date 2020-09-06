package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func newWordList(language string, file string) (list []string, err error) {
	var asset []byte
	if file == "" {
		var assetFile string
		switch language {
		case "de":
			assetFile = "data/german.txt"
		default:
			assetFile = "data/english.txt"
		}
		asset, err = Asset(assetFile)
		if err != nil {
			return nil, err
		}
	} else {
		asset, err = ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
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
	numberOfWords := flag.Uint("n", 3, "number of words")
	delimiter := flag.String("d", "-", "delimiter as split element for the password")
	language := flag.String("l", "en", "language of word list: [en, de]")
	file := flag.String("f", "", "path to external word list")
	flag.Parse()

	list, err := newWordList(*language, *file)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(generatePassword(*numberOfWords, *delimiter, list))
}
