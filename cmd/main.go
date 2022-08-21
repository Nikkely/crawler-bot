package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/nikkely/crawler-bot/pkg/config"
	"github.com/nikkely/crawler-bot/pkg/source"
)

func main() {
	c := config.NewConfig("./config/config.yml")

	y := source.NewYoutubeSource(c)
	result, err := y.Get("#シャドウバースエボルヴ")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%v", result)
	outputJson(result)
	os.Exit(1)
}

func outputJson(r any) {
	outputPath := fmt.Sprintf("./output/%s", time.Now().Format("2006-01-02T15:04:05"))
	err := ioutil.WriteFile(outputPath, []byte(fmt.Sprintf("%v", r)), 0666)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
