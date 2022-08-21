package source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"net/http"
	"net/url"

	"github.com/nikkely/crawler-bot/pkg/config"
)

type youtubeSource struct {
	Config config.Config
}

func NewYoutubeSource(c config.Config) youtubeSource {
	return youtubeSource{
		Config: c,
	}
}

func formatRFC3339(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z07:00:00")
}

func (s youtubeSource) makeURL(query string, publishedAfter time.Time) string {
	u, err := url.Parse("https://www.googleapis.com/youtube/v3/search")
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("key", s.Config.YoutubeAPIKey)
	q.Set("part", "snippet")
	q.Set("type", "video")
	q.Set("q", query)
	// q.Set("publishedAfter", formatRFC3339(publishedAfter))
	u.RawQuery = q.Encode()
	log.Println(u.String())
	return u.String()
}

func (s youtubeSource) Get(keyword string) (*YoutubeSearchAPIResponse, error) {
	url := s.makeURL(keyword, time.Now().Add(time.Hour*-200))
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got HTTP Status:%d, URL:%s", res.StatusCode, url)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	var result YoutubeSearchAPIResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse responce (%s)", err.Error())
	}
	return &result, nil
}
