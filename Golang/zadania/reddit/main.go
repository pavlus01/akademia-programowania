package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reddit/fetcher"
	"time"
)

func main() {
	//var f fetcher.RedditFetcher // do not change
	// var w io.Writer // do not change
	c := NewFetcher("https://oauth.reddit.com/r/golang", time.Second*5)
	err := c.Fetch(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%+v", d.Data.Children)
	// for index, element := range d.Data.Children {
	// 	log.Println(index)
	// 	log.Println(element.Data.Title)
	// }
	f2, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY, 0600)
	// buf := new(bytes.Buffer)
	// enc := json.NewEncoder(buf)
	// if err := enc.Encode(d.Data.Children); err != nil {
	// 	panic(err)
	// }
	err2 := c.Save(f2)
	if err2 != nil {
		log.Fatal(err2)
	}

	// f.Write(buf.Bytes())
}

type Fetcher struct {
	c    *http.Client
	host string
	data fetcher.Response
}

func NewFetcher(host string, t time.Duration) *Fetcher {
	return &Fetcher{
		host: host,
		c: &http.Client{
			Timeout: t,
		},
	}
}

func (e *Fetcher) Save(writer io.Writer) error {
	for _, element := range e.data.Data.Children {
		_, err := writer.Write([]byte(element.Data.Title + "\n" + element.Data.URL + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Fetcher) Fetch(ctx context.Context) error {
	ctx = context.WithValue(ctx, "requestID", time.Now().Unix())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, e.host, http.NoBody)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}
	req.Header.Add("Authorization", "bearer 23647321236545-YceLNgAn65Xk339PUlzdJWxjqiKFMw")

	resp, err := e.c.Do(req)
	if err != nil {
		return fmt.Errorf("cannot get data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data fetcher.Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("cannot unmarshal data: %w", err)
	}
	e.data = data
	return nil
}
