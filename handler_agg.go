package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't make the request struct: %v", err)
	}

	req.Header.Set("User-Agent", "gator")
	c := http.Client{
		Timeout: time.Minute,
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("something wrong with the response: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 399 {
		return nil, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling into RSSFeed struct: %v", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &feed, nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println("Channel:")
	fmt.Printf("\tTitle: \"%s\"\n", feed.Channel.Title)
	fmt.Printf("\tLink: \"%s\"\n", feed.Channel.Link)
	fmt.Printf("\tDescription: \"%s\"\n", feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		fmt.Printf("\tItem %d:\n", i+1)
		fmt.Printf("\t\tTitle: \"%s\"\n", item.Title)
		fmt.Printf("\t\tLink: \"%s\"\n", item.Link)
		fmt.Printf("\t\tDescription: \"%s\"\n", item.Description)
		fmt.Printf("\t\tPublish Date: \"%s\"\n", item.PubDate)
	}

	return nil
}
