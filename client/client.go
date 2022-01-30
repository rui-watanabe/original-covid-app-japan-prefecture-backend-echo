package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetClient() (p Prefectures) {
	c, err := NewClient("https://opendata.corona.go.jp/api/covid19DailySurvey")
	if err != nil {
		log.Fatalln(err)
	}

	// http.Response として返却
	res, err := c.Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatalln(err)
	}
	return p
}
