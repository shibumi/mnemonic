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

func newWordList() (list []string, err error) {
	asset, err := Asset("lists/frequency.txt")
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
	list, err := newWordList()
	if err != nil {
		log.Fatalln(err)
	}
	numberOfWords := flag.Uint("n", 3, "the number of words")
	delimeter := flag.String("d", "-", "delimeter as split element for the password")
	flag.Parse()

	fmt.Println(generatePassword(*numberOfWords, *delimeter, list))
}
