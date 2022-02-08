package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
)

type ApiClient interface {
	Request() (Prefectures, error)
}

type DataGet struct {
	client ApiClient // インターフェイスに依存しているだけで実装は存在しない
}

func (d *DataGet) Get(req ApiClient) (Prefectures, error) {
	result, err := req.Request()
	if err != nil {
		return make(Prefectures, 0), err
	}
	return result, nil
}

func (d *DataGet) Request() (p Prefectures, err error) {
	c, err := NewClient("http://localhost:3000/api/client/data")
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
	return
}
