package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abadojack/whatlanggo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	url = "https://kapi.kakao.com/v1/translation/translate"
	apikey = "KakaoAK c9fe95d1f4043ea14b2e99d9a440307a"
)

type Data struct {
	Translated_text [][]string `json:"translated_text"`
}

func main() {
	srcPtr := flag.String("src", "en", "source language.")                                                    //src lang
	targetPtr := flag.String("target", "kr", "target language.")                                              //target lang
	textPtr := flag.String("text", "", "sentence or word to transsente.")                                                               //--data-urlencod
	debugPtr := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		reader := bufio.NewReader(os.Stdin)
		t, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
		if string(t) == "\n" {
			flag.PrintDefaults()
		}
		fmt.Println(string(t))
		*textPtr = string(t)
	}
	if flag.NArg() == 1 {
		*textPtr = flag.Args()[0]
	}
	userLang := whatlanggo.Detect(*textPtr)
	if *debugPtr {
		fmt.Println("Detect Language : " + whatlanggo.LangToString(userLang.Lang))
	}
	if *textPtr == "" {
		flag.PrintDefaults()
	}
	switch whatlanggo.LangToString(userLang.Lang) {
	case "kor": // 한국어
		*srcPtr = "kr"
		*targetPtr = "en"
	case "eng": // 영어
		*srcPtr = "en"
		*targetPtr = "kr"
	case "cmn": // 중국
		*srcPtr = "cn"
		*targetPtr = "kr"
	case "nya": // 인도네시아
		*srcPtr = "id"
		*targetPtr = "kr"
	case "vie": // 베트남
		*srcPtr = "vi"
		*targetPtr = "kr"
	case "jpn": // 일본
		*srcPtr = "jp"
		*targetPtr = "kr"
	default:
		*srcPtr = "en"
		*targetPtr = "kr"
	}
	body := strings.NewReader(fmt.Sprintf(`src_lang=%s&target_lang=%s&query=%s`, *srcPtr, *targetPtr, *textPtr))
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", apikey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if *debugPtr {
		fmt.Println(string(b))
	}
	data := Data{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}
	if *debugPtr {
		fmt.Println(data)
	}
	for _, i := range data.Translated_text {
		for _, j := range i {
			fmt.Printf(j)
		}
		fmt.Println()
	}
}
